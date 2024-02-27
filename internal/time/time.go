package time

import (
	"fmt"
	"strings"
	"time"
)

const (
	TIME_ONLY = "15:04"
	HOUR_ONLY = "15"
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
	layout := TIME_ONLY

	hourOnly := !strings.Contains(input, ":")
	if hourOnly {
		layout = HOUR_ONLY
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

func GetTotalDuration(start time.Time, end time.Time, brk time.Duration) time.Duration {
	return end.Sub(start) - brk
}

func FormatTimeOnly(input time.Time) string {
	return input.Format(TIME_ONLY)
}

func FormatDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func FormatDay(t time.Time) string {
	weekday := t.Weekday().String()[:3]
	dateOnly := t.Format(time.DateOnly)
	return fmt.Sprintf("%s %s", weekday, dateOnly)
}
