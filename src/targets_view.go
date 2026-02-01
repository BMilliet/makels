package src

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TargetsViewModel struct {
	targets       []MakeTarget
	config        *Config
	cursor        int
	viewportStart int
	maxVisible    int
	selected      *string
	quitting      bool
	styles        *Styles
	searchMode    bool
	searchQuery   string
	filteredList  []MakeTarget
}

func NewTargetsViewModel(targets []MakeTarget, config *Config) TargetsViewModel {
	m := TargetsViewModel{
		targets:       targets,
		config:        config,
		cursor:        0,
		viewportStart: 0,
		maxVisible:    4, // Show only 4 items at a time
		quitting:      false,
		styles:        DefaultStyles(),
		searchMode:    false,
		searchQuery:   "",
		filteredList:  []MakeTarget{},
	}

	return m
}

func (m TargetsViewModel) Init() tea.Cmd {
	return nil
}

func (m TargetsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.searchMode {
				m.searchMode = false
				m.searchQuery = ""
				m.filteredList = []MakeTarget{}
				m.cursor = 0
				m.viewportStart = 0
				return m, nil
			}
			*m.selected = ExitSignal
			m.quitting = true
			return m, tea.Quit

		case "esc":
			if m.searchMode {
				m.searchMode = false
				m.searchQuery = ""
				m.filteredList = []MakeTarget{}
				m.cursor = 0
				m.viewportStart = 0
				return m, nil
			}
			*m.selected = ExitSignal
			m.quitting = true
			return m, tea.Quit

		case "/":
			if !m.searchMode {
				m.searchMode = true
				m.searchQuery = ""
				m.filteredList = []MakeTarget{}
				m.cursor = 0
				m.viewportStart = 0
				return m, nil
			}

		case "backspace":
			if m.searchMode && len(m.searchQuery) > 0 {
				m.searchQuery = m.searchQuery[:len(m.searchQuery)-1]
				m.updateFilteredList()
				m.cursor = 0
				m.viewportStart = 0
				return m, nil
			}

		case "up", "k":
			items := m.getActiveList()
			if m.cursor > 0 {
				m.cursor--
				// Adjust viewport to keep cursor visible with offset
				if m.cursor < m.viewportStart {
					m.viewportStart = m.cursor
				}
			} else {
				// Wrap to bottom
				m.cursor = len(items) - 1
				if len(items) > m.maxVisible {
					m.viewportStart = len(items) - m.maxVisible
				} else {
					m.viewportStart = 0
				}
			}

		case "down", "j":
			items := m.getActiveList()
			if m.cursor < len(items)-1 {
				m.cursor++
				// Adjust viewport to keep cursor visible with offset
				// Keep cursor in the middle-ish of viewport
				if m.cursor >= m.viewportStart+m.maxVisible {
					m.viewportStart = m.cursor - m.maxVisible + 1
				}
			} else {
				// Wrap to top
				m.cursor = 0
				m.viewportStart = 0
			}

		case "enter":
			items := m.getActiveList()
			if len(items) > 0 && m.cursor < len(items) {
				// Target selection
				selectedTarget := items[m.cursor]
				result := fmt.Sprintf("target|%s", selectedTarget.Name)
				*m.selected = result
				m.quitting = true
				return m, tea.Quit
			}

		default:
			if m.searchMode {
				key := msg.String()
				if len(key) == 1 {
					m.searchQuery += key
					m.updateFilteredList()
					m.cursor = 0
					m.viewportStart = 0
					return m, nil
				}
			}
		}
	}

	return m, nil
}

