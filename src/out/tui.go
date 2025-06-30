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
