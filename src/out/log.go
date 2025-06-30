package out

import (
	"fmt"
	"messh/src/config"
	"messh/src/helpers"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var SCLogger *log.Logger

type LoggerOptions struct {
	Level     string
	LogToFile bool
	FilePath  string
}

// Helper Functions ////////////////////////////////////////////////

func getLogLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "debug":
		return log.DebugLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	default:
		return log.InfoLevel
	}
}

func getSimpleConsoleLoggerStyles() *log.Styles {
	commonStyle := lipgloss.NewStyle().Bold(true)
	styles := log.DefaultStyles()
	styles.Levels[log.DebugLevel] = commonStyle.SetString("[-]").Foreground(lipgloss.Color("63"))
	styles.Levels[log.InfoLevel] = commonStyle.SetString("[i]").Foreground(lipgloss.Color("86"))
	styles.Levels[log.WarnLevel] = commonStyle.SetString("[!]").Foreground(lipgloss.Color("192"))
	styles.Levels[log.ErrorLevel] = commonStyle.SetString("[✗]").Foreground(lipgloss.Color("204"))
	styles.Levels[log.FatalLevel] = commonStyle.SetString("[◈]").Foreground(lipgloss.Color("134"))
	return styles
}

// Main Functions ////////////////////////////////////////////////

func init() {
	cfg, _, err := config.GetConfig()
	if err != nil {
		fmt.Println("failed to load config:", err)
	}
	SCLogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		TimeFormat:      "2006-01-02 15:04:05",
		Level:           getLogLevel(cfg.AppManagement.LogLevel),
		Formatter:       log.TextFormatter,
	})
	SCLogger.SetStyles(getSimpleConsoleLoggerStyles())
}

func NewLogger(opts ...LoggerOptions) (*log.Logger, error) {

	options := LoggerOptions{
		Level:     "info",
		LogToFile: false,
		FilePath:  "",
	}

	if len(opts) > 0 {
		options = opts[0]
	}

	if !options.LogToFile {
		return log.NewWithOptions(os.Stderr, log.Options{
			ReportTimestamp: true,
			Level:           getLogLevel(options.Level),
			Formatter:       log.TextFormatter,
		}), nil
	}

	const defaultLogFileName = "messh.log"
	const defaultLogDirName = "log"

	if options.FilePath == "" {
		rootPath, err := helpers.ResolveRootPath("")
		if err != nil {
			return nil, fmt.Errorf("failed to app path: %w", err)
		}
		options.FilePath = filepath.Join(rootPath, defaultLogDirName, defaultLogFileName)
	}

	// Check if the path is a dir & make parent dirs if needed
	info, err := os.Stat(options.FilePath)
	switch {
	case err == nil && info.IsDir():
		options.FilePath = filepath.Join(options.FilePath, defaultLogFileName)
	case os.IsNotExist(err):
		if err := os.MkdirAll(filepath.Dir(options.FilePath), 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
	case err != nil:
		return nil, fmt.Errorf("failed to stat log path: %w", err)
	}

	// Enforce .log extension
	if ext := strings.ToLower(filepath.Ext(options.FilePath)); ext != ".log" {
		options.FilePath += ".log"
	}

	// Open the file
	file, err := os.OpenFile(options.FilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Create the logger & return it
	return log.NewWithOptions(file, log.Options{
		TimeFormat:      "2006-01-02 15:04:05",
		Level:           getLogLevel(options.Level),
		ReportTimestamp: true,
		Formatter:       log.TextFormatter,
	}), nil

}
