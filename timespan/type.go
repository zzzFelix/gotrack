package timespan

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/zzzFelix/gotrack/util"
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
	day := util.FormatDay(date)
	fmtVal := fmt.Sprintf("%s    %s to %s", day, util.FormatTimeOnly(ts.Start), util.FormatTimeOnly(ts.End))
	if ts.Brk.Minutes() > 0 {
		fmtBreak := fmt.Sprintf(" (-%s)", util.FormatDuration(ts.Brk))
		fmtVal += fmtBreak
	}
	fmtVal += fmt.Sprintf("    total: %s", util.FormatDuration(util.GetTotalDuration(ts.Start, ts.End, ts.Brk)))
	fmt.Println(fmtVal)
}
