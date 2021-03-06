package params

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zloeber/githubinfo/pkg/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version.Version)
		if verbose {
			fmt.Println("Application:", version.AppName)
			fmt.Println("Build Date:", version.BuildDate)
			fmt.Println("Git Commit:", version.GitCommit)
			fmt.Println("Go Version:", version.GoVersion)
			fmt.Println("OS / Arch:", version.OsArch)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
