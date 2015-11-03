package jwt
import (
	"github.com/dgrijalva/jwt-go"
)
func NewToken(username, secret string) (string, error) {
	jt := jwt.New(jwt.SigningMethodHS256)
	jt.Claims["username"] = username
	token, err := jt.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
func CheckToken(tokenstr, secret string) (string, bool, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", false, nil
	}
	if token.Valid {
		return token.Claims["username"].(string), true, nil
	}
	return "", false, nil
}
// go run /home/psycho/go/src/web/jwt/jwt.go
