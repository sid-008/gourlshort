package main

import (
	"log"
	"net/http"
)

var (
	store map[string]string
)

func main() {
	store = map[string]string{
		"twitter": "http://twitter.com",
		"yahoo":   "http://yahoo.com",
	}
	http.HandleFunc("/lookup/", redirect)
	log.Println("Server started on port 3000")
	http.ListenAndServe(":3000", nil)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Query().Get("url")
	//res := strings.Split(path, "/")
	for key, value := range store {
		if key == req {
			http.Redirect(w, r, value, 307)
		}
	}
	log.Println(req)

	//http.Redirect(w, r, "http://google.com", 307)
}
