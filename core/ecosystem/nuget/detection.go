package nuget

import (
	"os"
	"strings"
)

func DetectSln(path string) string {
	files, err := os.ReadDir(path)
	if err != nil {
		return ""
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sln") {
			return file.Name()
		}
	}

	return ""
}
