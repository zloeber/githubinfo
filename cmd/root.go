package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/zloeber/githubinfo/config"
)

var (
	// Used for flags.
	cfgFile       string
	verbose       bool
	defaultConfig *viper.Viper
	rootCmd       = &cobra.Command{
		Use:   "githubinfo zloeber/githubinfo",
		Short: "Gather and return information about a github project via the github api without a token.",
		Long: `Gather information about another github project without scraping the github api via
		 curl and bash scripts.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//      Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func initConfig() {
// 	if err := config.LoadConfigProvider("GITHUBINFO"); err != nil {
// 		log.Fatalf("Error reading config file, %s", err)
// 	}
// }

// init gets this thing started
func init() {
	cobra.OnInitialize(config.LoadConfigProvider("GITHUBINFO"))
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", config.Config().GetBool("verbose"), "verbose output")
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is config.yml)")
	rootCmd.PersistentFlags().String("config", os.ExpandEnv("./.config.yml"), "config file (default is config.yml)")

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
