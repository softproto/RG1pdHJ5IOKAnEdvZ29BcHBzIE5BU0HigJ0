package main

import (
	"gogoapps-nasa/internal/urlcollector"
	"os"
	"strconv"
)

func getStringFromEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getIntFromEnv(key string, defaultValue int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}

	return defaultValue
}

//Setting up the app using Env variables
func getConfigFromEnv() *urlcollector.Config {
	c := urlcollector.Config{}

	apiKey := getStringFromEnv("API_KEY", "DEMO_KEY")
	port := getStringFromEnv("PORT", "8080")
	url := getStringFromEnv("APOD_URL", "https://api.nasa.gov/planetary/apod")
	concurrentRequests := getIntFromEnv("CONCURRENT_REQUESTS", 5)
	transportTimeout := getIntFromEnv("TRANSPORT_TIMEOUT", 5)
	handshakeTimeout := getIntFromEnv("HANDSGAKE_TIMEOUT", 5)
	clientTimeout := getIntFromEnv("CLIENT_TIMEOUT", 10)

	c.Setup(apiKey, port, url, concurrentRequests, transportTimeout, handshakeTimeout, clientTimeout)

	return &c
}
