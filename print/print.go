package print

import (
	"time"

	"github.com/zzzFelix/gotrack/database"
)

func TimesToday() {
	Times(time.Now())
}

func Times(date time.Time) {
	dateOnly := date.Format(time.DateOnly)
	database.Get(dateOnly)
}
