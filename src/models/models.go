package models

type GenerateKeyParams struct {
	KeyName    string
	KeyType    string
	KeySize    int
	Passphrase string
	Comment    string
	OutputDir  string
	Expiry     string
	Force      bool
	Tags       []string
}

type GeneratedKey struct {
	KeyName        string
	KeysOutputDir  string
	PrivateKeyPath string
	PublicKeyPath  string
	Warnings       []string
}

type KeysManagement struct {
	KeysPaths []KeyPath   `yaml:"keys-paths"`
	Defaults  KeyDefaults `yaml:"defaults"`
}

type KeyPath struct {
	Path      string `yaml:"path"`
	Recursive bool   `yaml:"recursive"`
}

type KeyDefaults struct {
	KeyType        string   `yaml:"key-type"`
	KeySize        int      `yaml:"key-size"`
	OutDir         string   `yaml:"out-dir"`
	Comment        string   `yaml:"comment"`
	Passphrase     string   `yaml:"passphrase"`
	Expiry         string   `yaml:"expiry"`
	ForceOverwrite bool     `yaml:"force-overwrite"`
	Tags           []string `yaml:"tags"`
}

type AppManagement struct {
	LogLevel string `yaml:"log-level"`
}

type Config struct {
	AppManagement  AppManagement  `yaml:"app-management"`
	KeysManagement KeysManagement `yaml:"keys-management"`
}
