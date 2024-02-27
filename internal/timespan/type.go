package timespan

import (
	"encoding/json"
	"fmt"
	"time"

	internaltime "github.com/zzzFelix/gotrack/internal/time"
)

type Timespan struct {
	Start time.Time
	End   time.Time
	Brk   time.Duration
}

func (ts *Timespan) marshal() []byte {
	val, err := json.Marshal(ts)
	if err != nil {
		panic(err)
	}
	return val
}

func (ts *Timespan) unmarshal(bytes []byte) *Timespan {
	err := json.Unmarshal(bytes, &ts)
	if err != nil {
		panic(err)
	}
	return ts
}

func (ts *Timespan) Print(date time.Time) {
	day := internaltime.FormatDay(date)
	fmtVal := fmt.Sprintf("%s    %s to %s", day, internaltime.FormatTimeOnly(ts.Start), internaltime.FormatTimeOnly(ts.End))
	if ts.Brk.Minutes() > 0 {
		fmtBreak := fmt.Sprintf(" (-%s)", internaltime.FormatDuration(ts.Brk))
		fmtVal += fmtBreak
	}
	fmtVal += fmt.Sprintf("    total: %s", internaltime.FormatDuration(internaltime.GetTotalDuration(ts.Start, ts.End, ts.Brk)))
	fmt.Println(fmtVal)
}
