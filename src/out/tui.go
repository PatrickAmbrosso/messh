package out

import (
	"fmt"
	"messh/src/config"
	"strings"

	"github.com/pterm/pterm"
)

// ASCII Banner
func Banner(message string) string {
	var sb strings.Builder
	sb.WriteString(pterm.Sprintf(config.AppBanner))
	sb.WriteString("\n")
	sb.WriteString(pterm.FgBlue.Sprint(message))
	sb.WriteString("\n")
	return sb.String()
}

// Log Messages
func Debug(msg string, args ...any) {
	style := pterm.FgLightMagenta
	logMessage(style, "[-]", msg, args...)
}

func Info(msg string, args ...any) {
	style := pterm.FgLightGreen
	logMessage(style, "[i]", msg, args...)
}

func Warn(msg string, args ...any) {
	style := pterm.FgYellow
	logMessage(style, "[!]", msg, args...)
}

func Error(msg string, args ...any) {
	style := pterm.FgLightRed
	logMessage(style, "[✗]", msg, args...)
}

func logMessage(color pterm.Color, icon, msg string, args ...any) {
	var b strings.Builder

	emphasisStyle := pterm.NewStyle(color, pterm.Bold)
	messageStyle := pterm.NewStyle(pterm.FgLightWhite)

	b.WriteString(emphasisStyle.Sprint(icon))
	b.WriteString(" ")
	b.WriteString(messageStyle.Sprint(msg))

	// Key-value pairs
	if len(args) > 0 {
		b.WriteString("  ")
		for i := 0; i < len(args); i += 2 {
			var key any = ""
			var val any = ""
			key = args[i]
			if i+1 < len(args) {
				val = args[i+1]
			}
			// Color the key
			b.WriteString(emphasisStyle.Sprint(fmt.Sprintf("%v", key)))
			b.WriteString(messageStyle.Sprint(fmt.Sprintf("=%v ", val)))
		}
	}
	fmt.Println(b.String())
}
