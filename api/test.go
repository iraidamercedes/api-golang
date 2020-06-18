package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

type endpoint struct {
	IpAddress string
	Grade     string
}

type Server struct {
	Host      string
	Status    string
	Endpoints []endpoint
}

func requestURL(ctx *fasthttp.RequestCtx) {

	req, _ := http.NewRequest("GET", "https://api.ssllabs.com/api/v3/analyze?host=sibcolombia.net", nil)

	client := &http.Client{}
	resp, _ := client.Do(req)
	servers, _ := ioutil.ReadAll(resp.Body)
	//println(string(servers))
	fmt.Fprintf(ctx, "%s", string(servers))

	var info Server
	if err := json.Unmarshal([]byte(servers), &info); err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Printf("Host: %s, Status: %s, Endpoints: %+v", info.Host, info.Status, info.Endpoints)

}
