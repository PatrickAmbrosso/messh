package config

import (
	_ "embed"
	"fmt"
	"messh/src/helpers"
	"messh/src/models"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	once   sync.Once
	cfg    *models.Config
	cfgSrc string
	err    error
)

// Embeds /////////////////////////////////////////////////////////////////////

//go:embed messh-config.yaml
var defaultConfig []byte

// Configuration Accessors ////////////////////////////////////////////////////

// loads the active config (singleton) from default path or falls back to embedded config
func GetConfig(getTemplate bool) (*models.Config, string, error) {
	if getTemplate {
		return loadConfig("", true)
	} else {
		once.Do(func() {
			cfg, cfgSrc, err = loadConfig("", false)
		})
		return cfg, cfgSrc, err
	}
}

func GetDefaultConfig() (*models.Config, error) {
	decoder := yaml.NewDecoder(strings.NewReader(string(defaultConfig)))
	decoder.KnownFields(true)

	var cfg models.Config
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("error decoding YAML: %w", err)
	}

	return &cfg, nil
}

// loads a config from another file and saves it to the default path
func ImportConfig(fromPath string, overwrite bool) (string, error) {
	absFromPath, err := resolveConfigPath(fromPath)
	if err != nil {
		return "", fmt.Errorf("invalid source config path: %w", err)
	}

	data, err := os.ReadFile(absFromPath)
	if err != nil {
		return "", fmt.Errorf("failed to read source config file: %w", err)
	}

	decoder := yaml.NewDecoder(strings.NewReader(string(data)))
	decoder.KnownFields(true)

	var cfg models.Config
	if err := decoder.Decode(&cfg); err != nil {
		return "", fmt.Errorf("failed to decode source config: %w", err)
	}

	return saveConfig(&cfg, "", overwrite) // write to default path
}

// writes the embedded default config to a path
func WriteConfig(path string, overwrite bool) (string, error) {
	if path == "" {
		path = "config-template.yaml"
	}
	if _, err := os.Stat(path); err == nil && !overwrite {
		return "", fmt.Errorf("file already exists at %s (overwrite disabled)", path)
	}
	if err := os.WriteFile(path, defaultConfig, 0644); err != nil {
		return "", fmt.Errorf("failed to write embedded config: %w", err)
	}
	return path, nil
}

// exports the currently loaded config to a given path
func ExportConfig(path string, overwrite bool) (string, error) {

	if path == "" {
		path = "config-export.yaml"
	}

	cfg, _, err := GetConfig(false)
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}
	return saveConfig(cfg, path, overwrite)
}

// Configuration Helpers ////////////////////////////////////////////

func resolveConfigPath(configPath string) (string, error) {
	if configPath == "" {
		rootPath, err := helpers.ResolveRootPath(configPath)
		if err != nil {
			return "", fmt.Errorf("received no config path, error resolving root path: %w", err)
		}
		configPath = filepath.Join(rootPath, "config.yaml")
	}

	absConfigPath, err := filepath.Abs(configPath)
	if err != nil {
		return "", fmt.Errorf("error resolving absolute path for config file: %w", err)
	}

	if ext := filepath.Ext(absConfigPath); !(ext == ".yaml" || ext == ".yml") {
		return "", fmt.Errorf("unknown file %s received. Only .yaml & .yml files supported", filepath.Base(configPath))
	}

	return absConfigPath, nil
}

func loadConfig(configPath string, getTemplate bool) (*models.Config, string, error) {

	marshallConfig := func(data []byte) (*models.Config, error) {
		decoder := yaml.NewDecoder(strings.NewReader(string(data)))
		decoder.KnownFields(true)
		var cfg models.Config
		if err := decoder.Decode(&cfg); err != nil {
			return nil, fmt.Errorf("error decoding YAML: %w", err)
		}
		return &cfg, nil
	}

	if getTemplate {
		cfg, err := marshallConfig(defaultConfig)
		if err != nil {
			return nil, "", err
		}
		return cfg, "config template", nil
	}

	isEmbedded := false

	absConfigPath, err := resolveConfigPath(configPath)
	if err != nil {
		return nil, "", err
	}

	var data []byte

	if _, err := os.Stat(absConfigPath); os.IsNotExist(err) {
		data = defaultConfig
		isEmbedded = true
	} else if err == nil {
		data, err = os.ReadFile(absConfigPath)
		if err != nil {
			return nil, "", fmt.Errorf("error reading config file: %w", err)
		}
	} else {
		return nil, "", fmt.Errorf("error checking config path: %w", err)
	}

	if isEmbedded {
		cfgSrc = "embedded config"
	} else {
		cfgSrc = absConfigPath
	}

	cfg, err := marshallConfig(data)
	if err != nil {
		return nil, "", err
	}

	return cfg, cfgSrc, nil
}

func saveConfig(cfg *models.Config, path string, overwrite bool) (string, error) {
	if cfg == nil {
		return "", fmt.Errorf("cannot save: config is nil")
	}

	absPath, err := resolveConfigPath(path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve config path: %w", err)
	}

	if _, err := os.Stat(absPath); err == nil && !overwrite {
		return "", fmt.Errorf("file already exists at %s (overwrite disabled)", absPath)
	}

	out, err := yaml.Marshal(cfg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return "", fmt.Errorf("failed to create directories: %w", err)
	}

	if err := os.WriteFile(absPath, out, 0644); err != nil {
		return "", fmt.Errorf("error writing config to file: %w", err)
	}

	return absPath, nil
}
