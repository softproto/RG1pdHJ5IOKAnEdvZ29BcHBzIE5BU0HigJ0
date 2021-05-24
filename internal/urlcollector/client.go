package urlcollector

import (
	"errors"
	"fmt"
	"gogoapps-nasa/internal/apod"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func sendRequest(config *Config, date string, cd *collectedData) {
	fmt.Println("...sendRequest()")
	defer fmt.Println("...sendRequest() done")

	client := customHTTPClient(config.transportTimeout, config.handshakeTimeout, config.clientTimeout)

	req, err := http.NewRequest("GET", apod.URL, nil)
	if err != nil {
		cd.collectError(date, err)
		return
	}

	q := req.URL.Query()
	q.Add("api_key", config.apiKey)
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		cd.collectError(date, err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cd.collectError(date, err)
		return
	}

	if resp.StatusCode == 200 {
		a, err := apod.UnmarshallResponse(b)
		if err != nil {
			cd.collectError(date, err)
			return
		}
		cd.collectURL(a.URL)
	} else {
		cd.collectError(date, errors.New(resp.Status))
	}
}

func customHTTPClient(transportTimeout time.Duration, handshakeTimeout time.Duration, clientTimeout time.Duration) (client *http.Client) {
	var transport = &http.Transport{
		Dial:                (&net.Dialer{Timeout: transportTimeout * time.Second}).Dial,
		TLSHandshakeTimeout: handshakeTimeout * time.Second,
	}
	client = &http.Client{
		Timeout:   time.Second * clientTimeout,
		Transport: transport,
	}
	return client
}
