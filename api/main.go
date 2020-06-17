package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/Domain", Retrieve)
	router.GET("/test", requestURL)

	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
}
