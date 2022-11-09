package repository

import (
	"fetcher/internal/config"
	"fmt"
	"io"
	"net/http"
)

type NetworkRepository struct {
	config config.Config
	client *http.Client
}

func NewNetworkRepository(config config.Config, client *http.Client) *NetworkRepository {
	return &NetworkRepository{config: config, client: client}
}

func (n *NetworkRepository) Source(endpoint string) (source string, err error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}

	response, err := n.client.Do(req)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get failed \"%v\" status code %v", endpoint, response.StatusCode)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
