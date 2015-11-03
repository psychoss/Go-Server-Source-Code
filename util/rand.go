package util
import (
	"math/rand"
	"time"
)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
func String(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(RandomInt('a', 'z'))
		time.Sleep(100 * time.Nanosecond)
	}
	return string(bytes)
}
// go run /home/psycho/go/src/web/util/rand.go
