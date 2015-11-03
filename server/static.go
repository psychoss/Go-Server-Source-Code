package server
import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
)
// Actually,the front-end Web server,nginx will be serve the static files.
func staticHandler(w http.ResponseWriter, url string) {
	output, err := ioutil.ReadFile(PATH_STATIC + url)
	if err != nil {
		http.Error(w, "File not found!", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(url)))
	w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", 60*60*24))
	w.Write(output)
}
// go run /home/psycho/go/src/web/server/static.go
