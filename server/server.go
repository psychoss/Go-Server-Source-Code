package server
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"web/level"
	"web/rubex"
)
var (
	rxp       *rubex.Regexp
	staticRxp *rubex.Regexp
	leveldb   *level.Level
)
type RequestHandler struct {
}
func (*RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	filter(w, req)
}
func filter(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	bu := []byte(url)
	if url == "/" {
		indexHandler(w, req)
	} else if staticRxp.Match(bu) {
		staticHandler(w, url)
	} else if rxp.Match(bu) {
		fecthHandler(w, req)
	}
	if req.Method == "POST" {
		if url == "/update" {
			updateHandler(w, req)
		} else if url == "/post" {
			getHandler(w, req)
		}
	}
	fmt.Println(url)
}
// Cache the file for nginx
func cacheFile(fileName string, datas []byte) {
	ioutil.WriteFile(PATH_CACHE+fileName+".html", datas, 07777)
}
// 允许跨域方法
// Allow cross
func crossDomain(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
func Close() {
	leveldb.Close()
}
func StartNewServer(address string) error {
	var err error
	serverHandler := &RequestHandler{}
	leveldb, err = level.New(PATH_DB)
	if err != nil {
		return err
	}
	rxp = rubex.MustCompile("^/(?:css|go|mongo|node|python|rethinkdb|rust)[\\-a-zA-Z0-9/]*$")
	staticRxp = rubex.MustCompile("\\.(?:jpg|jpeg|png|gif|woff2|woff|js|css|html|ico)$")
	return http.ListenAndServe(address, serverHandler)
}
// go run /home/psycho/go/src/web/server/server.go
