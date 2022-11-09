package utility

import (
	"net/url"
)

func HostNameFromEndpoint(endpoint string) (string, error) {
	url, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	return url.Hostname(), nil
}

func IsValidURL(endpoint string) bool {
	u, err := url.Parse(endpoint)
	return err == nil && u.Scheme != "" && u.Host != ""
}
