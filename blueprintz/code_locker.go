package blueprintz

import "blueprintz/global"

type CodeLockers []*CodeLocker
type CodeLocker struct {
	Url  global.Url
	Type global.LockerType
}
