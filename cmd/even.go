package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "A brief description of your command",
	Long:  `Add Event Integer Numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 == 0 {
				evenSum = evenSum + itemp
			}
		}
		fmt.Printf("The even addition of %s is %d\n", args, evenSum)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)
}
