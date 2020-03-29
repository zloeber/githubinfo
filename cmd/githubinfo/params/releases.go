package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// releasesCmd represents the get command
var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Github releases information.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a vendor/repo argument")
		}
		if api.IsValidProject(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid Github project specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		releasesJSON := string(api.ReleasesJSON(args[0]))
		releases := api.ReleaseURLs(releasesJSON)
		if !api.IsJSON(releasesJSON) {
			log.Error("cannot parse for json, possibly not online?")
		} else {
			fmt.Println("Project: ", args[0])
			fmt.Println("JSON: ", releases)
		}
	},
}

func init() {
	rootCmd.AddCommand(releasesCmd)
}
