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

func Test_Plural(t *testing.T) {
	singular := "package"
	plural := "packages"
	result := Plural(1, singular, plural)
	if result != singular {
		t.Error("Expected singular, got ", result)
	}
	result = Plural(2, singular, plural)
	if result != plural {
		t.Error("Expected plural, got ", result)
	}
}
