package main

import (
	"fmt"
	"os"

	"github.com/davidlu1997/fast-shortener/config"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("Alive at %s%s\n", config.Server.BaseUrl, config.Server.Port)
}
