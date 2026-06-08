package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// filePath returns the path where the AWS config file is located.
func filePath() (string, error) {
	configPath := os.Getenv("AWS_CONFIG_FILE")

	if configPath != "" {
		return requireFile(configPath)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolving home directory: %w", err)
	}
	configPath = filepath.Join(homeDir, ".aws/config")

	return requireFile(configPath)
}

// requireFile returns the given path if the file exists, or an error otherwise.
func requireFile(location string) (string, error) {
	_, err := os.Stat(location)
	if err != nil {
		return "", fmt.Errorf("checking aws config file: %w", err)
	}

	return location, nil
}
