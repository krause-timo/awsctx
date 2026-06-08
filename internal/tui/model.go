// Package tui provides the interactive terminal picker for selecting an AWS profile.
package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/krause-timo/awsctx/internal/config"
)

// model is the Bubble Tea state for the profile picker: the available profiles,
// the cursor position, and the chosen profile (nil until the user confirms one).
type model struct {
	cursor   int
	choice   *config.Profile
	profiles []config.Profile
}

// initialModel returns a model seeded with the given profiles and the cursor at the top.
func initialModel(profiles []config.Profile) model {
	return model{
		profiles: profiles,
	}
}

// Init implements tea.Model. The picker needs no startup command, so it returns nil.
func (m model) Init() tea.Cmd {
	return nil
}
