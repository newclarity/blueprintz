package fileheaders

import (
	"blueprintz/global"
	"blueprintz/only"
	"blueprintz/util"
	"fmt"
	"github.com/gearboxworks/go-status"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

var NilComponent = (*Component)(nil)
var _ Componenter = NilComponent

type HeaderValueFieldMap = map[global.FileHeader]*reflect.Value

type Componenter interface {
	GetHeaderValueFieldMap(component ...Componenter) HeaderValueFieldMap
}

type Component struct {
	Filepath    global.Filepath
	Description string         `fileheader:"Description"`
	Version     global.Version `fileheader:"Version"`
	Author      string         `fileheader:"Author"`
	AuthorURI   global.Url     `fileheader:"Author URI"`
	License     string         `fileheader:"License"`
	LicenseURI  global.Url     `fileheader:"License URI"`
	TextDomain  string         `fileheader:"Text Domain"`
	DomainPath  string         `fileheader:"Domain Path"`
}

func NewComponent(fp global.Filepath) *Component {
	return &Component{
		Filepath: fp,
	}
}

func (me *Component) GetSlug() global.Slug {
	return filepath.Base(filepath.Dir(me.Filepath))
}

func (me *Component) Read(component Componenter) (sts Status) {
	for range only.Once {
		if me.Filepath == "" {
			log.Fatalf("component filepath is empty")
		}
		if !util.FileExists(me.Filepath) {
			sts = status.Warn("file '%s' does not exist", me.Filepath)
			break
		}
		file, err := os.Open(me.Filepath)
		if err != nil {
			sts = status.Wrap(err).
				SetWarning(true).
				SetMessage("unable to open '%s'", me.Filepath)
		}
		b := make([]byte, 8192) // Same size as WordPress uses
		_, err = file.Read(b)
		if err != nil {
			sts = status.Wrap(err).
				SetWarning(true).
				SetMessage("unable to read from '%s'", me.Filepath)
		}
		if !strings.Contains(string(b), "Plugin Name:") {
			sts = status.Warn("file '%s' is not a plugin header file", me.Filepath)
			break
		}
		// @TODO Replace these two lines with on regex
		headertxt := strings.Replace(string(b), "\r", "\n", -1)
		headertxt = strings.Replace(headertxt, "\n\n", "\n", -1)

		for h, f := range me.GetHeaderValueFieldMap(component) {
			// Same regex logic in WordPress' get_file_data()
			regex := fmt.Sprintf("(?im)^[ \t/*#@]*%s:(.*)$", regexp.QuoteMeta(h))
			re := regexp.MustCompile(regex)
			m := re.FindStringSubmatch(headertxt)
			if m == nil {
				continue
			}
			f.SetString(strings.TrimSpace(m[1]))
		}
	}
	return sts
}

func (me *Component) GetHeaderValueFieldMap(component ...Componenter) (vm HeaderValueFieldMap) {
	if len(component) == 0 {
		panic("component not passed as a parameter")
	}
	tme := reflect.TypeOf(me)
	vme := reflect.ValueOf(me)
	if tme.Kind() == reflect.Ptr {
		tme = tme.Elem()
		vme = vme.Elem()
	}
	tc := reflect.TypeOf(component[0])
	vc := reflect.ValueOf(component[0])
	if tc.Kind() == reflect.Ptr {
		tc = tc.Elem()
		vc = vc.Elem()
	}
	vm = make(HeaderValueFieldMap, tme.NumField()+tc.NumField())
	vs := []reflect.Value{vc, vme}
	for x, t := range []reflect.Type{tc, tme} {
		for i := 0; i < t.NumField(); i++ {
			tag := t.Field(i).Tag.Get("fileheader")
			if tag == "" {
				continue
			}
			f := vs[x].Field(i)
			vm[tag] = &f
		}
	}
	return vm
}
