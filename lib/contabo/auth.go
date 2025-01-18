package contabo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type AuthConfig struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GetAccessToken() (string, error) {
	config, err := getAuthConfig()
	if err != nil {
		return "", err
	}

	endpoint := "https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("client_id", config.ClientID)
	data.Set("client_secret", config.ClientSecret)
	data.Set("username", config.Username)
	data.Set("password", config.Password)
	data.Set("grant_type", "password")

	resp, err := http.Post(endpoint,
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var token TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func getAuthConfig() (AuthConfig, error) {
	config := AuthConfig{
		ClientID:     viper.GetString("auth.client_id"),
		ClientSecret: viper.GetString("auth.client_secret"),
		Username:     viper.GetString("auth.api_user"),
		Password:     viper.GetString("auth.api_password"),
	}

	if config.ClientID == "" {
		return AuthConfig{}, fmt.Errorf("auth.client_id not configured")
	}
	if config.ClientSecret == "" {
		return AuthConfig{}, fmt.Errorf("auth.client_secret not configured")
	}
	if config.Username == "" {
		return AuthConfig{}, fmt.Errorf("auth.api_user not configured")
	}
	if config.Password == "" {
		return AuthConfig{}, fmt.Errorf("auth.api_password not configured")
	}

	return config, nil
}
