package common

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func DownloadUrlContent(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to download %s. Status code: %d", url, response.StatusCode)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func UrlExists(url string, timeLimit int) bool {
	if timeLimit <= 0 {
		timeLimit = 5
	}
	client := http.Client{
		Timeout: time.Duration(timeLimit) * time.Second,
	}
	response, err := client.Head(url)
	if err != nil {
		return false
	}
	if response.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func FirstUrlExists(urls []string, timeLimit int) string {
	for _, url := range urls {
		if UrlExists(url, timeLimit) {
			return url
		}
	}
	return ""
}
