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
