package readFile

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Proxies() ([]string, error) {

	var proxyData []string

	// Open the file
	file, err := os.Open("1_proxies.txt")
	if err != nil {
		x := fmt.Sprintf("Error opening file:", err)
		return proxyData, errors.New(x)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line) // Trim spaces from the line
		if trimmedLine != "" {                 // Only add non-empty lines
			proxyData = append(proxyData, trimmedLine)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		x := fmt.Sprintf("Error reading file:", err)
		return proxyData, errors.New(x)
	}

	/* Print the proxy data to verify
	fmt.Println("Proxies read from file:")
	for _, proxy := range proxyData.IP {
		fmt.Println(proxy)
	}
	*/
	if len(proxyData) < 1 {
		x := fmt.Sprintf("No proxies added?")
		return proxyData, errors.New(x)
	}

	/*
		for _, proxy := range proxyData {
			fmt.Println(proxy)
		}
	*/

	return proxyData, nil
}
