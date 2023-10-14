package track

import (
	"time"

	"github.com/zzzFelix/gotrack/database"
)

func Time(start time.Time, end time.Time, brk time.Duration, date time.Time) { // brk == break; break is a reserved keyword
	fmtDate := date.Format(time.DateOnly)
	database.Persist(fmtDate, start.String()+" "+end.String()+" "+brk.String())
	database.Get(fmtDate)
}

func TimeToday(start time.Time, end time.Time, brk time.Duration) {
	Time(start, end, brk, time.Now())
}

func TimeUntilNow(start time.Time) {
	TimeToday(start, time.Now(), time.Duration(0))
}
