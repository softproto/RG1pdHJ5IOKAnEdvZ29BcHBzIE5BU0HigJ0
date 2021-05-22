package urlcollector

import (
	"errors"
	"fmt"
	"gogoapps-nasa/internal/apod"
	"sync"
	"time"
)

const timeLayout = "2006-01-02"
const maxConnCount = 3

func getDatesList(startDate, endDate string) (list []string, err error) {
	fmt.Println("getDatesList()")

	sd, err := time.Parse(timeLayout, startDate)
	if err != nil {
		return nil, err
	}
	ed, err := time.Parse(timeLayout, endDate)
	if err != nil {
		return nil, err
	}

	if sd.Before(ed) || sd == ed {
		duration := int(ed.Sub(sd).Hours() / 24)
		// fmt.Println(duration + 1)
		for i := 0; i <= duration; i++ {
			date := sd.AddDate(0, 0, i)
			list = append(list, date.Format(timeLayout))
			//fmt.Println(date.Format(timeLayout))
		}
		return list, nil
	}
	return nil, errors.New("start_date should be before or equal end_date")
}

func RunCollector(apiKey string, out chan string) {
	fmt.Println("RunCollector()")
	dates, err := getDatesList("2000-01-01", "2000-01-5")
	if err != nil {
		return
	}
	fmt.Println(dates)

	d := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < maxConnCount; i++ {
		wg.Add(1)
		go urlsFetcher(d, out, apiKey, &wg)
	}

	for _, date := range dates {
		d <- date
	}
	close(d)
	wg.Wait()

}

func urlsFetcher(d chan string, out chan string, apiKey string, wg *sync.WaitGroup) {
	fmt.Println("urlsFetcher()")
	defer wg.Done()

	for date := range d {
		apod.SendRequest(apiKey, date, out)
	}

}


