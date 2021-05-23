package urlcollector

import (
	"errors"
	"fmt"
	"gogoapps-nasa/internal/apod"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	dateLayout       = "2006-01-02"
	maxConnLimit     = 2
	transportTimeout = 5
	handshakeTimeout = 5
	clientTimeout    = 10
)

type collectedData struct {
	mutex  sync.Mutex
	urls   []string
	errors []string
}

func getDatesList(startDate, endDate string) (list []string, err error) {
	fmt.Println("getDatesList()")

	if startDate == "" || endDate == "" {
		return nil, errors.New("start_date and end_date must be provided")
	}

	sd, err := time.Parse(dateLayout, startDate)
	if err != nil {
		return nil, err
	}
	ed, err := time.Parse(dateLayout, endDate)
	if err != nil {
		return nil, err
	}

	if sd.Before(ed) || sd == ed {
		duration := int(ed.Sub(sd).Hours() / 24)
		for i := 0; i <= duration; i++ {
			date := sd.AddDate(0, 0, i)
			list = append(list, date.Format(dateLayout))
		}
		return list, nil
	}
	return nil, errors.New("start_date should be before or equal end_date")
}

func runCollector(apiKey, start_date, end_date string) *collectedData {
	fmt.Println("runCollector()")
	cd := collectedData{
		mutex:  sync.Mutex{},
		urls:   []string{},
		errors: []string{},
	}

	dates, err := getDatesList(start_date, end_date)
	if err != nil {
		collectError(fmt.Sprintf("getDatesList(%s, %s)", start_date, end_date), err, &cd)
		return &cd
	}

	d := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < maxConnLimit; i++ {
		wg.Add(1)
		go urlsFetcher(d, apiKey, &wg, &cd)
	}
	for _, date := range dates {
		d <- date
	}
	close(d)
	wg.Wait()

	// fmt.Println(cd.Urls)
	// fmt.Println(cd.Errors)
	return &cd
}

func urlsFetcher(d chan string, apiKey string, wg *sync.WaitGroup, cd *collectedData) {
	fmt.Println("urlsFetcher()")
	defer fmt.Println("urlsFetcher() done")

	defer wg.Done()

	for date := range d {
		SendRequest(apiKey, date, cd)
	}
}

func SendRequest(apiKey, date string, cd *collectedData) {
	fmt.Println("SendRequest()")
	defer fmt.Println("SendRequest() done")

	client := customHTTPClient(transportTimeout, handshakeTimeout, clientTimeout)

	req, err := http.NewRequest("GET", apod.URL, nil)
	if err != nil {
		collectError(date, err, cd)
		return
	}

	q := req.URL.Query()
	q.Add("api_key", apiKey)
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		collectError(date, err, cd)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		collectError(date, err, cd)
		return
	}

	if resp.StatusCode == 200 {
		a, err := apod.UnmarshallResponse(b)
		if err != nil {
			collectError(date, err, cd)
			return
		}
		cd.mutex.Lock()
		cd.urls = append(cd.urls, a.Url)
		cd.mutex.Unlock()
	} else {
		collectError(date, errors.New(resp.Status), cd)
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

func collectError(reason string, err error, cd *collectedData) {
	e := fmt.Sprintf("with %s got error: %s", reason, err)
	cd.mutex.Lock()
	cd.errors = append(cd.errors, e)
	cd.mutex.Unlock()
}
