/*
a smiple web server for http://www.mean101.com
backend for the nginx
*/
package server
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"web/level"
	"web/log"
	"web/rubex"
)
var (
	// regular expression for match the request for  articles
	rxp *rubex.Regexp
	// regular expression for match the request for static files
	staticRxp *rubex.Regexp
	// the wrapper for leveldb
	leveldb *level.Level
	secret  []byte
)
// the holder for request handler
type RequestHandler struct {
}
// the main function for process the request
func (*RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// recover any possible panic
	// this defer function must place first in the sequence
	// because the nature of the defer
	// in-first then out-last
	// in other word，make sure this defer preform at the end of function
	defer func() {
		if err := recover(); err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}
	}()
	// in the future, may need to optimize the router
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
		if url == "/comment" {
			putCommentHandler(w, req)
		} else if url == "/comment/get" { /*for fectch the comments*/
			getCommentHandler(w, req)
		} else if url == "/update" {
			updateHandler(w, req)
		} else if url == "/post" {
			getHandler(w, req)
		} else if url == "/signin" {
			signinHandler(w, req)
		} else if url == "/signup" {
			signupHandler(w, req)
		}
	}
}
// Cache the file for nginx
func cacheFile(fileName string, datas []byte) {
	ioutil.WriteFile(PATH_CACHE+fileName+".html", datas, 0777)
}
func removeFile(fileName string) {
	os.Remove(PATH_CACHE + fileName + ".html")
	os.Remove(PATH_CACHE + "index.html")
}
// 允许跨域方法
// Allow cross-domain
func crossDomain(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
// Close the LevelDB
func Close() {
	leveldb.Close()
}
// Start the Server
func StartNewServer(address string) error {
	var err error
	serverHandler := &RequestHandler{}
	leveldb, err = level.New(PATH_DB)
	if err != nil {
		return err
	}
	
	rxp = rubex.MustCompile("^/(?:css|go|mongo|node|python|rethinkdb|rust)[\\-a-zA-Z0-9/]*$")
	staticRxp = rubex.MustCompile("\\.(?:jpg|jpeg|png|gif|woff2|woff|js|css|html|ico)$")
	secret, err = leveldb.Get([]byte(KEY_ADMIN_CERTIFICATION_TOKEN))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return http.ListenAndServe(address, serverHandler)
}
// go run /home/psycho/go/src/web/server/server.go
