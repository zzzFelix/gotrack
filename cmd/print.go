package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zzzFelix/gotrack/timespan"
	"github.com/zzzFelix/gotrack/util"
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
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			timespan.GetToday()
		case 1:
			date := util.ParseDate(args[0])
			timespan.Get(date)
		case 2:
			start := util.ParseDate(args[0])
			end := util.ParseDate(args[1])
			timespan.GetMultipleDays(start, end)
		}
	},
}
