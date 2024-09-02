package writefile

import (
	"bufio"
	"fmt"
	"os"
)

// SaveToFile writes each string from the slice to the specified file, one per line
func Array(data []string) error {

	//fmt.Println("save this string ", data)
	// Open or create the file
	file, err := os.OpenFile("0_results.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a writer to write to the file
	writer := bufio.NewWriter(file)

	// Write each string to the file
	for _, line := range data {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	// Ensure all buffered data is written to the file
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing buffer: %w", err)
	}

	return nil
}

// SaveString appends a single string to the specified file, followed by a newline
func String(data string) error {
	// Open or create the file
	file, err := os.OpenFile("0_results.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a writer to write to the file
	writer := bufio.NewWriter(file)

	// Write the string to the file followed by a newline
	_, err = writer.WriteString(data + "\n")
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	// Ensure all buffered data is written to the file
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error flushing buffer: %w", err)
	}

	return nil
}
