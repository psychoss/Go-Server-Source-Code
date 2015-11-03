package server
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"web/jwt"
)
// Handle the request for the comments.
func getCommentHandler(w http.ResponseWriter, req *http.Request) {
	bs, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	bs, err = leveldb.Get(bs)
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	w.Write(bs)
}
func putCommentHandler(w http.ResponseWriter, req *http.Request) {
	// Parse the json from the request body
	v, err := parseRequestJSON(req)
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	// dereference the pointer variable and type assertion
	m := (*v).(map[string]interface{})
	href, token, content := m["href"].(string), m["token"].(string), m["content"].(string)
	fmt.Println(refid)
	valide, err := jwt.CheckToken(token, string(secret))
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	if !valide {
		http.Error(w, ERROR_NOT_PERMISSION, http.StatusNotAcceptable)
		return
	}
	time := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	refid := m["refid"]
	var b []byte
	if refid == nil {
		bs, err = json.Marshal(map[string]string{"created": time, "content": content})
	} else {
		bs, err = json.Marshal(map[string]string{"created": time, "content": content, "refid": refid.(string)})
	}
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	key := []byte(PREFIX_COMMENT + href + time)
	err = leveldb.Put(key, bs)
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	anchor := []byte(PREFIX_COMMENT + href)
    leveldb.Put(anchor,anchor)
	w.Write(key)

}
// go run /home/psycho/go/src/web/server/comment.go
