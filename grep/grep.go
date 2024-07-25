package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Run executes the search based on the provided configuration
func Run(config *Config) error {
	contents, err := os.ReadFile(config.FileName)
	if err != nil {
		return err
	}

	var results []string
	if config.CaseSensitive {
		results = search(config.Query, string(contents))
	} else {
		results = searchCaseInsensitive(config.Query, string(contents))
	}

	for _, line := range results {
		fmt.Println(line)
	}

	return nil
}

// search performs a case-sensitive search
func search(query, contents string) (results []string) {
	scanner := bufio.NewScanner(strings.NewReader(contents))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, query) {
			results = append(results, line)
		}
	}
	return results
}

// searchCaseInsensitive performs a case-insensitive search
func searchCaseInsensitive(query, contents string) (results []string) {
	lowerQuery := strings.ToLower(query)
	scanner := bufio.NewScanner(strings.NewReader(contents))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), lowerQuery) {
			results = append(results, line)
		}
	}
	return results
}
