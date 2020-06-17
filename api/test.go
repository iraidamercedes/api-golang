package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

func requestURL(ctx *fasthttp.RequestCtx) {

	req, _ := http.NewRequest("GET", "https://api.ssllabs.com/api/v3/analyze?host=sibcolombia.net", nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	servers, _ := ioutil.ReadAll(resp.Body)
	//println(string(servers))

	fmt.Fprintf(ctx, "%s", string(servers))

	type Server struct {
		Host   string
		Status string
	}

	var info Server

	if err := json.Unmarshal([]byte(servers), &info); err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Printf("Host: %s, Status: %s", info.Host, info.Status)
}
