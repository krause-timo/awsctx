package tui

import tea "charm.land/bubbletea/v2"

// Update implements tea.Model. It handles key presses: up/down (k/j) move the cursor,
// enter/space confirm the profile under the cursor and quit, and q/ctrl+c quit without choosing.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.profiles)-1 {
				m.cursor++
			}
		case "enter", "space":
			m.choice = &m.profiles[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}
