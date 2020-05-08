package model

// AccessToken access_token struct
type AccessToken struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpires  string `json:"access_token_expires"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpires string `json:"refresh_token_expires"`
}
