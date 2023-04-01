package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
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

type UrlStruct struct { // NOTE: struct fields must start with upper case for json package to see its value
	Url   string //`json:"url"`
	Short string //`json:"short"`
}

func main() {
	store = map[string]string{
		"twitter": "http://twitter.com",
		"yahoo":   "http://yahoo.com",
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/view", func(c *gin.Context) {
		c.JSON(200, store)
	})

	r.GET("/:short", func(c *gin.Context) {
		req := c.Param("short")
		log.Println("Store:", store)

		for key, value := range store {
			if key == req {
				c.Redirect(301, value)
			}
		}
		log.Println(req)
	})

	r.POST("/addlink", func(c *gin.Context) {
		var reqBody UrlStruct
		if err := c.BindJSON(&reqBody); err != nil {
			log.Fatal(err)
		}
		requrl := reqBody.Url
		if !strings.HasPrefix(requrl, "http://") {
			requrl = "http://" + requrl
		}
		short := RandString()
		store[short] = requrl
		fmt.Println(short)

	})

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
