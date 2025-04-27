package scrapper

import (
	"fmt"
	"io"
	"net/http"
)

func ReadUrl(url string) {
	response, err := makeRequest(url)

	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	body, err := parseResponse(response)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	fmt.Println(body)

}

func makeRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func parseResponse(response *http.Response) (string, error) {
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", fmt.Errorf("status code: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	stringBody := string(body)

	return stringBody, nil
}
