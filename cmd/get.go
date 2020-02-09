package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zloeber/githubinfo/lib"
)

// versionCmd represents the version command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get project information.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
		  return errors.New("requires a vendor/repo argument")
		}
		if lib.isValidProject(args[0]) {
		  return nil
		}
		return fmt.Errorf("Invalid Github project specified: %s", args[0])
	  },
	  Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Input: ", args[0])
	  },
}

func init() {
	rootCmd.AddCommand(getCmd)
}
