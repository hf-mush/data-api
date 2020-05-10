package usecases

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/shuufujita/data-api/domain/model"
	"github.com/shuufujita/data-api/domain/repository"
)

// AccessTokenUseCase usecase of access_token
type AccessTokenUseCase interface {
	Generate(userID string) (model.AccessToken, error)
	Parse(tokenString string) (*jwt.Token, error)
	Validate(token *jwt.Token) bool
	GetUserID(token *jwt.Token) string
	GetUserInfo(userID string) (string, error)
	SaveUserInfo(userID string, data interface{}) error
}

type accessTokenUseCase struct {
	repository repository.AccessTokenRepository
}

// NewAccessTokenUseCase return access_token usecase entity
func NewAccessTokenUseCase(atr repository.AccessTokenRepository) AccessTokenUseCase {
	return &accessTokenUseCase{
		repository: atr,
	}
}

func (at accessTokenUseCase) Generate(userID string) (model.AccessToken, error) {
	now := time.Now().Unix()

	accessTokenExpire, err := getAccessTokenExpire(now)
	if err != nil {
		return model.AccessToken{}, err
	}

	accessTokenString, err := generateTokenString(userID, now, accessTokenExpire)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return model.AccessToken{}, err
	}

	refreshTokenExpire, err := getRefreshTokenExpire(now)
	if err != nil {
		return model.AccessToken{}, err
	}

	refreshTokenString, err := generateTokenString(userID, now, refreshTokenExpire)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return model.AccessToken{}, err
	}

	tokens := model.AccessToken{
		AccessToken:         accessTokenString,
		AccessTokenExpires:  time.Unix(accessTokenExpire, 0).Format(time.RFC3339),
		RefreshToken:        refreshTokenString,
		RefreshTokenExpires: time.Unix(refreshTokenExpire, 0).Format(time.RFC3339),
	}

	return tokens, nil
}

func getAccessTokenExpire(now int64) (int64, error) {
	tokenExpireMinutes, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Unix(now, 0).Add(time.Minute * time.Duration(tokenExpireMinutes)).Unix(), nil
}

func getRefreshTokenExpire(now int64) (int64, error) {
	tokenExpireMinutes, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Unix(now, 0).Add(time.Minute * time.Duration(tokenExpireMinutes)).Unix(), nil
}

func generateTokenString(userID string, issuedAt int64, expire int64) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    "DATA-API",
		Subject:   userID,
		Audience:  "DATA-API",
		IssuedAt:  issuedAt,
		NotBefore: issuedAt,
		ExpiresAt: expire,
	}

	keyData, err := ioutil.ReadFile(os.Getenv("PRYVATE_KEY_PATH"))
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "privateKey", err.Error()))
		return "", nil
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "privateKey", err.Error()))
		return "", nil
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return jwtToken.SignedString(privateKey)
}

func (at accessTokenUseCase) Parse(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		keyData, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_PATH"))
		if err != nil {
			log.Println(fmt.Sprintf("%v: [%v] %v", "error", "publicKey", err.Error()))
			return nil, err
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
		if err != nil {
			log.Println(fmt.Sprintf("%v: [%v] %v", "error", "publicKey", err.Error()))
			return nil, err
		}
		return key, nil
	})
}

func (at accessTokenUseCase) Validate(token *jwt.Token) bool {
	return token.Valid && (int64(token.Claims.(jwt.MapClaims)["exp"].(float64)) > time.Now().Unix())
}

func (at accessTokenUseCase) GetUserID(token *jwt.Token) string {
	return token.Claims.(jwt.MapClaims)["sub"].(string)
}

func (at accessTokenUseCase) GetUserInfo(userID string) (string, error) {
	return at.repository.GetUserInfo(userID)
}

func (at accessTokenUseCase) SaveUserInfo(userID string, data interface{}) error {
	return at.repository.SetUserInfo(userID, data)
}
