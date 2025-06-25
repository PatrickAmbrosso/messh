package cmd

import (
	"fmt"
	"messh/src/config"
	"messh/src/out"
	"os"

	"github.com/spf13/cobra"
)

const (
	helpRootCmd          = "A CLI tool to manage SSH connections from your terminal."
	helpAddCmd           = "Add a new SSH connection."
	helpRemoveCmd        = "Remove an SSH connection."
	helpEditCmd          = "Edit an SSH connection."
	helpKeysCmd          = "Generate and manage SSH key pairs."
	helpKeysGenerateCmd  = "Generate a new SSH key pair."
	helpKeysListCmd      = "List all generated SSH keys."
	helpKeysExportCmd    = "Export SSH keys to a file."
	helpKeysRemoveCmd    = "Remove a specific SSH key."
	helpKeysPruneCmd     = "Delete all unused or expired SSH keys."
	helpFmtCmd           = "Format the SSH connections file."
	helpSessionsCmd      = "List the SSH sessions."
	helpSessionsPruneCmd = "Prune the SSH sessions."
)

var (
	flagKeysGenerateKeyName      string
	flagKeysGenerateKeyType      string
	flagKeysGenerateKeySize      int
	flagKeysGeneratePassphrase   string
	flagKeysGenerateComment      string
	flagKeysGenerateOutputDir    string
	flagKeysGenerateExpiry       string
	flagKeysGenerateForce        bool
	flagKeysGenerateTags         []string
	flagKeysListMatchPattern     string
	flagKeysListOutputFile       string
	flagKeysExportByIDs          []string
	flagKeysExportByNames        []string
	flagKeysExportByMatchPattern string
	flagKeysExportOutputDir      string
	flagKeysExportFormatZip      bool
	flagKeysExportFormatGzip     bool
	flagKeysExportFormatRaw      bool
	flagKeysExportPublicKeysOnly bool
	flagKeysRemoveByIDs          []string
	flagKeysRemoveByNames        []string
	flagKeysRemoveByMatchPattern string
	flagKeysRemoveConfirm        bool
	flagKeysPruneConfirm         bool
)

var rootCmd = &cobra.Command{
	Use:           config.AppAbbrName,
	Short:         helpRootCmd,
	Long:          out.Banner(helpRootCmd),
	Version:       config.AppVersion,
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
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Help()
		fmt.Println()
		out.Error("Flag parse error", "error", err.Error())
		os.Exit(1)
		return nil
	})
}

func validateCLIArgsCount(n int, cmd *cobra.Command, args []string) {
	if len(args) != n {
		_ = cmd.Help()
		fmt.Println()
		out.Error(fmt.Sprintf("Command '%s' expects exactly %d argument(s).", cmd.CommandPath(), n))
		out.Info(fmt.Sprintf("Run '%s --help' for usage.", cmd.CommandPath()))
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
