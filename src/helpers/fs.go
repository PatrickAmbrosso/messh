package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetExecPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}
	return filepath.Dir(execPath), nil
}

func ResolveRootPath(rootPath string) (string, error) {

	if rootPath == "" {
		return GetExecPath()
	}

	if strings.HasPrefix(rootPath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to resolve ~ to home directory: %w", err)
		}
		rootPath = filepath.Join(homeDir, strings.TrimPrefix(rootPath, "~"))
	}

	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path for %q: %w", rootPath, err)
	}

	return absPath, nil
}

func CheckPathExists(path string, isDir bool) (string, error) {

	if path == "" && !isDir {
		return "", fmt.Errorf("file checking failed: received empty path")
	}

	if path == "" && isDir {
		execPath, err := GetExecPath()
		if err != nil {
			return "", fmt.Errorf("folder check failed: executable path cannot be determined: %w", err)
		}
		path = execPath
	}

	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		path = filepath.Join(homeDir, path[1:])
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("path not found: %s", absPath)
	} else if err != nil {
		return "", fmt.Errorf("error checking path %s: %w", absPath, err)
	}

	if isDir && !info.IsDir() {
		return "", fmt.Errorf("path exists but is NOT a directory: %s", absPath)
	}
	if !isDir && info.IsDir() {
		return "", fmt.Errorf("path exists but is a directory: %s", absPath)
	}

	return absPath, nil
}
