package main

import (
	"fmt"
	"mini-grep/grep"
	"os"
)

func main() {
	config, err := grep.NewConfig(os.Args)
	if err != nil {
		fmt.Printf("problem parsing arguments: %v", err)
		return
	}

	if err = grep.Run(config); err != nil {
		fmt.Printf("application error: %v", err)
	}
}
