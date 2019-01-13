package cmd

import (
	"fmt"
	"os"
	"strconv"

	goliday "github.com/hlts2/goliday_jp"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "holijp",
	Short: "holijp is japanese holiday viewer",
	Run: func(cmd *cobra.Command, args []string) {
		holidays := goliday.Holidays(
			goliday.WithYear(year),
			goliday.WithMonth(month),
			goliday.WithDay(day),
		)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Year", "Month", "Day", "Week", "Name"})

		for _, holiday := range holidays {
			table.Append([]string{
				strconv.Itoa(holiday.Year()),
				strconv.Itoa(holiday.Month()),
				strconv.Itoa(holiday.Day()),
				holiday.Week(),
				holiday.Name(),
			})
		}

		table.Render()
	},
}

var (
	year  int
	month int
	day   int
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", -1, "set year")
	rootCmd.PersistentFlags().IntVarP(&month, "month", "m", -1, "set year")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", -1, "set year")
}

// Execute executes holijp application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
