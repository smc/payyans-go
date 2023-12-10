package main

import (
	"os"
	"strings"
)

func ReadAndCleanFile(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dataInString := string(data)

	mapping := make(map[string]string)

	for _, line := range strings.Split(dataInString, "\n") {
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		lhs, rhs := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		mapping[lhs] = rhs
	}

	return mapping, nil
}
