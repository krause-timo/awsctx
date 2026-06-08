package tui

import (
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/krause-timo/awsctx/internal/config"
)

// Select runs the interactive picker over the given profiles, rendering to stderr so
// stdout stays free for the result. It returns the chosen profile, or nil if the user cancelled.
func Select(profiles []config.Profile) (*config.Profile, error) {
	program := tea.NewProgram(initialModel(profiles), tea.WithOutput(os.Stderr))

	finalModel, err := program.Run()
	if err != nil {
		return nil, err
	}

	m := finalModel.(model)

	return m.choice, nil
}
