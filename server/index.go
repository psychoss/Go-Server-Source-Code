package server
import (
	"bytes"
	//"fmt"
	"net/http"
	//"web/level"
	"strings"
	"web/rubex"
	"web/template"
	"web/util"
)
// the regexp for check the key
var (
	filterRxp *rubex.Regexp
)
// the item of the home page
type Item struct {
	Href, SubTitle, SubDescription string
}
// the extension function use by the template engine
func (i *Item) BuildLanguage() (map[string]string, error) {
	return map[string]string{"Language": strings.Split(i.Href, "-")[0], "Mark": util.Substr(i.Href, 0, 1)}, nil
}
// the struct to filter and process the datas return by the leveldb
type Filter struct {
}
func (*Filter) Accept(key []byte) bool {
	return filterRxp.Match(key)
}
func (*Filter) Process(key, value []byte) interface{} {
	separated := bytes.Split(value, []byte(DELIMITER))
	if len(separated) > 2 {
		return &Item{string(key), string(separated[0]), string(separated[2])}
	} else {
		return nil
	}
}
func init() {
	filterRxp = rubex.MustCompile("^(?:css|rethinkdb|go|mongo|node|rust|python)\\-")
}
// Handle the home page request.
func indexHandler(w http.ResponseWriter, req *http.Request) {

	layout, err := template.ParseFile(PATH_PUBLIC + TEMPLATE_LAYOUT)
	if err != nil {
		http.Error(w, ERROR_TEMPLATE_NOT_FOUND, http.StatusNotFound)
		return
	}
	index, err := template.ParseFile(PATH_PUBLIC + TEMPLATE_INDEX)
	//artical, err := template.ParseFile(PATH_PUBLIC + TEMPLATE_ARTICAL)
	if err != nil {
		http.Error(w, ERROR_TEMPLATE_NOT_FOUND, http.StatusNotFound)
		return
	}
	mapOutput := map[string]interface{}{"Title": TITLE, "Keyword": KEYWORD, "Description": DESCRIPTION, "Base": BASE_URL, "Url": BASE_URL, "Carousel": getAddition(PREFIX_INDEX), "Script": getAddition(PREFIX_SCRIPT), "Items": leveldb.GetRandomContents(20, &Filter{})}
	content := []byte(index.RenderInLayout(layout, mapOutput))
	w.Write(content)
	go cacheFile(PATH_CACHE+"index"+EXTENSION_HTML, content)
}
func getAddition(prefix string) string {
	key := []byte(prefix + util.NowShortDateTime())
	bs, err := leveldb.Get(key)
	if err != nil || len(bs) == 0 {
		key = []byte(prefix + "default")
		bs, _ = leveldb.Get(key)
	}
	return string(bs)
}
// go run /home/psycho/go/src/web/server/index.go
