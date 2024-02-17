package payyans

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
