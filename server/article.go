package server
import (
	"bytes"
	"encoding/json"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"web/template"
	//"fmt"
)
func convertToHTML(input []byte) string {
	return string(blackfriday.MarkdownCommon(input))
}
func fecthHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	layout, err := template.ParseFile(PATH_PUBLIC + TEMPLATE_LAYOUT)
	if err != nil {
		http.Error(w, ERROR_TEMPLATE_NOT_FOUND, http.StatusNotFound)
		return
	}
	artical, err := template.ParseFile(PATH_PUBLIC + TEMPLATE_ARTICAL)
	if err != nil {
		http.Error(w, ERROR_TEMPLATE_NOT_FOUND, http.StatusNotFound)
		return
	}
	key := bytes.Trim([]byte(url), "/")
	bs, err := leveldb.Get(key)
	if err != nil {
		http.Error(w, ERROR_TEMPLATE_NOT_FOUND, http.StatusNotFound)
		return
	}
	sps := bytes.Split(bs, []byte(DELIMITER))
	if len(sps) > 3 {
		page := struct {
			Title, Keyword, Description, Markdown, Base, Url string
		}{
			Title:       string(sps[0]) + TITLE,
			Keyword:     string(sps[1]),
			Description: string(sps[2]),
			Markdown:    convertToHTML(bytes.Replace(sps[3], []byte("{{Base}}"), []byte(BASE_URL), -1)),
			Base:        BASE_URL,
			Url:         BASE_URL + string(key) + "/",
		}
		content := []byte(artical.RenderInLayout(layout, &page))
        go cacheFile(PATH_CACHE +string(key) +EXTENSION_HTML,content)
		w.Write(content)
	} else {
		http.Error(w, ERROR_TEMPLATE_NOT_FOUND, http.StatusNotFound)
		return
	}
}
func getHandler(w http.ResponseWriter, req *http.Request) {
	crossDomain(w)
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	var v interface{}
	err = json.Unmarshal(bs, &v)
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	m := v.(map[string]interface{})
	href, token := []byte(m["Id"].(string)), []byte(m["token"].(string))
	lbs, err := leveldb.Get([]byte(KEY_ADMIN_CERTIFICATION_TOKEN))
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	if bytes.Equal(token, lbs) {
		bs, err = leveldb.Get(href)
		if err != nil {
			http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
			return
		}
		w.Write(bs)
	} else {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
}
func updateHandler(w http.ResponseWriter, req *http.Request) {
	crossDomain(w)
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	var v interface{}
	err = json.Unmarshal(bs, &v)
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	m := v.(map[string]interface{})
	href, content, token := []byte(m["href"].(string)), []byte(m["content"].(string)), []byte(m["token"].(string))
	lbs, err := leveldb.Get([]byte(KEY_ADMIN_CERTIFICATION_TOKEN))
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	if bytes.Equal(token, lbs) {
		err = leveldb.Put(href, content)
		if err != nil {
			http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
}
// go run /home/psycho/go/src/web/server/article.go
