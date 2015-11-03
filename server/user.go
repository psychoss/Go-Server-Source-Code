package server
import (
	"bytes"
	"fmt"
	"net/http"
	"web/jwt"
)
// parse the datas from the request body
func getRequiredDatas(req *http.Request) ([]byte, []byte, error) {
	v, err := parseRequestJSON(req)
	if err != nil {
		return nil, nil, err
	}
	m := (*v).(map[string]interface{})
	return []byte(m["id"].(string)), []byte(m["password"].(string)), nil
}
// For the user sign in.
func signinHandler(w http.ResponseWriter, req *http.Request) {
	id, password, err := getRequiredDatas(req)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	id = append([]byte(PREFIX_USER), id[0:]...)
	lbs, err := leveldb.Get(id)
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	if bytes.Equal(lbs, password) {
		fmt.Println(string(id[len([]byte(PREFIX_USER)):]))
		token, err := jwt.NewToken(string(id[len([]byte(PREFIX_USER)):]), string(secret))
		if err != nil {
			http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
	}
}
// For the user sign up.
func signupHandler(w http.ResponseWriter, req *http.Request) {
	// Parse the username and password from the request body
	id, password, err := getRequiredDatas(req)
	defer req.Body.Close()
	// if error is not a nil
	// end the function immediately after sending the response
	if err != nil {
		http.Error(w, ERROR_BAD_REQUEST, http.StatusBadRequest)
		return
	}
	// append the prefix
	id = append([]byte(PREFIX_USER), id[0:]...)
	// check if the username have been used
	lbs, err := leveldb.Get(id)
	if err != nil {
		http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
		return
	}
	// if the username is not used
	if len(lbs) == 0 {
		err := leveldb.Put(id, password)
		if err != nil {
			http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
			return
		}
		// make the token for the response
		// the token only include the username and the hash which for use by the validator
		token, err := jwt.NewToken(string(id[len([]byte(PREFIX_USER)):]), string(secret))
		if err != nil {
			http.Error(w, ERROR_SERVER_INTERNAL, http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
	}
}
// go run /home/psycho/go/src/web/server/user.go
