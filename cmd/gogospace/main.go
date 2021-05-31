package main

import (
	"fmt"
	"gogoapps-nasa/internal/urlcollector"
)

func main() {
	fmt.Println("GoGoSpace")

	config := getConfigFromEnv()
	fmt.Println(*config)

	urlcollector.RunServer(config)
}

