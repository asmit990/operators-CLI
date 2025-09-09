package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"addition"},
	Short: "Add 2 numbers",
	Long: "Carry out  addition operation on 2 numbers",
    Run: func(cmd * cobra.Command, args []string) {
		fmt.Printf("Addition of %s and %s = %s \n\n", args[0], args[1], add(args[0], args[1]))
	},
}

func init() {
    rootCmd.AddCommand(addCmd)
}