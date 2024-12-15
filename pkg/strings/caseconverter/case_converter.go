package caseconverter

import "strings"

// ToSnakeCase converts a camelCase or PascalCase string into snake_case.
func ToSnakeCase(s string) string {
	var snake []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			snake = append(snake, '_', r+('a'-'A'))
		} else {
			snake = append(snake, r)
		}
	}
	return strings.ToLower(string(snake))
}
