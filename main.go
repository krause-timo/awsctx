// Command awsctx is an interactive TUI for switching the active AWS SSO profile.
package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/krause-timo/awsctx/internal/config"
	"github.com/krause-timo/awsctx/internal/sso"
	"github.com/krause-timo/awsctx/internal/tui"
)

// main runs the program and reports any error to stderr with a non-zero exit code.
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "awsctx:", err)
		os.Exit(1)
	}
}

// run loads the AWS profiles, lets the user pick one via the TUI, and logs in to it.
func run() error {
	logFile, err := os.Create("awsctx.log")

	if err != nil {
		return fmt.Errorf("creating log file: %w", err)
	}

	defer logFile.Close()

	handler := slog.NewTextHandler(logFile, nil)
	logger := slog.New(handler)

	profiles, err := config.Profiles()
	if err != nil {
		logger.Error("loading aws profiles", "error", err)
		return err
	}

	chosen, err := tui.Select(profiles)
	if err != nil {
		return fmt.Errorf("launching tui: %w", err)
	}

	if chosen == nil {
		return nil // user cancelled (q/ctrl+c) → exit quietly, no login, no print
	}

	sso.Login(chosen.Name)

	return nil
}
