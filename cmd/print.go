package cmd

import (
	"github.com/spf13/cobra"
	internaltime "github.com/zzzFelix/gotrack/internal/time"
	"github.com/zzzFelix/gotrack/internal/timespan"
)

func init() {
	saveCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print [(start) date] [end date]",
	Short: "Prints times for given date(s).",
	Long: `Prints times for a given time period. When no date given, today's times are printed. Examples:
	'gotrack print 2023-05-01 2023-05-30' -- prints working hours from 1st May 2023 to 30th May 2023 (both inclusive).
	'gotrack print 2023-05-01' -- prints working hours for 1st May 2023.
	'gotrack print' -- prints working hours for today.`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			timespan.GetToday()
		case 1:
			date := internaltime.ParseDate(args[0])
			timespan.Get(date)
		case 2:
			start := internaltime.ParseDate(args[0])
			end := internaltime.ParseDate(args[1])
			timespan.GetMultipleDays(start, end)
		}
		return nil
	},
}
