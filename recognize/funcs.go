package recognize

import (
	"blueprintz/global"
	"github.com/gearboxworks/go-status"
	"github.com/gearboxworks/go-status/only"
	"net/http"
)

func IsValidType(c Componenter, r Recognizer) (ok bool) {
	cts := r.ValidTypes()
	for _, ct := range cts {
		if ct != c.GetType() {
			continue
		}
		ok = true
		break
	}
	return ok
}

func VerifyUrl(url global.Url) (sts Status) {
	for range only.Once {
		res, err := http.Head(url)
		if err != nil {
			sts = status.Fail().SetLogTo(status.DebugLog).
				SetMessage("HTTP HEAD request failed on '%s': %s", url, err.Error())
			break
		}
		if res == nil {
			sts = status.Fail().SetLogTo(status.WarnLog).
				SetMessage("HTTP HEAD request returned nil result for '%s'", url)
			break
		}
		if res.ContentLength == 0 {
			sts = status.Fail().SetLogTo(status.DebugLog).
				SetMessage("HTTP HEAD request returned content length of 0 result for '%s'", url)
			break
		}
		if res.StatusCode != http.StatusOK {
			sts = status.Fail().SetLogTo(status.DebugLog).
				SetMessage("HTTP HEAD request returned status of '%s' for '%s'", res.Status, url)
			break
		}
		sts = status.Success("HTTP HEAD request succeeded for '%s'", url)
	}
	return sts
}
