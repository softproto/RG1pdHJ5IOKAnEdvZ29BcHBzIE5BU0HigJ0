package main

import (
	"fmt"
	"gogoapps-nasa/internal/urlcollector"
	"os"
	"strconv"
)

func main() {
	fmt.Println("GoGoSpace")

	config := getConfigFromEnv()
	fmt.Println(*config)

	urlcollector.RunServer(config)
}


//Setting up the app using Env variables
func getConfigFromEnv() *urlcollector.Config {
	c := urlcollector.Config{}

	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		apiKey = "DEMO_KEY"
	}

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	url, exists := os.LookupEnv("APOD_URL")
	if !exists {
		url = "https://api.nasa.gov/planetary/apod"
	}

	r := os.Getenv("CONCURRENT_REQUESTS")
	concurrentRequests, err := strconv.Atoi(r)
	if err != nil {
		concurrentRequests = 5
	}

	tt := os.Getenv("TRANSPORT_TIMEOUT")
	transportTimeout, err := strconv.Atoi(tt)
	if err != nil {
		transportTimeout = 5
	}

	ht := os.Getenv("HANDSGAKE_TIMEOUT")
	handshakeTimeout, err := strconv.Atoi(ht)
	if err != nil {
		handshakeTimeout = 5
	}

	ct := os.Getenv("CLIENT_TIMEOUT")
	clientTimeout, err := strconv.Atoi(ct)
	if err != nil {
		clientTimeout = 10
	}

	c.Setup(apiKey, port, url, concurrentRequests, transportTimeout, handshakeTimeout, clientTimeout)

	return &c
}
