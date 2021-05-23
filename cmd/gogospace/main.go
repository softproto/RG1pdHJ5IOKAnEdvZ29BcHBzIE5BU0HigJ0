package main

import (
	"fmt"
	"gogoapps-nasa/internal/urlcollector"
)

const apiKey = "DEMO_KEY"

func main() {

	fmt.Println("main()")
	urlcollector.RunServer(apiKey)

	fmt.Println("main() end")
}
