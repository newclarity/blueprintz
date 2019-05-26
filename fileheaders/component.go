package fileheaders

import (
	"blueprintz/global"
	"blueprintz/jsonfile"
	"blueprintz/util"
	"fmt"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

var NilComponent = (*Component)(nil)
var _ Componenter = NilComponent
var _ jsonfile.Componenter = NilComponent

type HeaderValueFieldMap = map[global.FileHeader]*reflect.Value

type ComponenterMap map[global.Slug]Componenter

type Componenters []Componenter

type Componenter interface {
	GetName() global.ComponentName
	GetType() global.ComponentType
	GetSlug() global.Slug
	SetFilepath(global.Filepath)
	GetFilepath() global.Filepath
	ReadHeader(Componenter) Status
	GetHeaderValueFieldMap(...Componenter) HeaderValueFieldMap
	AllowHeaderless() bool
	SetAllowHeaderless(bool)
	IsRootFile() bool
	SetIsRootFile(bool)
}

type Component struct {
	Filepath        global.Filepath
	Description     string         `fileheader:"Description"`
	Version         global.Version `fileheader:"Version"`
	Author          string         `fileheader:"Author"`
	AuthorURI       global.Url     `fileheader:"Author URI"`
	License         string         `fileheader:"License"`
	LicenseURI      global.Url     `fileheader:"License URI"`
	TextDomain      string         `fileheader:"Text Domain"`
	DomainPath      string         `fileheader:"Domain Path"`
	allowheaderless bool
	isrootfile      bool
}

func (me *Component) GetFilepath() global.Filepath {
	return me.Filepath
}

func (me *Component) SetFilepath(fp global.Filepath) {
	me.Filepath = fp
}

func NewComponent(fp global.Filepath) *Component {
	return &Component{
		Filepath: fp,
	}
}

var panicMsg = "Cannot %s() of fileheaders.Component; use fileheaders.Plugin or fileheaders.Theme instead."

func (me *Component) GetName() global.ComponentName {
	panic(fmt.Sprintf(panicMsg, "GetName"))
}

func (me *Component) GetVersion() global.Version {
	return me.Version
}

func (me *Component) GetDownloadUrl() global.Url {
	panic(fmt.Sprintf(panicMsg, "GetComponentDownloadUrl"))
}

func (me *Component) GetSubdir() global.Slug {
	return filepath.Base(filepath.Dir(me.Filepath))
}

func (me *Component) GetBasefile() global.Slug {
	return filepath.Base(me.Filepath)
}

func (me *Component) GetSlug() global.Slug {
	return filepath.Base(filepath.Dir(me.Filepath))
}

func (me *Component) GetWebsite() global.Url {
	panic(fmt.Sprintf(panicMsg, "GetWebsite"))
}

func (me *Component) GetType() global.ComponentType {
	panic(fmt.Sprintf(panicMsg, "GetType"))
}

func (me *Component) AllowHeaderless() bool {
	return me.allowheaderless
}
func (me *Component) SetAllowHeaderless(allow bool) {
	me.allowheaderless = allow
}

func (me *Component) IsRootFile() bool {
	return me.isrootfile
}
func (me *Component) SetIsRootFile(isrootfile bool) {
	me.isrootfile = isrootfile
}

func (me *Component) GetSourceType() global.SourceType {
	panic(fmt.Sprintf(panicMsg, "GetSourceType"))
}

var headerFinder *regexp.Regexp

func init() {
	headerFinder = regexp.MustCompile(`(Plugin|Theme)\s+Name:`)
}

func (me *Component) ReadHeader(component Componenter) (sts Status) {
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
				SetWarn(true).
				SetMessage("unable to open '%s'", me.Filepath)
		}
		b := make([]byte, 8192) // Same size as WordPress uses
		_, err = file.Read(b)
		if err != nil {
			sts = status.Wrap(err).
				SetWarn(true).
				SetMessage("unable to read from '%s'", me.Filepath)
		}
		if headerFinder.Match(b) {
			// @TODO Replace these two lines with a regex
			headertxt := strings.Replace(string(b), "\r", "\n", -1)
			headertxt = strings.Replace(headertxt, "\n\n", "\n", -1)

			for h, f := range me.GetHeaderValueFieldMap(component) {
				// Same regex logic in WordPress' get_file_data()
				regex := fmt.Sprintf("(?im)^[ \t/*#@]*%s:(.*)$", regexp.QuoteMeta(h))
				re := regexp.MustCompile(regex)
				m := re.FindStringSubmatch(headertxt)
				// @todo Include this fix:  https://core.trac.wordpress.org/ticket/8497
				// Look in WordPress core code for get_file_data() to see the regex used, or:
				// https://core.trac.wordpress.org/attachment/ticket/8497/8497.diff
				if m == nil {
					continue
				}
				f.SetString(strings.TrimSpace(m[1]))
			}
			break
		}
		if component.AllowHeaderless() && component.IsRootFile() {
			noop()
			break
		}
		sts = status.Warn("file '%s' is not a %s header file", me.Filepath, component.GetType())
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
