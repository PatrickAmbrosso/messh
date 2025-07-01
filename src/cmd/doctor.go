package cmd

import (
	"messh/src/constants"
	"messh/src/out"
	"os"

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
		logger, err := out.NewLogger(out.LoggerOptions{
			Level:     "info",
			LogToFile: flagDoctorQuiet,
			FilePath:  flagDoctorFile,
		})
		if err != nil {
			out.SCLogger.Error("failed to initialize logger: " + err.Error())
			os.Exit(1)
		}

		logger.Info("running " + constants.AppAbbrName + " app diagnostics")

		// This command performs a series of environment and application checks.
		// Below are the planned diagnostics to implement:
		//
		// 1. Config File Checks
		// ----------------------
		// - Check if the config file exists
		// - Validate if the config is proper YAML
		// - Warn if critical fields like `keys-paths` or `defaults.key-type` are missing
		// - Log the path from which the config was loaded

		// 2. SSH Key Checks
		// ------------------
		// - Count number of SSH private/public key pairs in configured directories
		// - Validate private key format (e.g., starts with "-----BEGIN OPENSSH PRIVATE KEY-----")
		// - Check key file permissions (should be 0600 for private keys)
		// - If expiry is configured, show days left for each key (if metadata is tracked)

		// 3. Dependency Checks
		// ---------------------
		// - Check if `ssh`, `ssh-keygen`, and optionally `ssh-agent` exist in PATH
		// - Log warning if not found

		// 4. File System and Directory Checks
		// ------------------------------------
		// - Ensure default output directory for keys exists and is writable
		// - Check that paths in `keys-paths` exist (or warn if not)
		// - Warn if paths are not readable or are symlinks

		// 5. Output and Behavior
		// ------------------------
		// - Respect `--quiet` to suppress stdout logs
		// - Respect `--file` to write results to a specified log file
		// - Default to console logging if neither is set

		// 6. Recommendations
		// --------------------
		// - If config is missing, suggest running `messh config init`
		// - If key directory is missing, suggest creating it
		// - If key permissions are wrong, recommend `chmod 600 <keyfile>`// This command performs a series of environment and application checks.
		// Below are the planned diagnostics to implement:
		//
		// 1. Config File Checks
		// ----------------------
		// - Check if the config file exists
		// - Validate if the config is proper YAML
		// - Warn if critical fields like `keys-paths` or `defaults.key-type` are missing
		// - Log the path from which the config was loaded

		// 2. SSH Key Checks
		// ------------------
		// - Count number of SSH private/public key pairs in configured directories
		// - Validate private key format (e.g., starts with "-----BEGIN OPENSSH PRIVATE KEY-----")
		// - Check key file permissions (should be 0600 for private keys)
		// - If expiry is configured, show days left for each key (if metadata is tracked)

		// 3. Dependency Checks
		// ---------------------
		// - Check if `ssh`, `ssh-keygen`, and optionally `ssh-agent` exist in PATH
		// - Log warning if not found

		// 4. File System and Directory Checks
		// ------------------------------------
		// - Ensure default output directory for keys exists and is writable
		// - Check that paths in `keys-paths` exist (or warn if not)
		// - Warn if paths are not readable or are symlinks

		// 5. Output and Behavior
		// ------------------------
		// - Respect `--quiet` to suppress stdout logs
		// - Respect `--file` to write results to a specified log file
		// - Default to console logging if neither is set

		// 6. Recommendations
		// --------------------
		// - If config is missing, suggest running `messh config init`
		// - If key directory is missing, suggest creating it
		// - If key permissions are wrong, recommend `chmod 600 <keyfile>`
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)

	doctorCmd.Flags().BoolVarP(&flagDoctorQuiet, "quiet", "q", false, "quiet mode with no console output")
	doctorCmd.Flags().StringVarP(&flagDoctorFile, "file", "f", "", "file to write the diagnostic report")
}
