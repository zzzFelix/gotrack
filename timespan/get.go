package timespan

import (
	"fmt"
	"time"

	"github.com/zzzFelix/gotrack/database"
	"github.com/zzzFelix/gotrack/util"
)

func GetToday() Timespan {
	return Get(time.Now())
}

func Get(date time.Time) Timespan {
	dateOnly := date.Format(time.DateOnly)
	bytes, err := database.Get(dateOnly)

	ts := Timespan{}
	if err != nil {
		day := util.FormatDay(date)
		fmt.Println(day, "  ", "No tracked time")
		return ts
	}
	ts.unmarshal(bytes)
	ts.Print(date)
	return ts
}

func GetMultipleDays(start time.Time, end time.Time) time.Duration {
	totalDuration := time.Duration(0)
	dividend := 0
	for i := start; i.Before(end.AddDate(0, 0, 1)); i = i.AddDate(0, 0, 1) {
		ts := Get(i)
		duration := util.GetTotalDuration(ts.Start, ts.End, ts.Brk)
		totalDuration += duration
		if duration != time.Duration(0) {
			dividend++
		}
	}
	fmt.Printf("[TOTAL] %s\n[AVG] %s\n", util.FormatDuration(totalDuration), util.FormatDuration(average(totalDuration, dividend)))
	return totalDuration
}

func average(totalDuration time.Duration, days int) time.Duration {
	if days == 0 {
        return 0
    }
    return totalDuration / time.Duration(days)
}