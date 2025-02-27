package lua

import "strings"

func SingleLine(luaCode string) string {
	// Split the input into lines.
	lines := strings.Split(luaCode, "\n")

	// Trim each line and ignore empty ones.
	var trimmedLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			trimmedLines = append(trimmedLines, trimmed)
		}
	}

	// Join the lines with a space (or any separator you prefer).
	return strings.Join(trimmedLines, " ")
}