package payyans

import (
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
)

func Normalize(input string, rulesMap map[string]string) (string, error) {
	expression := strings.Join(maps.Keys(rulesMap), "|")

	pattern, err := regexp.Compile(expression)

	if err != nil {
		return "", err
	}

	some := pattern.ReplaceAllStringFunc(input, func(match string) string {
		return rulesMap[match]
	})

	return some, nil
}

func NormalizeByRulesFile(input string, rulesFile string) (string, error) {
	rulesMap, err := ParseEqualSplittedFile(rulesFile)
	if err != nil {
		return "", err
	}

	return Normalize(input, rulesMap)
}
