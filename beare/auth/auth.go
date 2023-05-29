package auth

import (
	"log"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256) // header

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "545465"
	claims["name"] = "taro"
	claims["iat"] = time.Now().Unix() //issued at
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte("secret_key_hogehoge")) // 電子署名
	log.Println(tokenString)
	w.Write([]byte(tokenString))
})

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key_hogehoge"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
