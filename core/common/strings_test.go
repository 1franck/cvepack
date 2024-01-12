package common

import "testing"

func Test_ReplacePlaceholders(t *testing.T) {
	template := "package {package}"
	replacements := map[string]string{
		"package": "common",
	}
	result := ReplacePlaceholders(template, replacements)
	if result != "package common" {
		t.Error("Expected 'package common', got ", result)
	}
}
