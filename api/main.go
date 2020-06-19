package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/test", Test)
	router.GET("/domain", RequestURL)

	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
}
