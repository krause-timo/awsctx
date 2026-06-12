package tui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// docStyle adds outer margins around the whole list; Update subtracts its frame
// size from the window dimensions so the list fits inside the margins.
var docStyle = lipgloss.NewStyle().Margin(1, 2)

// Update implements tea.Model. It handles enter (confirm the selected profile and quit),
// ctrl+c (quit without choosing), and window resizes; everything else is delegated to
// the list component, which owns cursor movement, filtering, and pagination.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			selected, ok := m.list.SelectedItem().(item)
			if !ok {
				return m, nil
			}
			m.choice = &selected.profile
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
