package usecases

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/shuufujita/data-api/domain/repository"
)

// AccessTokenUseCase usecase of access_token
type AccessTokenUseCase interface {
	Generate(userID string) (AccessToken, error)
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

// AccessToken access_token struct
type AccessToken struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpires  string `json:"access_token_expires"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpires string `json:"refresh_token_expires"`
}

// Generate return jwt token
func (at accessTokenUseCase) Generate(userID string) (AccessToken, error) {
	// アクセストークンの有効期限を取得する
	acExp, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
	}
	// アクセストークンを生成する
	accessTokenString, err := generateTokenString(userID, "access", acExp)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return AccessToken{}, err
	}
	// リフレッシュトークンの有効期限を取得する
	rfExp, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
	}
	// リフレッシュトークンを生成する
	refreshTokenString, err := generateTokenString(userID, "refresh", rfExp)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", userID, err.Error()))
		return AccessToken{}, err
	}
	// レスポンスを生成する
	tokens := AccessToken{
		AccessToken:         accessTokenString,
		AccessTokenExpires:  time.Now().Add(time.Minute * time.Duration(acExp)).Format(time.RFC3339),
		RefreshToken:        refreshTokenString,
		RefreshTokenExpires: time.Now().Add(time.Minute * time.Duration(rfExp)).Format(time.RFC3339),
	}

	return tokens, nil
}

// Parse token文字列をparseする
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
	// 署名と有効期限をチェックする(exp が jwt ライブラリ内での型が float64 になっているので int64 に変換しなおして比較)
	return token.Valid && (int64(token.Claims.(jwt.MapClaims)["exp"].(float64)) > time.Now().Unix())
}

func (at accessTokenUseCase) GetUserID(token *jwt.Token) string {
	return token.Claims.(jwt.MapClaims)["sub"].(string)
}

func (at accessTokenUseCase) GetUserInfo(userID string) (string, error) {
	return at.repository.GetUserInfo(userID)
}

func generateTokenString(userID string, tokenType string, expire int64) (string, error) {
	acToken := jwt.New(jwt.SigningMethodRS256)
	// claimsに値を設定する
	claims := acToken.Claims.(jwt.MapClaims)
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
	return acToken.SignedString(key)
}
