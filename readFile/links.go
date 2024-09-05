package readFile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Links() (string, error) {
	var builder strings.Builder

	// Open the file
	file, err := os.Open("2_link.txt")
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line) // Trim spaces from the line
		if trimmedLine != "" {                 // Only add non-empty lines
			if builder.Len() > 0 {
				builder.WriteString("\n") // Add a newline separator
			}
			builder.WriteString(trimmedLine)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return builder.String(), nil
}
