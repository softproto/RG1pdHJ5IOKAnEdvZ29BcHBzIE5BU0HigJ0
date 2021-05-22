package urlcollector

import (
	"errors"
	"fmt"
	"gogoapps-nasa/internal/apod"
	"time"
)

const timeLayout = "2006-01-02"

func addToQueue(startDate, endDate string) error {
	fmt.Println("addToQueue()")

	sd, err := time.Parse(timeLayout, startDate)
	if err != nil {
		return err
	}
	ed, err := time.Parse(timeLayout, endDate)
	if err != nil {
		return err
	}

	if sd.Before(ed) || sd == ed {
		duration := int(ed.Sub(sd).Hours() / 24)
		fmt.Println(duration)
		for i := 0; i <= duration; i++ {
			date := sd.AddDate(0, 0, i)
			fmt.Println(date.Format(timeLayout))
		}
		return nil
	}
	return errors.New("start_date should be before or equal end_date")
}

func RunCollector(apiKey string) {
	fmt.Println("RunCollector()")

	go urlCollector(apiKey)
}

func urlCollector(apiKey string) {
	apod.SendRequest(apiKey, "3001-03-02")
}
