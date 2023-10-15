package timespan

import (
	"fmt"
	"time"

	"github.com/zzzFelix/gotrack/database"
)

func Delete(date time.Time) {
	fmtDate := date.Format(time.DateOnly)
	database.Delete(fmtDate)
	fmt.Printf("Deleted times for %s\n", fmtDate)
}