func (m TargetsViewModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	// Title
	titleStyle := lipgloss.NewStyle().
		Foreground(m.styles.TitleColor).
		Bold(true).
		PaddingLeft(1)

	b.WriteString("\n")
	b.WriteString(titleStyle.Render("Makefile Targets"))
	b.WriteString("\n\n")

	// Show search box if in search mode
	if m.searchMode {
		searchBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(m.styles.SearchBoxColor).
			Padding(0, 1).
			Width(70).
			Foreground(m.styles.SearchTextColor)

		searchText := fmt.Sprintf("🔍 Search: %s", m.searchQuery)
		if m.searchQuery == "" {
			searchText = "🔍 Search: (type to search...)"
		}
		b.WriteString(searchBox.Render(searchText))
		b.WriteString("\n\n")
	}

	// Current page items
	items := m.getActiveList()
	if len(items) == 0 {
		if m.searchMode {
			b.WriteString(m.styles.FooterStyle.Render("  No matches found\n"))
		} else {
			b.WriteString(m.styles.FooterStyle.Render("  No targets found in Makefile\n"))
		}
	} else {
		visibleEnd := m.viewportStart + m.maxVisible
		if visibleEnd > len(items) {
			visibleEnd = len(items)
		}

		if m.viewportStart > 0 {
			b.WriteString(m.styles.FooterStyle.Render("  ⬆ More items above..."))
			b.WriteString("\n\n")
		}

		for i := m.viewportStart; i < visibleEnd; i++ {
			var itemBox lipgloss.Style
			var borderColor lipgloss.Color

			if m.cursor == i {
				borderColor = m.styles.SelectedTitleColor
				itemBox = lipgloss.NewStyle().
					Border(lipgloss.RoundedBorder()).
					BorderForeground(borderColor).
					Padding(0, 1).
					Width(70).
					MarginLeft(2)
			} else {
				borderColor = m.styles.MutedBorderColor
				itemBox = lipgloss.NewStyle().
					Border(lipgloss.RoundedBorder()).
					BorderForeground(borderColor).
					Padding(0, 1).
					Width(70)
			}

			titleStyle := lipgloss.NewStyle().Bold(true)
			var valueColor lipgloss.Color
			if m.cursor == i {
				titleStyle = titleStyle.Foreground(m.styles.SelectedTitleColor)
				valueColor = m.styles.FooterColor
			} else {
				titleStyle = titleStyle.Foreground(m.styles.MutedTitleColor)
				valueColor = m.styles.MutedTitleColor
			}

			valueStyle := lipgloss.NewStyle().
				Foreground(valueColor).
				Width(66).
				Italic(true)

			// Render target item
			item := items[i]
			var titleText string
			if m.searchMode && m.searchQuery != "" {
				titleText = m.highlightMatches(item.Name, m.searchQuery)
			} else {
				titleText = item.Name
			}

			targetStyle := lipgloss.NewStyle().
				Foreground(m.styles.AquamarineColor).
				Bold(true)

			var descriptionText string
			if item.Description != "" {
				// Has description from comment
				descriptionText = valueStyle.Render(item.Description)
			} else if len(item.Recipe) > 0 {
				// No description, show code preview
				codeStyle := lipgloss.NewStyle().
					Foreground(m.styles.FadedCodeColor).
					Italic(true)
				
				var previewLines []string
				for _, recipeLine := range item.Recipe {
					// Truncate long lines (max 60 chars)
					if len(recipeLine) > 60 {
						recipeLine = recipeLine[:57] + "..."
					}
					previewLines = append(previewLines, recipeLine)
				}
				descriptionText = codeStyle.Render(strings.Join(previewLines, "\n"))
			} else {
				// No description and no recipe
				descriptionText = valueStyle.Render("No description")
			}

			content := fmt.Sprintf("%s\n%s\n%s",
				titleStyle.Render(titleText),
				targetStyle.Render("make"),
				descriptionText,
			)

			b.WriteString(itemBox.Render(content))
			b.WriteString("\n")
		}

		if visibleEnd < len(items) {
			b.WriteString("\n")
			b.WriteString(m.styles.FooterStyle.Render("  ⬇ More items below..."))
		}
	}

	// Footer
	b.WriteString("\n")
	var helpText string
	if m.searchMode {
		helpText = "  type to search • ↑↓/jk navigate • enter select • esc cancel"
	} else {
		helpText = "  / search • ↑↓/jk navigate • enter select • q/esc quit"
	}
	b.WriteString(m.styles.FooterStyle.Render(helpText + "\n"))

	return b.String()
}

func (m TargetsViewModel) getActiveList() []MakeTarget {
	if m.searchMode && m.searchQuery != "" {
		return m.filteredList
	}

	return m.targets
}

func (m *TargetsViewModel) updateFilteredList() {
	if m.searchQuery == "" {
		m.filteredList = []MakeTarget{}
		return
	}

	m.filteredList = []MakeTarget{}
	query := strings.ToLower(m.searchQuery)

	for _, target := range m.targets {
		nameLower := strings.ToLower(target.Name)
		descLower := strings.ToLower(target.Description)

		if fuzzyMatch(nameLower, query) || fuzzyMatch(descLower, query) {
			m.filteredList = append(m.filteredList, target)
		}
	}
}

func fuzzyMatch(text, query string) bool {
	if query == "" {
		return true
	}

	textIdx := 0
	queryIdx := 0

	for textIdx < len(text) && queryIdx < len(query) {
		if text[textIdx] == query[queryIdx] {
			queryIdx++
		}
		textIdx++
	}

	return queryIdx == len(query)
}

func (m TargetsViewModel) highlightMatches(text, query string) string {
	if query == "" {
		return text
	}

	highlightStyle := lipgloss.NewStyle().
		Background(m.styles.HighlightBgColor).
		Foreground(m.styles.HighlightFgColor)

	textLower := strings.ToLower(text)
	queryLower := strings.ToLower(query)

	var result strings.Builder
	textIdx := 0
	queryIdx := 0

	for textIdx < len(text) {
		if queryIdx < len(queryLower) && textLower[textIdx] == queryLower[queryIdx] {
			result.WriteString(highlightStyle.Render(string(text[textIdx])))
			queryIdx++
		} else {
			result.WriteByte(text[textIdx])
		}
		textIdx++
	}

	return result.String()
}

func TargetsView(targets []MakeTarget, config *Config, selected *string) {
	m := NewTargetsViewModel(targets, config)
	m.selected = selected

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("TargetsView -> ", err)
		os.Exit(1)
	}
}
