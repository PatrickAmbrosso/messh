package cmd

import (
	"messh/src/out"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:           "history",
	Short:         helpHistoryCmd,
	Long:          out.Banner(helpHistoryCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
