package apod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const servicePath = "https://api.nasa.gov/planetary/apod"

type ApodResponce struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

type ApodError struct {
	Code           string `json:"code"`
	Message        string `json:"msg"`
	ServiceVersion string `json:"service_version"`
}

func UnmarshallApodResponce(responce []byte) (resp *ApodResponce, err error) {
	err = json.Unmarshal(responce, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func UnmarshallApodError(responce []byte) (resp *ApodError, err error) {
	err = json.Unmarshal(responce, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func SendRequest(apiKey, date string, out chan string) {

	//fmt.Println("SendRequest()")

	client := &http.Client{}

	req, err := http.NewRequest("GET", servicePath, nil)
	if err != nil {
		log.Println(err)
		return
	}

	q := req.URL.Query()
	q.Add("api_key", apiKey)
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()
	// fmt.Println(req.URL.RawQuery)

	// req.Header.Add("Accept", "text/html")
	// req.Header.Add("User-Agent", "MSIE/15.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//http.Error(w, err.Error(), 500)
		return
	}

	s := resp.StatusCode

	if s == 200 {
		a, err := UnmarshallApodResponce(b)
		if err != nil {
			log.Println(err)
			return
		}
		// fmt.Println(a.Title)
	//	fmt.Println(a.Url)
		out <- a.Url
	} else {
		// fmt.Println(resp.StatusCode)
		// fmt.Println(http.StatusText(resp.StatusCode))
		fmt.Println(resp.Status)
	}

}
