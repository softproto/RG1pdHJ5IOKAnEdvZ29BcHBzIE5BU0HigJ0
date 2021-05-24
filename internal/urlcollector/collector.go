package urlcollector

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const dateLayout = "2006-01-02"

//Splitting a date range into a slice of dates
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

//Preparing and concurrently running a pool of urlsFetcher() goroutines with a limit of CONCURRENT_REQUESTS
func runCollector(config *Config, start_date, end_date string) *collectedData {
	fmt.Println("runCollector()")
	defer fmt.Println("runCollector() done")
	cd := collectedData{
		mutex:  sync.Mutex{},
		Urls:   []string{},
		Errors: []string{},
	}

	dates, err := getDatesList(start_date, end_date)
	if err != nil {
		cd.collectError(fmt.Sprintf("getDatesList(%s, %s)", start_date, end_date), err)
		return &cd
	}

	d := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < config.concurrentRequests; i++ {
		wg.Add(1)
		go urlsFetcher(d, config, &wg, &cd)
	}
	for _, date := range dates {
		d <- date
	}
	close(d)
	wg.Wait()

	return &cd
}

//Fetch URLs from Apod server for each date from dates range
func urlsFetcher(d chan string, config *Config, wg *sync.WaitGroup, cd *collectedData) {
	fmt.Println("..urlsFetcher()")
	defer fmt.Println("..urlsFetcher() done")

	defer wg.Done()

	for date := range d {
		sendRequest(config, date, cd)
	}
}
