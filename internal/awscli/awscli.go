// Package awscli wraps invocations of the AWS CLI binary.
package awscli

import (
	"fmt"
	"os"
	"os/exec"
)

// Require searches for the aws cli in the current path, and returns an error if it was not found.
func Require() error {
	_, err := exec.LookPath("aws")
	if err != nil {
		return fmt.Errorf("aws cli not in path: %w", err)
	}

	return nil
}

// SSOLogin performs the sso login with the aws cli.
func SSOLogin(profile string) error {
	cmd := exec.Command("aws", "sso", "login", "--profile", profile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("aws sso login: %w", err)
	}

	return nil
}
