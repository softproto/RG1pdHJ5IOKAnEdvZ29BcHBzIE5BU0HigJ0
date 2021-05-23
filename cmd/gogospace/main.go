package main

import (
	"fmt"
	"gogoapps-nasa/internal/urlcollector"
	"os"
)

// const apiKey = "DEMO_KEY"
// const port = "8080"

func main() {

	// os.Setenv("API_KEY", "DEMO_KEY")
	// os.Setenv("PORT", "8080")
	//  fmt.Println("FOO:", os.Getenv("FOO"))

	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		apiKey = "KEY"
	}
	fmt.Println(apiKey)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "80"
	}
	fmt.Println(port)

	fmt.Println("main()")
	urlcollector.RunServer(apiKey, port)

	fmt.Println("main() end")

}
