package cmd

import (
	"fmt"
	"messh/src/config"
	"messh/src/constants"
	"messh/src/models"
	"messh/src/out"
	"os"
	"strings"

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

		var cfg *models.Config
		var cfgSrc string
		var err error

		if flagConfigShowTemplate {
			cfg, cfgSrc, err = config.GetConfig(true)
		} else {
			cfg, cfgSrc, err = config.GetConfig(false)
		}

		if err != nil {
			out.SCLogger.Error(err.Error())
			os.Exit(1)
		}

		if !flagConfigShowQuiet {
			msg := "Config for " + constants.AppAbbrName + " app - version " + constants.AppVersion + " (source: " + cfgSrc + ")"
			fmt.Println(out.Banner(msg))
			fmt.Println()

			fmt.Print(consolePrintConfig(cfg))

			// cfgStr, err := yaml.Marshal(&cfg)
			// if err != nil {
			// 	out.SCLogger.Error(err.Error())
			// 	os.Exit(1)
			// }

			// if !flagConfigShowQuiet {
			// 	fmt.Print(string(cfgStr))
			// }

		}

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

func consolePrintConfig(cfg *models.Config) string {
	var b strings.Builder

	b.WriteString(out.SectionHeader("Application Configuration") + "\n")
	b.WriteString(out.KV("Log Level", cfg.AppManagement.LogLevel) + "\n")

	b.WriteString(out.SectionHeader("Keys Management") + "\n")
	b.WriteString(out.KV("Key Type", cfg.KeysManagement.Defaults.KeyType) + "\n")
	b.WriteString(out.KV("Key Size", fmt.Sprintf("%d", cfg.KeysManagement.Defaults.KeySize)) + "\n")
	b.WriteString(out.KV("Output Directory", cfg.KeysManagement.Defaults.OutDir) + "\n")
	b.WriteString(out.KV("Comment", cfg.KeysManagement.Defaults.Comment) + "\n")
	// b.WriteString(out.KV("Passphrase", maskIfSet(cfg.KeysManagement.Defaults.Passphrase)) + "\n")
	b.WriteString(out.KV("Expiry", cfg.KeysManagement.Defaults.Expiry) + "\n")
	b.WriteString(out.KV("Force Overwrite", fmt.Sprintf("%v", cfg.KeysManagement.Defaults.ForceOverwrite)) + "\n")

	// Optional: Add tags display
	if len(cfg.KeysManagement.Defaults.Tags) > 0 {
		b.WriteString(out.KV("Tags", strings.Join(cfg.KeysManagement.Defaults.Tags, ", ")) + "\n")
	}

	return b.String()
}
