package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// View implements tea.Model, rendering the profile list with a ">" marker on the current row.
func (m model) View() tea.View {
	var builder strings.Builder
	builder.WriteString("Select the AWS profile\n\n")

	for i, profile := range m.profiles {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		fmt.Fprintf(&builder, "%s %s\n", cursor, profile.Name)
	}

	builder.WriteString("\nPress q to quit.\n")

	view := tea.NewView(builder.String())
	view.WindowTitle = "awsctx: select your profile"

	return view
}
