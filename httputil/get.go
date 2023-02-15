package httputil

import (
	"io"
	"net/http"
)

func Get(url string) (int, string) {
	resp, err := http.Get(url)
	if err != nil {
		return resp.StatusCode, ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}
