package grep

import (
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
func search(query, contents string) (result []string) {
	var lines = strings.Split(contents, "\n")
	for _, line := range lines {
		if strings.Contains(line, query) {
			result = append(result, line)
		}
	}

	return
}

// searchCaseInsensitive performs a case-insensitive search
func searchCaseInsensitive(query, contents string) (result []string) {
	var lines = strings.Split(strings.ToLower(contents), "\n")
	for _, line := range lines {
		if strings.Contains(line, strings.ToLower(query)) {
			result = append(result, line)
		}
	}

	return
}
