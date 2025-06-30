package cmd

import (
	"fmt"
	"messh/src/config"
	"messh/src/out"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:           "config",
	Short:         helpConfigCmd,
	Long:          out.Banner(helpConfigCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var configInitCmd = &cobra.Command{
	Use:           "init",
	Short:         helpConfigInitCmd,
	Long:          out.Banner(helpConfigInitCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		var savedPath string
		var err error

		if flagConfigInitFile == "" {
			// case: no --file, write embedded config
			savedPath, err = config.WriteConfig("", flagConfigInitConfirm)
			if err != nil {
				out.SCLogger.Error(err.Error())
				os.Exit(1)
			}
			out.SCLogger.Info(fmt.Sprintf("Default config written to %s", savedPath))
		} else {
			// case: --file specified, import that config
			savedPath, err = config.ImportConfig(flagConfigInitFile, flagConfigInitConfirm)
			if err != nil {
				out.SCLogger.Error(err.Error())
				os.Exit(1)
			}
			out.SCLogger.Info(fmt.Sprintf("Config imported from %s to %s", flagConfigInitFile, savedPath))
		}
	},
}

var configShowCmd = &cobra.Command{
	Use:           "show",
	Short:         helpConfigShowCmd,
	Long:          out.Banner(helpConfigShowCmd),
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
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)

	configInitCmd.Flags().StringVarP(&flagConfigInitFile, "file", "f", "", "configuration file to create or update")
	configInitCmd.Flags().BoolVarP(&flagConfigInitConfirm, "confirm", "c", false, "confirm the configuration file creation or update")

	configShowCmd.Flags().BoolVarP(&flagConfigShowTemplate, "template", "t", false, "display the configuration template")
	configShowCmd.Flags().StringVarP(&flagConfigShowExportFile, "export", "e", "", "export the configuration to a file")
	configShowCmd.Flags().BoolVarP(&flagConfigShowExportConfirm, "confirm", "c", false, "confirm the configuration export")
	configShowCmd.Flags().BoolVarP(&flagConfigShowQuiet, "quiet", "q", false, "quiet mode (no console output, only file - if set)")
}
