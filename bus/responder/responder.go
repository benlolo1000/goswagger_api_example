package responder

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type CustomResponder struct {
	email     *string
	responder middleware.Responder
}

type Claims struct {
	Email *string `json:"email"`
	jwt.StandardClaims
}

func CookieHandler(responder middleware.Responder, email *string) *CustomResponder {
	return &CustomResponder{
		email:     email,
		responder: responder,
	}
}

func CreateToken(Email *string) []string {
	jwtKey := []byte("my_secret_key")
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Email: Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil
	}
	splitToken := strings.SplitN(tokenString, ".", 3)
	return splitToken
}

func (this *CustomResponder) WriteResponse(rw http.ResponseWriter, p runtime.Producer) {
	token := CreateToken(this.email)
	expirationTime := time.Now().Add(60 * time.Minute)

	cookie := http.Cookie{
		Name:     "token",
		Value:    token[0] + "." + token[1],
		Expires:  expirationTime,
		HttpOnly: true}
	http.SetCookie(rw, &cookie)
	this.responder.WriteResponse(rw, p)
}
