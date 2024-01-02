package github

import (
	"io"
	"net/http"
)

func getUrlContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Print the HTML content as a string
	return string(body), nil
}

//func ScanUrl(url string) []ecosystem.Project {
//    content, err := getUrlContent(url)
//    if err != nil {
//        return []ecosystem.Project{}
//    }
//
//
//    if strings.Contains(content, fmt.Sprintf("title=\"%s\"", golang.GoMod)) {
//
//    }
//    if (golang.GoMod)
//
//
//}
