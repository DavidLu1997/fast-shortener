package main

import (
	"fmt"
	"os"

	"github.com/davidlu1997/fast-shortener/api"
	"github.com/davidlu1997/fast-shortener/config"
	"github.com/valyala/fasthttp"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	api := api.InitAPI()
	err = fasthttp.ListenAndServe(config.Server.Port, api.RequestHandler)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
