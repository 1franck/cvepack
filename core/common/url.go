package common

import (
	"fmt"
	"io"
	"net/http"
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
