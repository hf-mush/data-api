package persistance

import (
	"github.com/shuufujita/data-api/domain/repository"
)

type accessTokenPersistance struct{}

// NewAccessTokenPersistance access_token repository persistance
func NewAccessTokenPersistance() repository.AccessTokenRepository {
	return &accessTokenPersistance{}
}

func (atp accessTokenPersistance) GetUserInfo(userID string) (string, error) {
	return RedisGet("USER:" + userID + ":user_info")
}

func (atp accessTokenPersistance) SetUserInfo(userID string, data interface{}) error {
	return RedisSetJSON("USER:"+userID+":user_info", data)
}
