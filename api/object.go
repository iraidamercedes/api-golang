package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type servers struct {
	Address  string `json:"adress"`
	Sslgrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}

type Domain struct {
	Serverlist       []servers `json:"Servers"`
	Serverschanged   bool      `json:"server_changed"`
	Sslgrade         string    `json:"ssl_grade"`
	Previoussslgrade string    `json:"previous_ssl_grade"`
	Logo             string    `json:"logo"`
	Title            string    `json:"title"`
	IsDown           bool      `json:"is_down"`
}

func Retrieve(ctx *fasthttp.RequestCtx) {

	myDomain := Domain{
		Serverlist: []servers{
			servers{"51.38.239.240", "B", "US", "Privacy Protect, LLC (PrivacyProtect.org)"},
			servers{"51.38.239.240", "A+", "US", "Privacy Protect, LLC (PrivacyProtect.org)"},
		},
		Serverschanged:   true,
		Sslgrade:         "B",
		Previoussslgrade: "A+",
		Logo:             "https://sibcolombia.net/wp-content/uploads/2018/06/logo-sib.svg",
		Title:            "SiB Colombia",
		IsDown:           false,
	}

	prettyJSON, err := json.MarshalIndent(myDomain, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}

	fmt.Fprintf(ctx, "%s\n", string(prettyJSON))
}
