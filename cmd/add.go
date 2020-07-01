package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func addInt(args []string) {
	var sum int
	// iterate over the arguments
	// the first return value is index of args, we can omit it using _

	for _, ival := range args {
		// strconv is the library used for type conversion. for string
		// to int conversion Atio method is used.
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, fval := range args {
		// convert string to float64
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + ftemp
	}
	fmt.Printf("Sum of floating numbers %s is %f\n", args, sum)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Number",
	Long: "Add Integer Or Float Numbers",
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus { // if status is true, call addFloat
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}
