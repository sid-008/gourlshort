package main

import "net/http"

func main() {
	//	store := make(map[string]string)

	http.HandleFunc("/google", redirect)
	http.ListenAndServe(":3000", nil)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://google.com", 301)
}
