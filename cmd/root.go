package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hlts2/goliday"
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
			week, name := weekAndName(holiday)
			table.Append([]string{
				strconv.Itoa(holiday.Year()),
				strconv.Itoa(holiday.Month()),
				strconv.Itoa(holiday.Day()),
				week,
				name,
			})
		}

		table.Render()
	},
}

var (
	year  int
	month int
	day   int
	mode  string
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", -1, "set year")
	rootCmd.PersistentFlags().IntVarP(&month, "month", "m", -1, "set month")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", -1, "set day")
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "ja", "set mode(ja, en)")
}

func weekAndName(h goliday.Holiday) (string, string) {
	switch mode {
	case "ja":
		return h.Week(), h.Name()
	default:
		return h.WeekEn(), h.NameEn()
	}
}

// Execute executes holijp application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
