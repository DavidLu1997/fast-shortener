package main

import (
	"fmt"
	"os"

	"github.com/davidlu1997/fast-shortener/api"
	"github.com/davidlu1997/fast-shortener/config"
	"github.com/valyala/fasthttp"
)

const configPath = "config/config.json"

func main() {
	config, err := config.GetConfig(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	api := api.InitAPI(config)

	fmt.Printf("Serving %s%s\n", config.Server.BaseUrl, config.Server.Port)
	err = fasthttp.ListenAndServe(config.Server.Port, api.RequestHandler)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
