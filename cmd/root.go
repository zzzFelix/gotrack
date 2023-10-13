package cmd

import (
	"fmt"
	"time"

	gotime "github.com/zzzFelix/gotrack/time"
)

func Execute(args []string) {
	times := gotime.ParseAllTimes(args)
	determineCmd(times)
}

func determineCmd(times []time.Time) {
	switch len(times) {
	case 0:
		fmt.Println("Help!")
	case 1:
		fmt.Println("Start time")
		fmt.Println(times)
	case 2:
		fmt.Println("Start and end time")
		fmt.Println(times)
	default:
		fmt.Println("Start, end, break")
		fmt.Println(times[:2], gotime.GetDuration(times[2]))
	}
}
