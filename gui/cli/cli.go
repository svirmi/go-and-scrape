package cli

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GuiModel struct {
	news     []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func PopulateGeneralNewsModel(news []string) GuiModel {
	var renderedNews []string
	for _, n := range news {
		renderedNews = append(renderedNews, historyStyle.Copy().Render(n))
	}

	return GuiModel{
		// Our shopping list is a grocery list
		news: renderedNews,

		// A map which indicates which news are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `news` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m GuiModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m GuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.news)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated GuiModel to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m GuiModel) View() string {
	// The header
	s := "What should we buy at the market?\n"

	// Render the row
	//doc := strings.Builder{}
	rows := ""

	// Iterate over our news
	for i, choice := range m.news {

		// Is the cursor pointing at this choice?
		cursor := "" // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		if i == 0 || i%3 != 0 {
			rows = lipgloss.JoinHorizontal(lipgloss.Top, rows, fmt.Sprintf("%d. %s [%s] %s   ", i, cursor, checked, choice))
		} else {
			s = lipgloss.JoinVertical(lipgloss.Left, s, rows)
			rows = ""
			rows = lipgloss.JoinHorizontal(lipgloss.Top, rows, fmt.Sprintf("%d. %s [%s] %s   ", i, cursor, checked, choice))
		}
	}

	// The footer
	s = lipgloss.JoinVertical(lipgloss.Left, s, rows, "\nPress q to quit.\n")

	// Send the UI for rendering
	return s
}