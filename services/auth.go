package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Oskar-jansson/m5adminapi/internal/client"
	"github.com/Oskar-jansson/m5adminapi/models"
)

type AuthService struct {
	client *client.Client
}

func NewAuthService(c *client.Client) *AuthService {
	return &AuthService{client: c}
}

// Login authenticates the user with the stored credentials and retrieves an access token.
// The access token is automatically stored in the client for subsequent authenticated requests.
func (s *AuthService) Login(ctx context.Context) *models.SdkError {

	url := fmt.Sprintf("%s/login", s.client.Path)

	body, err := json.Marshal(s.client.Credentials)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	rs := client.RequestSettings{
		URL:                url,
		Method:             "POST",
		Body:               body,
		IncludeAccessToken: false,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}

	// Expected login response
	type loginResponse struct {
		AccessToken string `json:"accesstoken"`
	}

	r, sdkErr := client.ResponseConvert[loginResponse](resp)
	if r != nil {
		s.client.SetAccessToken(r.AccessToken)
	}
	if sdkErr != nil {
		return sdkErr
	}

	return nil

}

func (s *AuthService) Logout(ctx context.Context) *models.SdkError {
	rs := client.RequestSettings{
		URL:                "/m5adminapi/api/logout",
		Method:             "GET",
		IncludeAccessToken: true,
	}

	resp, err := s.client.HttpRequest(ctx, rs)
	if err != nil {
		return &models.SdkError{Err: err}
	}
	type logoutResponse struct{}

	_, sdkErr := client.ResponseConvert[logoutResponse](resp)
	if sdkErr != nil {
		return sdkErr
	}

	return nil
}

// DiscardToken clears the stored access token by setting it to an empty string.
// NOTE: only discards token, does not perform logout from api.
func (s *AuthService) DiscardToken() {
	s.client.SetAccessToken("")
}

// SetAccesstoken updates the stored access token with the provided value.
// NOTE: only discards token, does not perform logout from api.
func (s *AuthService) SetAccesstoken(newToken string) {
	s.client.SetAccessToken(newToken)
}

// Returns stored accesstoken in client object.
func (s *AuthService) GetAccesstoken() string {
	return s.client.GetAccessToken()
}
