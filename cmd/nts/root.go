package nts

import "fmt"
import "github.com/spf13/cobra"
import "os"

var rootCmd &cobra.Command{
	Use: "nts",
	Short: "Novus Time Server configuration tool",
	Long: "Novus Time Server configuration tool is used for updating the parameters in Novus Time Server Products"

	Run: func(cmd *cobra.Command, args []string){

	},

}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}