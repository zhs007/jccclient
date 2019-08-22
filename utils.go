package jccclient

import "net/url"

// GetHostName - get hostname in url
func GetHostName(strurl string) (string, error) {
	u, err := url.Parse(strurl)
	if err != nil {
		return "", err
	}

	return u.Host, nil
}
