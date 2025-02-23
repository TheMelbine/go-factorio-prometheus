package setup

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func Logger() {
	level := os.Getenv("LOG_LEVEL")
	format := os.Getenv("LOG_FORMAT")
	logOptions := log.Options{
		TimeFormat:      time.DateTime,
		ReportCaller:    true,
		ReportTimestamp: true,
		Formatter:       log.TextFormatter,
	}

	// level
	if level != "" {
		l, err := log.ParseLevel(level)
		if err != nil {
			log.Fatal("invalid log level", "error", err)
		}
		logOptions.Level = l
	}

	// format
	switch format {
	default:
		log.Warn("unknown log format, falling back to text, expected `text`, `json` or `logfmt`", "format", format)
		fallthrough
	case "text", "":
		logOptions.Formatter = log.TextFormatter
	case "json":
		logOptions.Formatter = log.JSONFormatter
	case "logfmt":
		logOptions.Formatter = log.LogfmtFormatter
	}

	// Initialize the default logger.
	logger := log.NewWithOptions(os.Stderr, logOptions)
	logger.SetStyles(CreateStyle())

	log.SetDefault(logger)

	logger.Debug("setup the logger", "level", level, "format", format)
}

func CreateStyle() *log.Styles {
	styles := log.DefaultStyles()

	styles.Levels[log.DebugLevel] = styles.Levels[log.DebugLevel].SetString("DEBUG")
	styles.Levels[log.InfoLevel] = styles.Levels[log.InfoLevel].SetString("INFO")
	styles.Levels[log.WarnLevel] = styles.Levels[log.WarnLevel].SetString("WARN")
	styles.Levels[log.ErrorLevel] = styles.Levels[log.ErrorLevel].SetString("ERROR")
	styles.Levels[log.FatalLevel] = styles.Levels[log.FatalLevel].SetString("FATAL")

	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Keys["error"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["error"] = lipgloss.NewStyle().Bold(true)
	styles.Values["error"] = lipgloss.NewStyle().Bold(true)

	return styles
}
