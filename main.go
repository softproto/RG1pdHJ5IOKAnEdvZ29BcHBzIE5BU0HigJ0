package main

import "gogoapps-nasa/internal/urlcollector"

const apiKey = "DEMO_KEY"

func main() {
	urlcollector.RunCollector(apiKey)
	//urlcollector.RunServer()
}
