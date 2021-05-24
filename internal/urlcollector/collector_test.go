package urlcollector

import (
	"testing"
)

var pos = []struct {
	start_date string
	end_date   string
	list       []string
}{
	{"2001-01-01", "2001-01-03", []string{"2001-01-01", "2001-01-02", "2001-01-03"}},
	{"2001-01-01", "2001-01-01", []string{"2001-01-01"}},
}

var neg = []struct {
	start_date string
	end_date   string
	list       []string
}{
	{"2001-01-03", "2001-01-01", []string{}}, //"start_date should be before or equal end_date"
	{"2001-01-01", "", []string{}},           //"start_date and end_date must be provided"
	{"2001-01-00", "2001-01-01", []string{}}, //"parsing time \"2001-04-00\": day out of range"
}

func TestGetDatesListPositive(t *testing.T) {
	for _, e := range pos {
		list, err := getDatesList(e.start_date, e.end_date)
		if err != nil {
			t.Errorf("Got result %v with %s, %s , expected %v", list, e.start_date, e.end_date, e.list)
		}
	}
}

func TestGetDatesListNegative(t *testing.T) {
	for _, e := range neg {
		list, err := getDatesList(e.start_date, e.end_date)
		if err == nil {
			t.Errorf("Got result %v with %s, %s , expected %v", list, e.start_date, e.end_date, e.list)
		}
	}
}
