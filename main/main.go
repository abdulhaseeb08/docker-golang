package main

import (
	"fmt"
	"net/http"
)

//Handler Interface
type welcomeMessage string

func (wc welcomeMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello welcome to this simple webpage")
}

//Handler Functions
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the login page")
}

func funkyPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ":P")
}

func main() {
	router := http.NewServeMux()

	var wc welcomeMessage
	router.Handle("/", wc)
	router.HandleFunc("/login", login)
	router.HandleFunc("/funkyPage", funkyPage)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
