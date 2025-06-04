package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
    file, err := os.Open("./config/.env")
	if err != nil {
		fmt.Errorf("failed to open file %w", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// skip comments and new lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// TODO: probably don't allow escape characters
		// research what other characters that shouldn't be allowed

		pair := strings.Split(line, ": ")
		os.Setenv(pair[0], pair[1])
	}
}
