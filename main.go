package main

import (
	"fmt"
	"gogoapps-nasa/internal/urlcollector"
)

const apiKey = "DEMO_KEY"

func main() {
	out := make(chan string)
	urlcollector.RunCollector(apiKey, out)

	fmt.Println("main()")

	// urls := []string{}
	for url := range out {
		fmt.Println(url)
	}

	//	fmt.Println(<-out)
	//urlcollector.RunServer()
}
