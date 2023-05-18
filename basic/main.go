package main

import (
	"fmt"
	"net/http"
)

func BasicAuth(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("BasicAuth")
		user, pass, ok := r.BasicAuth()

		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are unauthorized to access the application.\n"))
			return
		}

		handler(w, r)
	}
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a private area, welcome!\n"))
}

func MyHandler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is another private area, welcome!\n"))
}

func main() {
	// curl --basic -u myusername1:mypassword1 http://localhost:8080/area1
	http.HandleFunc("/area1", BasicAuth(MyHandler, "myusername1", "mypassword1", "Please enter your username and password for area 1"))
	http.HandleFunc("/area2", BasicAuth(MyHandler2, "myusername2", "mypassword2", "Please enter your username and password for area 2"))
	http.ListenAndServe(":8080", nil)
}
