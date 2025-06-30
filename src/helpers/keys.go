package helpers

import (
	"fmt"
	mathRand "math/rand"
	"messh/src/constants"
	"messh/src/models"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

var adjectives = []string{
	"curious", "brave", "quiet", "lucky", "mighty", "gentle",
	"clever", "proud", "nimble", "bold", "shy", "happy", "grumpy",
}

var nouns = []string{
	"salamander", "falcon", "otter", "tiger", "penguin", "eagle",
	"panther", "koala", "badger", "whale", "rhino", "dragon", "wolf",
}

func GenerateKeyName() string {
	adj := adjectives[mathRand.Intn(len(adjectives))]
	noun := nouns[mathRand.Intn(len(nouns))]
	return fmt.Sprintf("%s-%s", adj, noun)
}

func GenerateSSHKey(params *models.GenerateKeyParams) (*models.GeneratedKey, error) {

	var warnings []string

	keyType := strings.ToLower(params.KeyType)

	if keyType != "" && !slices.Contains(constants.AllowedSSHKeyTypes, keyType) {
		return nil, fmt.Errorf("key type %s is invalid - allowed types are: %s", keyType, strings.Join(constants.AllowedSSHKeyTypes, ", "))
	}

	if keyType == "" {
		warnings = append(warnings, "no key type provided, defaulting to ed25519")
		keyType = "ed25519"
	}

	if params.KeyName == "" {
		warnings = append(warnings, "no key name provided, generating a random name")
		params.KeyName = GenerateKeyName()
	}

	outputDir := params.OutputDir
	if outputDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("unable to determine user home directory: %w", err)
		}
		outputDir = filepath.Join(home, ".ssh")
		warnings = append(warnings, fmt.Sprintf("no output directory provided, using default: %s", outputDir))
	}

	// Clean and normalize the path
	outputDir, err := filepath.Abs(outputDir)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve output location: %w", err)
	}

	// Stat the path to check if it exists and what it is
	info, err := os.Stat(outputDir)
	if err == nil {
		if !info.IsDir() {
			// It’s a file, not a directory — use its parent
			warnings = append(warnings, fmt.Sprintf("provided path is a file, using its parent directory: %s", filepath.Dir(outputDir)))
			outputDir = filepath.Dir(outputDir)
		}
	} else if os.IsNotExist(err) {
		// Path does not exist — optionally create or warn
		warnings = append(warnings, fmt.Sprintf("output directory does not exist, creating: %s", outputDir))
		if err := os.MkdirAll(outputDir, 0700); err != nil {
			return nil, fmt.Errorf("failed to create output directory: %w", err)
		}
	} else {
		return nil, fmt.Errorf("failed to stat output directory: %w", err)
	}

	outputPath := filepath.Join(outputDir, params.KeyName)

	if _, err := os.Stat(outputPath); err == nil {
		if !params.Force {
			return nil, fmt.Errorf("file already exists at %s (use --force to overwrite)", outputPath)
		}
		warnings = append(warnings, "existing key files will be overwritten due to --force")
		_ = os.Remove(outputPath)
		_ = os.Remove(outputPath + ".pub")
	}

	cmdArgs := []string{"-t", keyType, "-f", outputPath, "-q"}
	if params.Comment != "" {
		cmdArgs = append(cmdArgs, "-C", params.Comment)
	}

	switch keyType {
	case "rsa":
		validSizes := []int{2048, 4096, 8192}
		if !slices.Contains(validSizes, params.KeySize) {
			warnings = append(warnings, fmt.Sprintf("unsupported RSA key size %d, using 4096", params.KeySize))
			params.KeySize = 4096
		}
		if params.KeySize == 2048 {
			warnings = append(warnings, "RSA key size 2048 is considered weak, recommended: 4096 or 8192")
		}
		cmdArgs = append(cmdArgs, "-b", fmt.Sprintf("%d", params.KeySize))

	case "ecdsa":
		validSizes := []int{256, 384, 521}
		if !slices.Contains(validSizes, params.KeySize) {
			warnings = append(warnings, fmt.Sprintf("unsupported ECDSA key size %d, using 256", params.KeySize))
			params.KeySize = 256
		}
		cmdArgs = append(cmdArgs, "-b", fmt.Sprintf("%d", params.KeySize))
	}

	if params.Passphrase != "" {
		cmdArgs = append(cmdArgs, "-N", params.Passphrase)
	} else {
		cmdArgs = append(cmdArgs, "-N", "")
		warnings = append(warnings, "private key will be unencrypted (no passphrase set)")
	}

	cmd := exec.Command("ssh-keygen", cmdArgs...)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("ssh-keygen failed: %w", err)
	}

	return &models.GeneratedKey{
		KeyName:        params.KeyName,
		KeysOutputDir:  outputDir,
		PrivateKeyPath: outputPath,
		PublicKeyPath:  outputPath + ".pub",
		Warnings:       warnings,
	}, nil
}
