package util
import (
	"testing"
	"web/util"
)
func TestUnion(t *testing.T) {
	a := []byte("123")
	b := []byte("456")
	if string(util.Bytes.Union(a, b)) != "123456" {
		t.Error("util.Bytes.Union")
	}
}
func TestLeftN(t *testing.T) {
	a := []byte("123")
	if util.Bytes.LeftN(a, 0) != nil {
		t.Error("util.Bytes.LeftN")
	}
	if string(util.Bytes.LeftN(a, 1)) != "1" {
		t.Error("util.Bytes.LeftN")
	}
}
// go run /home/psycho/go/src/web/util/bytes_test.go
