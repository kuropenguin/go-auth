package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/secret", secretHandler)
	http.ListenAndServe(":8080", nil)
}

type LoginForm struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// cookieの設定
	var loginForm LoginForm
	if err := json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !authentication(loginForm.Name, loginForm.Pass) {
		http.Error(w, "auth failed", http.StatusBadRequest)
		return
	}
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	setCookie := http.Cookie{Name: "cookie-name-hoge", Value: "cookie-test", Expires: expiration}
	http.SetCookie(w, &setCookie)

	redirectPath := "/"
	cookie, _ := r.Cookie("original_url")
	if cookie.Value != "" {
		redirectPath = cookie.Value
	}

	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

func authentication(id string, pass string) bool {
	if id == "id" && pass == "pass" {
		return true
	}
	return false
}

func secretHandler(w http.ResponseWriter, r *http.Request) {

	// IF NOT LOGIN
	// http.SetCookie(w, &http.Cookie{
	// Name:  "original_url",
	// Value: r.URL.Path,
	// })
	// Redirect to the login page
	// http.Redirect(w, r, "/login", http.StatusSeeOther)

	// // Check if count exists
	// if session.Values["count"] != nil {
	// 	session.Values["count"] = session.Values["count"].(int) + 1
	// } else {
	// 	session.Values["count"] = 1
	// }

	// // Save the session before writing to the response.
	// session.Save(r, w)

	// // Write response
	// count := session.Values["count"].(int)
	// fmt.Fprintf(w, "Count: %d", count)
}
