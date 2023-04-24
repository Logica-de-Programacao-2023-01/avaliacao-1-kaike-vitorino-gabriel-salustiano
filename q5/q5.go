package q5

import "strings"

func ProcessString(s string) string {
	var resultado strings.Builder
	vogal := "AEIOUaeiou"
	for _, let := range s {
		if strings.ContainsRune(vogal, let) {
			continue
		} else if let >= 'A' && let <= 'Z' {
			resultado.WriteString(".")
			resultado.WriteString(strings.ToLower(string(let)))

		} else {
			resultado.WriteString(".")
			resultado.WriteString(string(let))

		}
	}
	return resultado.String()
}
