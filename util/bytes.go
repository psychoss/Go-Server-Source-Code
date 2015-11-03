package util

type utilBytes struct {
}
var Bytes = utilBytes{}
func (*utilBytes) Union(a, b []byte) []byte {
	return append(a, b[0:]...)
}
func (*utilBytes) RightN(a []byte, start int) []byte {
    l :=len(a)
	if start+1 > l || start<1{
		return nil
	}
	return a[start:]
}
func (*utilBytes) LeftN(a []byte, start int) []byte {
    l :=len(a)
	if start+1 > l || start<1{
		return nil
	}
    return a[:start]
}
// go run /home/psycho/go/src/web/util/bytes.go
