package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/zzzFelix/gotrack/timespan"
	"github.com/zzzFelix/gotrack/util"
)

func init() {
	saveCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print [date]",
	Short: "Prints times for a given date.",
	Long: `Prints times for a given date. When no date given, today's times are printed. Examples:
	'gotrack print 2023-05-01' -- prints working hours for 1st May 2023.
	'gotrack print' -- prints working hours for today.`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		ts := timespan.Timespan{}
		date := time.Now()
		switch len(args) {
		case 0:
			ts = timespan.GetToday()
		case 1:
			date = util.ParseDate(args[0])
			ts = timespan.Get(date)
		}
		ts.Print(date)
	},
}
