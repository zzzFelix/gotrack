package timespan

import (
	"time"

	"github.com/zzzFelix/gotrack/database"
)

func Save(ts Timespan, date time.Time) {
	ts.Print(date)
	database.Persist(date.Format(time.DateOnly), ts.marshal())
}

func SaveToday(ts Timespan) {
	Save(ts, time.Now())
}
