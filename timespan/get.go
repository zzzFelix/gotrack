package timespan

import (
	"time"

	"github.com/zzzFelix/gotrack/database"
)

func GetToday() Timespan {
	return Get(time.Now())
}

func Get(date time.Time) Timespan {
	dateOnly := date.Format(time.DateOnly)
	bytes := database.Get(dateOnly)
	ts := Timespan{}
	ts.unmarshal(bytes)
	return ts
}
