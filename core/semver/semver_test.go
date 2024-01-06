package semver

import (
	"testing"
)

func Test_IsVersionAffectedByFixedVersion(T *testing.T) {
	type scenario struct {
		version      string
		minVersion   string
		fixedVersion string
		expected     bool
	}

	scenarios := []scenario{
		{"1.0.0", "0", "5.0.0", true},
		{"1.0.0", "0", "1.0.0", false},
		{"1.0.0", "1.0.0", "1.0.1", true},
		{"1.0.0", "1.0.0", "2.0.2", true},
		{"1.0.0", "2.0.0", "2.0.3", false},
		{"1.0.0", "2.0.0", "2.0.3", false},
		{"5.1.2", "0", "5.1.2", false},
	}

	for _, s := range scenarios {
		result := IsVersionAffectedByFixedVersion(s.version, s.minVersion, s.fixedVersion)
		if result != s.expected {
			T.Errorf("IsVersionInRange(%s, %s, %s) failed, expected %t, got %t", s.version, s.minVersion, s.fixedVersion, s.expected, result)
		}
	}
}

func Test_IsVersionInRange(T *testing.T) {
	type scenario struct {
		version    string
		minVersion string
		maxVersion string
		expected   bool
	}

	scenarios := []scenario{
		{"1.0.0", "0", "5.0.0", true},
		{"1.0.0", "1.0.0", "1.0.0", true},
		{"1.0.0", "1.0.0", "1.0.1", true},
		{"1.0.0", "1.0.0", "2.0.2", true},
		{"1.0.0", "2.0.0", "2.0.3", false},
		{"1.0.0", "2.0.0", "2.0.3", false},
		{"5.1.2", "0", "5.1.2", true},
	}

	for _, s := range scenarios {
		result := IsVersionInRange(s.version, s.minVersion, s.maxVersion)
		if result != s.expected {
			T.Errorf("IsVersionInRange(%s, %s, %s) failed, expected %t, got %t", s.version, s.minVersion, s.maxVersion, s.expected, result)
		}
	}
}
