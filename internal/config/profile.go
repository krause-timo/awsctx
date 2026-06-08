// Package config reads AWS profiles from the local AWS config file.
package config

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

// Profile represents a single named AWS profile from the config file.
type Profile struct {
	Name   string
	Region string
}

// Profiles reads the AWS config file and returns all named profiles found in it.
func Profiles() ([]Profile, error) {
	configPath, err := filePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)

	}
	defer file.Close()

	profiles, err := parseProfiles(file)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

// parseProfiles reads INI-formatted AWS config from r and returns its named profiles.
func parseProfiles(r io.Reader) ([]Profile, error) {
	var profiles []Profile

	cfg, err := ini.Load(r)
	if err != nil {
		return nil, fmt.Errorf("loading config file: %w", err)
	}

	for _, section := range cfg.Sections() {

		// Only named [profile X] sections; the bare [default] and [sso-session X] are intentionally skipped.
		if profileName, ok := strings.CutPrefix(section.Name(), "profile "); ok {
			profiles = append(profiles, Profile{
				Name:   profileName,
				Region: section.Key("region").String(),
			})
		}
	}

	return profiles, nil
}
