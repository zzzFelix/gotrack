package util

import (
	"fmt"
	"strings"
	"time"
)

func ParseAllTimes(input []string) []time.Time {
	output := make([]time.Time, len(input))
	for id, element := range input {
		output[id] = ParseTime(element)
	}
	return output
}

func ParseDate(input string) time.Time {
	date, error := time.Parse(time.DateOnly, input)

	if error != nil {
		fmt.Println(error)
		return time.Now()
	}

	return date
}

func ParseTime(input string) time.Time {
	layout := "15:04" // HH:mm

	hourOnly := !strings.Contains(input, ":")
	if hourOnly {
		layout = "15" // HH
	}

	date, error := time.Parse(layout, input)

	if error != nil {
		fmt.Println(error)
		return time.Now()
	}

	return date
}

func GetDuration(input time.Time) time.Duration {
	midnight := time.Date(input.Year(), input.Month(), input.Day(), 0, 0, 0, 0, input.Location())
	return input.Sub(midnight)
}
