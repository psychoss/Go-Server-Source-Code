package util
type utilString struct {
}
var String = utilString{}
func (*utilString) ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func (*utilString) Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
// go run /home/psycho/go/src/web/util/string.go
