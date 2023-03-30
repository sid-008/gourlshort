package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/helth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "All good!",
		})
	})

	r.GET("/:short", redirect)

	r.POST("/addlink/:url", addlink)

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
	//http.HandleFunc("/lookup/:url", redirect)
	//http.HandleFunc("/add/", addlink)
	//log.Println("Server started on port 3000")
	//fmt.Println("GET /lookup/?url=YOUR_STRING_HERE")
	//fmt.Println("POST /add/?url=YOUR_STRING_HERE")
	//str := RandString()
	//fmt.Println(str)
	//	err := http.ListenAndServe(":3000", nil)
	//	if err != nil {
	//		log.Fatal(err)
	//}
}

func addlink(c *gin.Context) {
	url := c.Param("url")
	short := RandString()
	/*if !strings.HasPrefix(req, "http://") {
		req2 = "http://" + req
	}*/
	store[short] = url
	log.Println(url)
	log.Println(short)
	log.Println(store)
	c.JSON(200, url)
}

func redirect(c *gin.Context) {
	log.Println("store:", store)
	req := c.Param("short")

	for key, value := range store {
		if key == req {
			c.Redirect(301, value)
		}
	}
	log.Println(req)
}
