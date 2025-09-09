package cmd
import (
	"github.com/spf13/cobra"

 "fmt"
 "os"
)

var rootCmd = &cobra.Command{
   Use: "asmit",
   Short: "asmit is a simple cli for perfoming simple tasks",
   Long: "asmit is a simple cli for perfoming simple tasks",
   Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Welcome to asmit CLI!")
   },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. an error '%s\n", err)
	    os.Exit(1)
	}

}