package gopayyans

import (
	"os"
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
)

// TODO: move this function in a single file where the root can reuse it.
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

func Normalize(text string) (string, error) {
	mapping, err := ReadAndCleanFile("../rules/normalizer_ml.rules")
	if err != nil {
		return "", err
	}

	expression := strings.Join(maps.Keys(mapping), "|")

	pattern, err := regexp.Compile(expression)

	if err != nil {
		return "", err
	}

	some := pattern.ReplaceAllStringFunc(text, func(match string) string {
		return mapping[match]
	})

	return some, nil
}
