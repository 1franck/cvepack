package github

import (
	"cvepack/core/ecosystem"
	"testing"
)

func Test_DetectGithubRepoUrl(t *testing.T) {
	type scenario struct {
		url     string
		isValid bool
		repo    string
	}

	tests := []scenario{
		{"https://github.com/someuser", false, ""},
		{"https://github.com/someuser/repo", true, "someuser/repo"},
		{"github.com/some_user/repo", true, "some_user/repo"},
		{"github.com/someUser1/repo/", true, "someUser1/repo"},
		{"http://github.com/someUser1/repo/", true, "someUser1/repo"},
		{"https://github.com/someUser_1/repo_1/blob/master/README.md", true, "someUser_1/repo_1"},
		{"https://github.com/someuser/repo/issues", true, "someuser/repo"},
	}

	for _, test := range tests {
		source := ecosystem.NewSource(test.url, ecosystem.UrlSource)
		gh := NewUrl(source)
		if gh.IsValid != test.isValid {
			t.Errorf("Expected %s to be %t", test.url, test.isValid)
		}
		if gh.Repo != test.repo {
			t.Errorf("Expected repo to be %s, got %s", test.repo, gh.Repo)
		}
	}
}

func Test_GetFileRawUrl(t *testing.T) {
	type scenario struct {
		url      string
		branch   string
		filepath string
		expected string
	}

	tests := []scenario{
		{"https://github.com/someuser/repo", "main", "README.md", "https://raw.githubusercontent.com/someuser/repo/main/README.md"},
	}

	for _, test := range tests {
		source := ecosystem.NewSource(test.url, ecosystem.UrlSource)
		gh := NewUrl(source)
		got := gh.GetFileRawUrl(test.branch, test.filepath)
		if got != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, got)
		}
	}
}
