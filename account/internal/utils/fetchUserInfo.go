package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchUserInfo(accessToken, provider string) (string, error) {
	var apiUrl string

	// Set the API URL based on the provider
	switch provider {
	case "google":
		apiUrl = "https://www.googleapis.com/oauth2/v3/userinfo"
	case "facebook":
		apiUrl = "https://graph.facebook.com/v12.0/me"
	default:
		return "", fmt.Errorf("unsupported provider: %s", provider)
	}

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return "", err
	}

	// Set the access token in the request header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch user info. Status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
