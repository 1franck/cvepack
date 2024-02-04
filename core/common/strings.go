package common

import (
	"fmt"
	"strings"
)

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

func DetectStringLineEnding(content string) string {
	if strings.Contains(content, "\r\n") {
		return "\r\n"
	}

	return "\n"
}

func TextPad(text string, padding int) string {
	pad := strings.Repeat(" ", padding)
	return fmt.Sprintf("%s%s%s", pad, text, pad)
}

func DefaultTextPad(text string) string {
	return TextPad(text, 1)
}
