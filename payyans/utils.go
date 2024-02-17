package payyans

import (
	"strings"
)

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func keyExists(m map[string]bool, key string) bool {
	_, ok := m[key]
	return ok
}

func ParseEqualSplittedData(data string) (map[string]string, error) {
	mapping := make(map[string]string)

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		lhs, rhs := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		mapping[lhs] = rhs
	}

	return mapping, nil
}
