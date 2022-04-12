package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("<h1>Welcome to my web server!</h1>"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is the about page</h1>"))
}

func redirect(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, 301)
	}
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "The id query parameter is missing", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "<h1>The user id is: %s</h1>", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/example", redirect("https://www.securemeeting.org/"))
	// /user?id=123 <-- search this way
	mux.HandleFunc("/user", userHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
