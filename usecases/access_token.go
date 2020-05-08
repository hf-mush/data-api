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
	acExp, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
	}
	accessTokenString, err := generateTokenString(userID, "access", acExp)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return model.AccessToken{}, err
	}

	rfExp, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
	}
	refreshTokenString, err := generateTokenString(userID, "refresh", rfExp)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return model.AccessToken{}, err
	}

	tokens := model.AccessToken{
		AccessToken:         accessTokenString,
		AccessTokenExpires:  time.Now().Add(time.Minute * time.Duration(acExp)).Format(time.RFC3339),
		RefreshToken:        refreshTokenString,
		RefreshTokenExpires: time.Now().Add(time.Minute * time.Duration(rfExp)).Format(time.RFC3339),
	}

	return tokens, nil
}

func (at accessTokenUseCase) Parse(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		keyData, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_PATH"))
		if err != nil {
			log.Println(fmt.Sprintf("%v: [%v] %v", "error", "-", err.Error()))
			return nil, err
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
		if err != nil {
			log.Println(fmt.Sprintf("%v: [%v] %v", "error", "-", err.Error()))
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

func generateTokenString(userID string, tokenType string, expire int64) (string, error) {
	jwtToken := jwt.New(jwt.SigningMethodRS256)
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expire)).Unix()
	claims["iss"] = "DATA-API"
	claims["iat"] = time.Now().Unix()
	claims["token_type"] = tokenType
	keyData, err := ioutil.ReadFile(os.Getenv("PRYVATE_KEY_PATH"))
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return "", err
	}
	return jwtToken.SignedString(key)
}
