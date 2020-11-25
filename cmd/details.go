package cmd

import (
	u "github.com/go143/cliq/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(detailsCmd)
}

var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Details of a certain something",
	Long:  `This command is a sample command to get you up to speed`,
	Run:   RUN(detailsRunner),
}

func detailsRunner() {
	REQUIRED(PIPELINE)
	u.Yellow("Project Details are printed")
}
