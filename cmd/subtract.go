package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var subtractCmd = &cobra.Command{
	Use:      "subtract",
	Aliases:  []string{"sub"},
    Short:    "Subtract a number from another",
	Long:     "Carry out subtraction operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd * cobra.Command, args []string) {
		     fmt.Printf("Subtraction of %s and %s = %s \n\n", args[0], args[1], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}