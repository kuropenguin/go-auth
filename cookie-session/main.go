package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

func main() {
	http.HandleFunc("/set_cookie", setCookieHandler)
	http.HandleFunc("/check_cookies", checkCookiesHandler)
	http.HandleFunc("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}

// ただ cookie を設定するだけ
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// cookieの設定
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "cookie-name-hoge", Value: "cookie-test", Expires: expiration}
	http.SetCookie(w, &cookie)

	w.Write([]byte("Set Cookie\n"))
}

func checkCookiesHandler(w http.ResponseWriter, r *http.Request) {
	// クライアントからきたリクエストに埋め込まれているcookieの確認
	for _, c := range r.Cookies() {
		log.Print("Name:", c.Name, "Value:", c.Value)
	}
	w.Write([]byte("check Cookie\n"))
}

// Initialize a session store with a secret key used for authentication.
// Replace the string "something-very-secret" with your own secret.
var store = sessions.NewCookieStore([]byte("secret-hoge"))

func countHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve or initialize session
	session, _ := store.Get(r, "session-name-hoge")

	// Check if count exists
	if session.Values["count"] != nil {
		session.Values["count"] = session.Values["count"].(int) + 1
	} else {
		session.Values["count"] = 1
	}

	// Save the session before writing to the response.
	session.Save(r, w)

	// Write response
	count := session.Values["count"].(int)
	fmt.Fprintf(w, "Count: %d", count)
}
