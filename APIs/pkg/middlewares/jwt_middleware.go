package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWTSecretKey = ""
	JWTExp       = 0
)

func SetJWTSecretKey(value string) {
	JWTSecretKey = value
}

func SetJWTExp(value int) {
	JWTExp = value
}

func CreateJWTToken(payload map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range payload {
		claims[key] = value
	}

	exp := time.Now().Add(time.Duration(JWTExp) * time.Second).Unix()

	claims["exp"] = exp
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func JWTVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header is required", http.StatusBadRequest)
			return
		}
		token = strings.Replace(token, "Bearer", "", 1)
		token = strings.Trim(token, " ")

		verifyToken := func(token *jwt.Token) (interface{}, error) {
			return []byte(JWTSecretKey), nil
		}

		_, err := jwt.Parse(token, verifyToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
