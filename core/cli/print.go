package cli

import (
	"cvepack/core/config"
	"fmt"
	"sort"
	"strings"
)

func PrintNameWithVersionHeader() {
	PrintWithUnderline(fmt.Sprintf("%s v%s", config.Default.Name, config.Default.Version))
}

func PrintWithUnderline(text string) {
	fmt.Println(text)
	fmt.Println(strings.Repeat("=", len(text)))
}

func PrintMap(m map[string]string, keyValFormatter func(string, string) (string, string)) {
	longestKey := 0
	longestValue := 0 // no needed for now, maybe remove or use later with border

	processedMap := make(map[string]string, len(m))
	for k, v := range m {
		key := k
		val := v
		if keyValFormatter != nil {
			key, val = keyValFormatter(k, v)
		}

		if len(key) > longestKey {
			longestKey = len(key)
		}
		if len(val) > longestValue {
			longestValue = len(val)
		}

		processedMap[key] = val
	}

	keys := make([]string, 0, len(processedMap))
	for k := range processedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := processedMap[key]
		padding := longestKey - len(key)
		paddingString := ""
		if padding > 0 {
			paddingString = strings.Repeat(".", padding)
		}
		fmt.Printf("%s %s.... %s\n", key, paddingString, val)
	}
}
