package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString() string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var (
	store map[string]string
)

func main() {
	store = map[string]string{
		"twitter": "http://twitter.com",
		"yahoo":   "http://yahoo.com",
	}
	http.HandleFunc("/lookup/", redirect)
	http.HandleFunc("/add/", addlink)
	log.Println("Server started on port 3000")
	fmt.Println("POST /lookup/?url=YOUR_STRING_HERE")
	fmt.Println("POST /add/?url=YOUR_STRING_HERE")
	//str := RandString()
	//fmt.Println(str)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func addlink(w http.ResponseWriter, r *http.Request) {
	req := r.URL.Query().Get("url")
	short := RandString()
	var req2 string
	if !strings.HasPrefix(req, "http://") {
		req2 = "http://" + req
	}
	store[short] = req2
	log.Println(req)
	log.Println(short)

}

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Println("store:", store)
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
