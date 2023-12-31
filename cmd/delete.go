package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/zzzFelix/gotrack/timespan"
	"github.com/zzzFelix/gotrack/util"
)

func init() {
	saveCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [date]",
	Short: "Deletes times for a given date.",
	Long: `Deletes times for a given date. When no date given, today's times are delete. Examples:
	'gotrack delete 2023-05-01' -- deletes working hours for 1st May 2023.
	'gotrack delete' -- deletes working hours for today.`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			DeleteToday()
		case 1:
			date := util.ParseDate(args[0])
			Delete(date)
		}
	},
}

func Delete(date time.Time) {
	timespan.Delete(date)
}

func DeleteToday() {
	Delete(time.Now())
}
