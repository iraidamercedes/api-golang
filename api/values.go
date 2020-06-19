package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
	"github.com/valyala/fasthttp"
)

//define strucs
type endpoint struct {
	IpAddress string `json:"IpAddress"`
	Grade     string `json:"Grade"`
}

type Server struct {
	Endpoints         []endpoint `json:"Endpoints"`
	Country           string     `json:"Country"`
	Owner             string     `json:"Owner"`
	Serverschanged    bool       `json:"server_changed"`
	Sslgrade          string     `json:"ssl_grade"`
	GradeTrustIgnored string     `json:"previous_ssl_grade"`
	Logo              string     `json:"logo"`
	Title             string     `json:"title"`
	IsDown            bool       `json:"is_down"`
	Status            string     `json:"status, omitempty"`
}

//create the request
func RequestURL(ctx *fasthttp.RequestCtx) {

	//bring the data from domain
	req, _ := http.NewRequest("GET", "https://api.ssllabs.com/api/v3/analyze?host=sibcolombia.net", nil)

	client := &http.Client{}
	resp, _ := client.Do(req)
	servers, _ := ioutil.ReadAll(resp.Body)

	var info Server
	if err := json.Unmarshal([]byte(servers), &info); err != nil {
		log.Fatal("Failed to generate json", err)
	}

	//bring the data from whois
	whois_raw, err := whois.Whois("sibcolombia.net")
	result, err := whoisparser.Parse(whois_raw)

	//create values
	Endpoints := info.Endpoints
	Country := result.Registrant.Country
	Owner := result.Administrative.Organization
	Sslgrade := "B"
	GradeTrustIgnored := "A+"
	Title := "SiB Colombia"
	IsDown := strings.ToLower(info.Status) != "ready"

	//procces dates
	layout := "2006-01-02"
	yesterday := time.Now().AddDate(0, 0, -1)
	input := result.Domain.UpdatedDate
	t, _ := time.Parse(layout, input)

	Serverschanged := (t).After(yesterday)

	//structure the API response
	myData := Server{
		Endpoints:         Endpoints,
		Country:           Country,
		Owner:             Owner,
		Serverschanged:    Serverschanged,
		Sslgrade:          Sslgrade,
		GradeTrustIgnored: GradeTrustIgnored,
		Title:             Title,
		IsDown:            IsDown,
	}

	//print the json
	prettyJSON, err := json.MarshalIndent(myData, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Fprintf(ctx, "%s\n", string(prettyJSON))

}
