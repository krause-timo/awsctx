package tui

import (
	tea "charm.land/bubbletea/v2"
)

// View implements tea.Model. It renders the list component inside docStyle's
// margins, in full-screen (alt screen) mode.
func (m model) View() tea.View {
	view := tea.NewView(docStyle.Render(m.list.View()))

	view.AltScreen = true
	view.WindowTitle = "awsctx: select your profile"

	return view
}
