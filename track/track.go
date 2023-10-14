package track

import (
	"fmt"
	"time"

	"github.com/zzzFelix/gotrack/database"
	"github.com/zzzFelix/gotrack/util"
)

func Time(start time.Time, end time.Time, brk time.Duration, date time.Time) { // brk == break; break is a reserved keyword
	fmtDate := date.Format(time.DateOnly)
	fmtVal := fmt.Sprintf("%s    %s to %s", fmtDate, util.FormatTimeOnly(start), util.FormatTimeOnly(end))
	if brk.Minutes() > 0 {
		fmtBreak := fmt.Sprintf(" (-%s)", util.FormatDuration(brk))
		fmtVal += fmtBreak
	}
	fmtVal += fmt.Sprintf("    total: %s", util.FormatDuration(util.GetTotalDuration(start, end, brk)))
	database.Persist(fmtDate, fmtVal)
	database.Get(fmtDate)
}

func TimeToday(start time.Time, end time.Time, brk time.Duration) {
	Time(start, end, brk, time.Now())
}

func TimeUntilNow(start time.Time) {
	TimeToday(start, time.Now(), time.Duration(0))
}
