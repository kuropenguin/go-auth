// reference :  https://iketechblog.com/go-jwt/
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kuropenguin/go-auth/beare/auth"

	"github.com/gorilla/mux"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {
	r := mux.NewRouter()

	r.Handle("/public", public)
	// curl localhost:8080/private -H "Authorization:Bearer token"
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "Google",
		Tag:   "search engine",
		URL:   "https://www.google.com",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "Private Title",
		Tag:   "Private Tag",
		URL:   "https://www.google.com",
	}
	json.NewEncoder(w).Encode(post)
})
