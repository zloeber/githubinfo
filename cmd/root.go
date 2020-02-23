package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "githubinfo zloeber/githubinfo",
	Short: "Gather and return information about a github project via the github api without a token.",
	Long: `Sometimes one needs information about another github project. In these cases, one can scrape the api
	with curl and bash scripts or just use this simple utility to do the same.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//      Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init gets this thing started
func init() {
	cobra.OnInitialize()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
