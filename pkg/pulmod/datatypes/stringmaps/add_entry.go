package stringmaps

func AddEntry(m map[string]string, key, value string) map[string]string {
	m[key] = value
	return m
}
