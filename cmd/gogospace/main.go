package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
        Timeout: 3 * time.Second,
    } 	

	req, err := http.NewRequest(
		"GET", "https://tut.by", nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// добавляем заголовки
	// req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	// req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
