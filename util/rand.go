package util
import(	"math/rand"
      "time")
// A type for random utilities
// 随机能力的类型
type utilRand struct {
}
// a instance of utilRand for hold the public functions
var Rand = utilRand{}
func (*utilRand) RandomInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// go run /home/psycho/go/src/web/util/rand.go