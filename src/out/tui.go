package out

import (
	"messh/src/constants"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ASCII Banner
func Banner(message string) string {
	bannerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		PaddingTop(1).
		PaddingBottom(1)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("4"))

	banner := bannerStyle.Render(strings.Trim(constants.AppBanner, "\n"))

	if strings.TrimSpace(message) != "" {
		helpBlock := helpStyle.Render(message)
		return lipgloss.JoinVertical(lipgloss.Left, banner, helpBlock)
	}

	return banner
}

func SectionHeader(header string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("11")). // Bright yellow
		Bold(true).
		Padding(1, 0).
		Render(header)
}

func KV(key, value string) string {
	keyStyle := lipgloss.NewStyle().
		PaddingLeft(2).
		Bold(true).
		Foreground(lipgloss.Color("11")) // Bright Yellow

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")) // Bright White

	return keyStyle.Render(key+":") + " " + valueStyle.Render(value)
}
