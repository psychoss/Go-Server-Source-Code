package server
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)
func parseRequestJSON(req *http.Request) (*interface{}, error) {
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(ba, &v)
	if err != nil {
		return nil, err
	}
	return &v,nil
}
// go run /home/psycho/go/src/web/server/parse.go
