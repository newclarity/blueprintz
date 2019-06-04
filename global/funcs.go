package global

import "strings"

func StringSliceIndex(slc []string, str string) int {
	pos := -1
	for i, v := range slc {
		if v != str {
			continue
		}
		pos = i
		break
	}
	return pos
}
func StringSlicePad(slc []string, n int) []string {
	for i, t := range slc {
		slc[i] = strings.Repeat(" ", n) + t + strings.Repeat(" ", n)
	}
	return slc
}
