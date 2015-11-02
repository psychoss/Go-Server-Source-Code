package server
import (
	"bytes"
	"fmt"
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
	return map[string]string{"Language": strings.Split(i.Href, "-")[0], "Mark": util.String.Substr(i.Href, 0, 1)}, nil
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
    }else{
    return nil
    }
}
func init() {
	filterRxp = rubex.MustCompile("^(?:css|rethinkdb|go|mongo|node|rust|python)\\-")
}
// Handle the home page request.
func indexHandler(w http.ResponseWriter, req *http.Request) {
	layout, err := template.ParseFile(PUBLIC_PATH + LAYOUT_TEMPLATE_FILE)
	if err != nil {
		http.Error(w, TEMPLATE_FILE_NOT_FOUND, http.StatusNotFound)
		return
	}
	index, err := template.ParseFile(PUBLIC_PATH + INDEX_TEMPLATE_FILE)
	//artical, err := template.ParseFile(PUBLIC_PATH + ARTICAL_TEMPLATE_FILE)
	if err != nil {
		http.Error(w, TEMPLATE_FILE_NOT_FOUND, http.StatusNotFound)
		return
	}
	mapOutput := map[string]interface{}{"Title": TITLE, "Keyword": KEYWORD, "Description": DESCRIPTION, "Base": BASE_URL, "Url": BASE_URL, "Carousel": getCarousel(), "Script": getScript(),"Items":leveldb.GetRandomContents(20, &Filter{})}
	content := index.RenderInLayout(layout, mapOutput)
	
	w.Write([]byte(content))
}
  
func getCarousel() string {
	key := []byte(PREFIX_INDEX + util.Time.NowShortDateTime())
	bs, err := leveldb.Get(key)
    if err != nil || len(bs) == 0 {
		key = []byte(PREFIX_INDEX + "default")
		bs, _ = leveldb.Get(key)
	}
	fmt.Println(len(bs))
	return string(bs)
}
func getScript() string {
	key := []byte(PREFIX_SCRIPT + util.Time.NowShortDateTime())
	bs, err := leveldb.Get(key)
	if err != nil || len(bs) == 0 {
		key = []byte(PREFIX_SCRIPT + "default")
		bs, _ = leveldb.Get(key)
	}
	fmt.Println(len(bs))
	return string(bs)
}
// go run /home/psycho/go/src/web/server/index.go
