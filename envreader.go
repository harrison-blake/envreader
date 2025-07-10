package envreader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Load(file string) (err error) {
    content, err := os.Open(file)
	if err != nil {
		fmt.Errorf("failed to open file %w", err)
		return err
	}

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()

		// skip comments and new lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" || strings.HasPrefix(line, "export") {
			continue
		}

		pair := strings.Split(line, "=")
		os.Setenv(pair[0], pair[1])
	}

	return
}
