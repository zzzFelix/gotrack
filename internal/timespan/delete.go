package timespan

import (
	"fmt"
	"time"

	"github.com/zzzFelix/gotrack/internal/database"
)

func Delete(date time.Time) error {
	fmtDate := date.Format(time.DateOnly)
	err := database.Delete(fmtDate)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted times for %s\n", fmtDate)
	return nil
}
