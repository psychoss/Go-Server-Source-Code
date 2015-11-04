package level
import (
	//"bytes"
	"bytes"
	"fmt"
	"github.com/jmhodges/levigo"
	"math/rand"
	"sync"
	"time"
)
type Level struct {
	mu sync.Mutex
	db *levigo.DB
}
// 一个用来筛选键和处理数据的接口
type Filter interface {
	Accept([]byte) bool
	Process(key, vaule []byte) interface{}
}
func (l *Level) GetRandomContents(amout int, filter Filter) []interface{} {
	// 获取所有键
	keys := l.GetAllKeys()
	// 线程锁，防止其他请求同时执行这个方法
	l.mu.Lock()
	// 当方法执行完成时释放锁
	defer l.mu.Unlock()
	// 获取键的长度
	keyLength := len(keys)
	// 最多能读取的长度
	stopPos := keyLength
	// 如果待读取的长度大于整个长度，读取所有
	if amout < stopPos {
		stopPos = amout
	}
	// 生成随机数的内嵌函数
	randProvier := func(min, max int) int {
		rand.Seed(time.Now().UTC().UnixNano())
		return min + rand.Intn(max-min)
	}
	// 随机数
	var randInt int
	// 防止获取重复随机数的map
	unique := map[int]bool{}
	// 读取 LevelDB 时使用的选项
	ro := levigo.NewReadOptions()
	// 延迟关闭选项
	defer ro.Close()
	// 不缓存，因为需要大量读取
	ro.SetFillCache(false)
	var v []interface{}
	// 防止进入死循环
	retry, maxRetry := 0, 100
	// 循环
	for stopPos > -1 {
		//获取随机数
		randInt = randProvier(0, keyLength)
		// 如果此随机数已使用，跳过继续
		if unique[randInt] {
			if retry == maxRetry {
				break
			}
			retry++
			continue
		} else {
			// 标记为已使用
			unique[randInt] = true
		}
		// 获取键
		key := keys[randInt]
		// 过滤掉不想要的键
		if !filter.Accept(key) {
			if retry == maxRetry {
				break
			}
			retry++
			continue
		}
		//读取内容
		r, err := l.db.Get(ro, key)
		if len(r) == 0 {
			continue
		}
		if err != nil {
			break
		}
		v = append(v, filter.Process(key, r))
		stopPos--
	}
	return v
}
func (l *Level) GetAllKeysByAnchor(anchor []byte) []byte {
	l.mu.Lock()
	defer l.mu.Unlock()
	var buffer bytes.Buffer
	ro := levigo.NewReadOptions()
	defer ro.Close()
	ro.SetFillCache(false)
	it := l.db.NewIterator(ro)
	defer it.Close()
	buffer.WriteByte('[')
	it.Seek(anchor)
	if !it.Valid() {
		return nil
	}
	it.Next()
	for it.Valid() && bytes.HasPrefix(it.Key(), anchor) {
		buffer.Write(it.Value())
		buffer.WriteByte(',')
		it.Next()
        
	}
	bs := buffer.Bytes()
	bs = bytes.TrimRight(bs, ",")
	bs = append(bs, ']')
	if err := it.GetError(); err != nil {
		fmt.Println(err)
	}
	return bs
}
func (l *Level) GetAllKeys() [][]byte {
	l.mu.Lock()
	defer l.mu.Unlock()
	ro := levigo.NewReadOptions()
	defer ro.Close()
	ro.SetFillCache(false)
	it := l.db.NewIterator(ro)
	defer it.Close()
	var hold [][]byte
	for it.SeekToFirst(); it.Valid(); it.Next() {
		hold = append(hold, it.Key())
	}
	if err := it.GetError(); err != nil {
		fmt.Println(err)
	}
	return hold
}
func (l *Level) BatchRead(v ...interface{}) [][]byte {
	l.mu.Lock()
	defer l.mu.Unlock()
	ro := levigo.NewReadOptions()
	ro.SetFillCache(false)
	var r [][]byte
	for _, value := range v {
		rs, err := l.db.Get(ro, value.([]byte))
		if err != nil {
			break
		}
		r = append(r, rs)
	}
	return r
}
// Save record to levelDB
func (l *Level) Put(key, value []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	wo := levigo.NewWriteOptions()
	err := l.db.Put(wo, key, value)
	wo.Close()
	return err
}
func (l *Level) Get(key []byte) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	ro := levigo.NewReadOptions()
	bs, err := l.db.Get(ro, key)
	ro.Close()
	return bs, err
}
// Delete the record by key
// Concurrent-safe
func (l *Level) Delete(key []byte) {
	l.mu.Lock()
	defer l.mu.Unlock()
	wo := levigo.NewWriteOptions()
	l.db.Delete(wo, key)
	wo.Close()
}
func (l *Level) Close() {
	l.db.Close()
}
func New(directory string) (*Level, error) {
	options := levigo.NewOptions()
	options.SetCreateIfMissing(true)
	options.SetCompression(levigo.SnappyCompression)
	db, err := levigo.Open(directory, options)
	return &Level{db: db}, err
}
// go run /home/psycho/go/src/web/db/leveldb.go
