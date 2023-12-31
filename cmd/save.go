package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/zzzFelix/gotrack/timespan"
	"github.com/zzzFelix/gotrack/util"
)

var (
	saveCmd = &cobra.Command{
		Use:   "gotrack [start time] [end time] [break duration] [date]",
		Short: "A simple time tracker.",
		Long: `Gotrack is a no-frills CLI tool to track working hours. Built with love in Go. Examples:
	'gotrack 09:00 17:00 0:45 2023-05-01' -- tracks working hours from 9 to 17h with a 45 minute break for 1st May 2023.
	'gotrack 9 17 1' -- tracks working hours from 9 to 17h with a 1 hour break. Since no date is specified, times are stored for today.
	'gotrack 9 12:30' -- tracks working hours from 9 to 12:30 with no break. Since no date is specified, times are stored for today.
	'gotrack 8:30' -- tracks working hours from 8:30 to the current time, no break.`,
		Args: cobra.MatchAll(cobra.RangeArgs(1, 4), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				cmd.Help()
			case 1:
				start := util.ParseTime(args[0])
				ts := timesToTimespan(start, time.Now(), time.Duration(0))
				timespan.SaveToday(ts)
			case 2:
				times := util.ParseAllTimes(args)
				ts := timesToTimespan(times[0], times[1], time.Duration(0))
				timespan.SaveToday(ts)
			case 3:
				times := util.ParseAllTimes(args)
				brk := util.GetDuration(times[2])
				ts := timesToTimespan(times[0], times[1], brk)
				timespan.SaveToday(ts)
			case 4:
				times := util.ParseAllTimes(args[:3])
				brk := util.GetDuration(times[2])
				ts := timesToTimespan(times[0], times[1], brk)
				date := util.ParseDate(args[3])
				timespan.Save(ts, date)
			}
		},
	}
)

func timesToTimespan(start time.Time, end time.Time, brk time.Duration) timespan.Timespan {
	return timespan.Timespan{
		Start: start,
		End:   end,
		Brk:   brk,
	}
}
