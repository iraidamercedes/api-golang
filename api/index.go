package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}
