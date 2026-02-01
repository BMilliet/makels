package src

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	FooterColor        lipgloss.Color
	BorderColor        lipgloss.Color
	TitleColor         lipgloss.Color
	SelectedTitleColor lipgloss.Color

	FooterStyle         lipgloss.Style
	TitleStyle          lipgloss.Style
	InputField          lipgloss.Style
	InputFieldWithError lipgloss.Style

	PaginationStyle   lipgloss.Style
	HelpStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style

	PeachColor      lipgloss.Color
	CoralColor      lipgloss.Color
	OrchidColor     lipgloss.Color
	ThistleColor    lipgloss.Color
	NyanzaColor     lipgloss.Color
	AquamarineColor lipgloss.Color
	ErrorColor      lipgloss.Color
	DividerColor    lipgloss.Color

	// Muted colors for unselected items
	MutedTitleColor  lipgloss.Color
	MutedBorderColor lipgloss.Color
	FadedCodeColor   lipgloss.Color

	// Search and highlight colors
	SearchBoxColor   lipgloss.Color
	SearchTextColor  lipgloss.Color
	HighlightBgColor lipgloss.Color
	HighlightFgColor lipgloss.Color

	// Settings colors
	SettingsTitleColor         lipgloss.Color
	SettingsSelectedTitleColor lipgloss.Color
	SettingsBorderColor        lipgloss.Color
	SettingsValueColor         lipgloss.Color
	SettingsEnabledColor       lipgloss.Color
	SettingsDisabledColor      lipgloss.Color
}

func DefaultStyles() *Styles {
	s := new(Styles)

	// Green palette
	s.PeachColor = lipgloss.Color("#A8E6A3")      // Light green
	s.CoralColor = lipgloss.Color("#7FD97F")      // Medium green
	s.OrchidColor = lipgloss.Color("#6BCF6B")     // Emerald green
	s.ThistleColor = lipgloss.Color("#90EE90")    // Mint green
	s.NyanzaColor = lipgloss.Color("#D4F1D4")     // Pale green
	s.ErrorColor = lipgloss.Color("#FF6B6B")      // Red for errors
	s.AquamarineColor = lipgloss.Color("#5FD3A6") // Aqua green
	s.DividerColor = lipgloss.Color("#6B6B6B")

	// Muted colors for unselected items
	s.MutedTitleColor = lipgloss.Color("#6B6B6B")  // Subtle gray
	s.MutedBorderColor = lipgloss.Color("#3A3A3A") // Very dark gray

	// Faded color for code preview
	s.FadedCodeColor = lipgloss.Color("#5A5A5A") // Very muted gray for code

	// Search and highlight colors
	s.SearchBoxColor = s.AquamarineColor
	s.SearchTextColor = s.ThistleColor
	s.HighlightBgColor = lipgloss.Color("#90EE90") // Light green highlight
	s.HighlightFgColor = lipgloss.Color("#1A1A1A") // Dark text for readability

	// Settings colors - green tones
	s.SettingsTitleColor = lipgloss.Color("#8FB98F")         // Muted sage green
	s.SettingsSelectedTitleColor = lipgloss.Color("#A8D5A8") // Soft green
	s.SettingsBorderColor = lipgloss.Color("#6B9B6B")        // Medium sage
	s.SettingsValueColor = lipgloss.Color("#B5D4B5")         // Light sage
	s.SettingsEnabledColor = lipgloss.Color("#7FD97F")       // Bright green
	s.SettingsDisabledColor = lipgloss.Color("#E8999D")      // Soft red

	s.BorderColor = s.OrchidColor
	s.FooterColor = s.NyanzaColor
	s.TitleColor = s.ThistleColor
	s.SelectedTitleColor = s.OrchidColor

	s.InputField = lipgloss.NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(80)

	s.InputFieldWithError = lipgloss.NewStyle().
		BorderForeground(s.ErrorColor).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(80)

	s.FooterStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		Foreground(s.FooterColor).
		Italic(true)

	s.TitleStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		Foreground(s.TitleColor).
		Bold(true)

	s.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	s.HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	s.SelectedItemStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(s.SelectedTitleColor).
		Foreground(s.SelectedTitleColor).
		Padding(0, 0, 0, 1)

	return s
}

func (s Styles) Text(t string, c lipgloss.Color) string {
	var style = lipgloss.NewStyle().Foreground(c)
	return style.Render(t)
}
