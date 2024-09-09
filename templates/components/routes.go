package components

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

type State struct {
	Open bool
}

func Routes() {
	err := InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	} else {
		log.Println("Database initialized")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(Index()).ServeHTTP(w, r)
	})

	http.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			templ.Handler(LoginContent()).ServeHTTP(w, r)
		} else {
			templ.Handler(Login()).ServeHTTP(w, r)
		}
	})

	http.HandleFunc("GET /sign-up", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") == "true" {
			templ.Handler(SignUpContent()).ServeHTTP(w, r)
		} else {
			templ.Handler(SignUp()).ServeHTTP(w, r)
		}
	})

	http.HandleFunc("POST /db", func(w http.ResponseWriter, r *http.Request) {
		CreateUser("Johnson", "yU3y3@example.com")
		users, err := GetAllUsers()
		if err != nil {
			log.Fatalf("Failed to retrieve users: %v", err)
		}
		fmt.Println(users)
	})
	
	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ErrorPage()).ServeHTTP(w, r)
	})
}
