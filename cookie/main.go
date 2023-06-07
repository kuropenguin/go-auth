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
	ID   string `json:"id"`
	Pass string `json:"pass"`
}

// curl -X POST -H "Content-Type: application/json" -d '{"id":"id" ,"pass":"pass"}' localhost:8080/login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// cookieの設定
	var loginForm LoginForm
	if err := json.NewDecoder(r.Body).Decode(&loginForm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !authentication(loginForm.ID, loginForm.Pass) {
		http.Error(w, "auth failed", http.StatusBadRequest)
		return
	}
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	setCookie := http.Cookie{Name: "sid", Value: "hogehoge", Expires: expiration}
	http.SetCookie(w, &setCookie)

	redirectPath := "/"
	cookie, err := r.Cookie("original_url")
	if err != nil {
		// cookieがない場合はトップに飛ばす
		http.Redirect(w, r, redirectPath, http.StatusSeeOther)
		return
	}
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
	sid, err := r.Cookie("sid")
	if err != nil {
		// Set cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "original_url",
			Value: r.URL.Path,
		})
		// Redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if sid.Value != "hogehoge" {
		// Set cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "original_url",
			Value: r.URL.Path,
		})
		// Redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	w.Write([]byte("This is a private area, welcome!\n"))
}
