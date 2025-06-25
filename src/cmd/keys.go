package cmd

import (
	"messh/src/config"
	"messh/src/helpers"
	"messh/src/models"
	"messh/src/out"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:           "keys",
	Short:         helpKeysCmd,
	Long:          out.Banner(helpKeysCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var keysListCmd = &cobra.Command{
	Use:           "list",
	Short:         helpKeysListCmd,
	Long:          out.Banner(helpKeysListCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var keysExportCmd = &cobra.Command{
	Use:           "export",
	Short:         helpKeysExportCmd,
	Long:          out.Banner(helpKeysExportCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var keysRemoveCmd = &cobra.Command{
	Use:           "remove",
	Short:         helpKeysRemoveCmd,
	Long:          out.Banner(helpKeysRemoveCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var keysPruneCmd = &cobra.Command{
	Use:           "prune",
	Short:         helpKeysPruneCmd,
	Long:          out.Banner(helpKeysPruneCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var keysGenerateCmd = &cobra.Command{
	Use:           "generate",
	Short:         helpKeysGenerateCmd,
	Long:          out.Banner(helpKeysGenerateCmd),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateCLIArgsCount(0, cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {

		if !helpers.SSHKeygenAvailable() {
			out.Error("ssh-keygen is not available in PATH")
			os.Exit(1)
		}

		key, err := helpers.GenerateSSHKey(&models.GenerateKeyParams{
			KeyName:    flagKeysGenerateKeyName,
			KeyType:    flagKeysGenerateKeyType,
			KeySize:    flagKeysGenerateKeySize,
			Passphrase: flagKeysGeneratePassphrase,
			Comment:    flagKeysGenerateComment,
			OutputDir:  flagKeysGenerateOutputDir,
			Expiry:     flagKeysGenerateExpiry,
			Force:      flagKeysGenerateForce,
			Tags:       flagKeysGenerateTags,
		})

		if err != nil {
			out.Error("Error generating key: " + err.Error())
			os.Exit(1)
		}

		for _, warning := range key.Warnings {
			out.Warn(warning)
		}
		out.Info("SSH keypair by name ( " + key.KeyName + " ) generated")
		out.Info("Private key saved to: " + key.PrivateKeyPath)
		out.Info("Public key saved to: " + key.PublicKeyPath)

	},
}

func init() {

	rootCmd.AddCommand(keysCmd)
	keysCmd.AddCommand(keysGenerateCmd)
	keysCmd.AddCommand(keysListCmd)
	keysCmd.AddCommand(keysExportCmd)
	keysCmd.AddCommand(keysRemoveCmd)
	keysCmd.AddCommand(keysPruneCmd)

	keysGenerateCmd.Flags().StringVarP(&flagKeysGenerateKeyName, "name", "n", "", "base name to use for the key")
	keysGenerateCmd.Flags().StringVarP(&flagKeysGenerateKeyType, "type", "t", "", "ssh key type to generate: "+strings.Join(config.AllowedSSHKeyTypes, ", "))
	keysGenerateCmd.Flags().IntVarP(&flagKeysGenerateKeySize, "size", "s", 0, "key size in bits (used only for rsa and ecdsa)")
	keysGenerateCmd.Flags().StringVarP(&flagKeysGeneratePassphrase, "passphrase", "x", "", "passphrase to protect the private key")
	keysGenerateCmd.Flags().StringVarP(&flagKeysGenerateComment, "comment", "c", "", "comment to embed in the public key (e.g., email or hostname)")
	keysGenerateCmd.Flags().StringVarP(&flagKeysGenerateOutputDir, "output-dir", "o", "", "directory to save the generated key")
	keysGenerateCmd.Flags().StringVarP(&flagKeysGenerateExpiry, "expiry", "e", "", "expiry date (natural language, e.g., '90d', '2025-12-31')")
	keysGenerateCmd.Flags().BoolVarP(&flagKeysGenerateForce, "force", "f", false, "overwrite existing key without confirmation")
	keysGenerateCmd.Flags().StringArrayVarP(&flagKeysGenerateTags, "tags", "g", []string{}, "tags to associate with the key")

}

// func saveRSAPrivateKeyToFile(key *rsa.PrivateKey, path string) error {
// 	if _, err := os.Stat(path); err == nil && !force {
// 		return fmt.Errorf("file %s already exists, use --force to overwrite", path)
// 	}

// 	keyFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		return err
// 	}
// 	defer keyFile.Close()

// 	block := &pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: x509.MarshalPKCS1PrivateKey(key),
// 	}

// 	return pem.Encode(keyFile, block)
// }

// func savePublicKey(signer ssh.Signer, path string) error {
// 	pubKeyBytes := ssh.MarshalAuthorizedKey(signer.PublicKey())
// 	return os.WriteFile(path, pubKeyBytes, 0644)
// }
