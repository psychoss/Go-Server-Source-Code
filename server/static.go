package server
import (
	"io/ioutil"
	"mime"
    "fmt"
	"net/http"
	"path/filepath"
)
func staticHandler(w http.ResponseWriter, url string) {
	output, err := ioutil.ReadFile(STATIC_PATH + url)
	if err != nil {
		http.Error(w, TEMPLATE_FILE_NOT_FOUND, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type",mime.TypeByExtension(filepath.Ext(url)))
    w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", 60*60*24))
	w.Write(output)
}
// go run /home/psycho/go/src/web/server/static.go
