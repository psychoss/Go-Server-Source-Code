package server
import (
	"net/http"
)
func getRequiredDatas(req *http.Request) ([]byte, []byte, error) {
	v, err := parseRequestJSON(req)
	if err != nil {
		return nil, nil, err
	}
	m := (*v).(map[string]interface{})
	return []byte(m["id"].(string)), []byte(m["password"].(string)),nil
}
// For the user sign in.
func signinHandler(w http.ResponseWriter, req *http.Request) {
   /* id,password,err:=getRequiredDatas(req)
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}*/
	/*lbs, err := leveldb.Get([]byte(KEY_ADMIN_CERTIFICATION_TOKEN))
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}*/

}
// For the user sign up.
func signupHandler(w http.ResponseWriter, req *http.Request) {
}
// go run /home/psycho/go/src/web/server/user.go
