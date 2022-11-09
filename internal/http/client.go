package http

import (
	"fetcher/internal/config"
	"net/http"
)

func NewHttpClient(config config.Config) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: config.MaxIdleConnsPerHost(),
		},
		Timeout: config.RequestTimeout(),
	}

	return client
}
