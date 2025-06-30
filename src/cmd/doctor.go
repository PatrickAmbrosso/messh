package cmd

import (
	"messh/src/out"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:           "doctor",
	Short:         helpDoctorCmd,
	Long:          out.Banner(helpDoctorCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		out.SCLogger.Info("this is an info message")
		out.SCLogger.Warn("this is a warning message")
		out.SCLogger.Error("this is an error message")
		out.SCLogger.Debug("this is a debug message")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)

	doctorCmd.Flags().BoolVarP(&flagDoctorQuiet, "quiet", "q", false, "quiet mode with no console output")
	doctorCmd.Flags().StringVarP(&flagDoctorFile, "file", "f", "", "file to write the diagnostic report (accepts .log, .json extensions)")
}
