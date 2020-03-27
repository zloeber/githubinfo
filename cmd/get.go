package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zloeber/githubinfo/log"
	githubinfo "github.com/zloeber/githubinfo/src"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get github project information.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a vendor/repo argument")
		}
		if githubinfo.IsValidProject(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid Github project specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectJSON := string(githubinfo.ProjectJSON(args[0]))
		if !githubinfo.IsJSON(projectJSON) {
			log.Error("cannot parse for json, possibly not online?")
		} else {
			fmt.Println("Project: ", args[0])
			fmt.Println("Description: ", githubinfo.Description(projectJSON))
			fmt.Println("License: ", githubinfo.License(projectJSON))
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
