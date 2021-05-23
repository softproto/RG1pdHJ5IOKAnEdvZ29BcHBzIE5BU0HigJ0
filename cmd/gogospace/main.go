package main

import (
	"fmt"
	"gogoapps-nasa/internal/urlcollector"
	"os"
)

const apiKey = "DEMO_KEY"
const port = "8080"

func main() {

	fmt.Println("main()")
	urlcollector.RunServer(apiKey, port)

	fmt.Println("main() end")

	os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
//	path, exists := os.LookupEnv("PATH")
}
