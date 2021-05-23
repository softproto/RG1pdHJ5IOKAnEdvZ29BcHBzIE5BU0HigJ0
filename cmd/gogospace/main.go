package main

import (
	"fmt"
	"gogoapps-nasa/internal/urlcollector"
)

const apiKey = "DEMO_KEY"
const port = "8080"

func main() {

	fmt.Println("main()")
	urlcollector.RunServer(apiKey, port)

	fmt.Println("main() end")
}
