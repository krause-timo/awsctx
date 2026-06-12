// Package tui provides the interactive terminal picker for selecting an AWS profile.
package tui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/krause-timo/awsctx/internal/config"
)

// model is the Bubble Tea state for the profile picker: the list component
// showing the profiles, and the chosen profile (nil until the user confirms one).
type model struct {
	choice *config.Profile
	list   list.Model
}

type item struct {
	profile config.Profile
}

func (i item) Title() string       { return i.profile.Name }
func (i item) FilterValue() string { return i.profile.Name }
func (i item) Description() string { return "" }

// initialModel returns a model whose list is populated with the given profiles.
// The list is created with zero size; the real dimensions arrive via tea.WindowSizeMsg.
func initialModel(profiles []config.Profile) model {
	var items []list.Item

	for _, profile := range profiles {
		items = append(items, item{
			profile: profile,
		})
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.SetSpacing(0)

	ls := list.New(items, delegate, 0, 0)
	ls.Title = "Select AWS profile"

	return model{list: ls}
}

// Init implements tea.Model. The picker needs no startup command, so it returns nil.
func (m model) Init() tea.Cmd {
	return nil
}
