package helpers

import (
	"os/exec"
	"sync"
)

var (
	checkOnce       sync.Once
	binaryAvailable = make(map[string]bool)
	binaries        = []string{
		"ssh",
		"ssh-keygen",
		"scp",
		"ssh-add",
		"ssh-agent",
		"ssh-copy-id",
	}
)

// checkBinaries ensures lookup is only done once
func checkBinaries() {
	checkOnce.Do(func() {
		for _, bin := range binaries {
			_, err := exec.LookPath(bin)
			binaryAvailable[bin] = err == nil
		}
	})
}

// IsAvailable returns true if the binary is found in PATH
func IsAvailable(name string) bool {
	checkBinaries()
	return binaryAvailable[name]
}

// Convenience functions

func SSHAvailable() bool       { return IsAvailable("ssh") }
func SSHKeygenAvailable() bool { return IsAvailable("ssh-keygen") }
func SCPAvailable() bool       { return IsAvailable("scp") }
func SSHAddAvailable() bool    { return IsAvailable("ssh-add") }
func SSHAgentAvailable() bool  { return IsAvailable("ssh-agent") }
func SSHCopyIDAvailable() bool { return IsAvailable("ssh-copy-id") }
