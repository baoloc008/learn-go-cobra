package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var oddCmd = &cobra.Command{
	Use:   "odd",
	Long: `Add Odd Integer Numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		var oddSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 != 0 {
				oddSum = oddSum + itemp
			}
		}
		fmt.Printf("The odd addition of %s is %d\n", args, oddSum)
	},
}

func init() {
	addCmd.AddCommand(oddCmd)
}
