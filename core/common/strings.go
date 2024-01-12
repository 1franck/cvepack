package common

import "strings"

func ReplacePlaceholders(template string, replacements map[string]string) string {
	for placeholder, value := range replacements {
		placeholder = "{" + placeholder + "}"
		template = strings.Replace(template, placeholder, value, -1)
	}
	return template
}

func Plural(count int, singular, plural string) string {
	if count > 1 {
		return plural
	}
	return singular
}
