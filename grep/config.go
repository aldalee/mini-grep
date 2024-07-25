package grep

import (
	"errors"
	"os"
)

// Config holds the configuration for the search
type Config struct {
	Query         string
	FileName      string
	CaseSensitive bool
}

// NewConfig creates a new Config from the command line arguments
func NewConfig(args []string) (*Config, error) {
	if len(args) < 3 {
		return nil, errors.New("not enough arguments")
	}

	return &Config{
		Query:         args[1],
		FileName:      args[2],
		CaseSensitive: os.Getenv("CASE_INSENSITIVE") == "",
	}, nil
}
