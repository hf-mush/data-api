package usecases

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func before() error {
	path := os.Getenv("DOTENV_PATH")
	err := godotenv.Load(path)
	if err != nil {
		log.Println("load .env from " + path)
		return err
	}
	return nil
}

var userID = "100"

func TestGenerateToken(t *testing.T) {
	err := before()
	if err != nil {
		t.Fatalf("%v", "load .env error. "+err.Error())
	}

	now := time.Now().Unix()
	expireMinutes, err := getAccessTokenExpire(now)
	if err != nil {
		t.Fatalf("%v", "get expire error. "+err.Error())
	}

	_, err = generateTokenString(userID, now, expireMinutes)
	if err != nil {
		t.Fatalf("%v", "generate token error. "+err.Error())
	}
}

func TestGenerateAccessToken(t *testing.T) {
	err := before()
	if err != nil {
		t.Fatalf("%v", "load .env error. "+err.Error())
	}

	atu := &accessTokenUseCase{}
	_, err = atu.Generate(userID)
	if err != nil {
		t.Fatalf("%v", "access token generate error. "+err.Error())
	}
}

func TestParseToken(t *testing.T) {
	err := before()
	if err != nil {
		t.Fatalf("%v", "load .env error. "+err.Error())
	}

	now := time.Now().Unix()
	expireMinutes, err := getAccessTokenExpire(now)
	if err != nil {
		t.Fatalf("%v", "get expire error. "+err.Error())
	}

	token, err := generateTokenString(userID, now, expireMinutes)
	if err != nil {
		t.Fatalf("%v", "generate token error. "+err.Error())
	}

	atu := &accessTokenUseCase{}
	_, err = atu.Parse(token)
	if err != nil {
		t.Fatalf("%v", "token parse error : "+err.Error())
	}
}

func TestValidateToken(t *testing.T) {
	err := before()
	if err != nil {
		t.Fatalf("%v", "load .env error. "+err.Error())
	}

	now := time.Now().Unix()
	expireMinutes, err := getAccessTokenExpire(now)
	if err != nil {
		t.Fatalf("%v", "get expire error. "+err.Error())
	}

	token, err := generateTokenString(userID, now, expireMinutes)
	if err != nil {
		t.Fatalf("%v", "generate token error. "+err.Error())
	}

	atu := &accessTokenUseCase{}
	jwtToken, err := atu.Parse(token)
	if err != nil {
		t.Fatalf("%v", "token parse error. "+err.Error())
	}

	valid := atu.Validate(jwtToken)
	assert.True(t, valid)
}

func TestGetUserID(t *testing.T) {
	err := before()
	if err != nil {
		t.Fatalf("%v", "load .env error. "+err.Error())
	}

	now := time.Now().Unix()
	expireMinutes, err := getAccessTokenExpire(now)
	if err != nil {
		t.Fatalf("%v", "get expire error. "+err.Error())
	}

	token, err := generateTokenString(userID, now, expireMinutes)
	if err != nil {
		t.Fatalf("%v", "generate token error. "+err.Error())
	}

	atu := &accessTokenUseCase{}
	jwtToken, err := atu.Parse(token)
	if err != nil {
		t.Fatalf("%v", "token parse error. "+err.Error())
	}

	valid := atu.Validate(jwtToken)
	assert.True(t, valid)

	tokenUserID := atu.GetUserID(jwtToken)
	assert.Equal(t, tokenUserID, userID)
}
