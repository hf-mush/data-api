package repository

// AccessTokenRepository access_token repository
type AccessTokenRepository interface {
	GetUserInfo(userID string) (string, error)
	SetUserInfo(userID string, data interface{}) error
}
