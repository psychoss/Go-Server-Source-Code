package server
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"web/jwt"
	"web/util"
)
// Handle the request for the comments.
func getCommentHandler(w http.ResponseWriter, req *http.Request) {
	bs, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	//	if err != nil {
	//		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
	//		return
	//	}
	bs = leveldb.GetAllKeysByAnchor(util.Union([]byte(PREFIX_COMMENT), bs))
	//fmt.Println(string(bs))
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
	// extract the datas for later
	// be careful，if the value of m["href"] is nil
	// type assert to string will be cause a panic
	// if not recover from this
	// the whole application will be stop
    fmt.Println(m)
	username, href, token, content := m["username"].(string), m["href"].(string), m["token"].(string), m["content"].(string)
	// check the token
	valide, err := jwt.CheckToken(token, string(secret))
	// check the error
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	// check the result return by the validator
	if !valide {
		http.Error(w, ERROR_NOT_PERMISSION, http.StatusNotAcceptable)
		return
	}
	// use the microsecond as the suffix
	time := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	// the refid is a reference to the comment which the user reply to
	refid := m["refid"]
	// the holde for the data which will be used to insert into the leveldb
	var bs []byte
	id := PREFIX_COMMENT + href + time
	if refid == nil {
		bs, err = json.Marshal(map[string]string{"username": username, "id": id, "created": time, "content": content})
	} else {
		bs, err = json.Marshal(map[string]string{"username": username, "id": id, "created": time, "content": content, "refid": refid.(string)})
	}
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	err = leveldb.Put([]byte(id), bs)
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	anchor := []byte(PREFIX_COMMENT + href)
	// The anchor is for find the range of the comments
	// Make sure the anchor have been inserted
	leveldb.Put(anchor, anchor)
	w.Write([]byte(id))
}
// go run /home/psycho/go/src/web/server/comment.go
