package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/shuufujita/data-api/interfaces/response"
	"github.com/shuufujita/data-api/usecases"
)

// AuthenticationHandler authentication handler
type AuthenticationHandler interface {
	Authentication(next echo.HandlerFunc) echo.HandlerFunc
}

type authenticationHandler struct {
	accessTokenUseCase usecases.AccessTokenUseCase
}

// NewAuthenticationHandler authentication handler entity
func NewAuthenticationHandler(atu usecases.AccessTokenUseCase) AuthenticationHandler {
	return &authenticationHandler{
		accessTokenUseCase: atu,
	}
}

func (ah authenticationHandler) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.ErrorResponse(c, "INVALID_FORMAT", `access token must set header and start with prefix 'Bearer '`)
		}
		accessToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse accessToken to JSON Web Token
		jwtToken, err := ah.accessTokenUseCase.Parse(accessToken)
		if err != nil {
			return response.ErrorResponse(c, "INVALID_TOKEN", err.Error())
		}
		if !ah.accessTokenUseCase.Validate(jwtToken) {
			return response.ErrorResponse(c, "INVALID_TOKEN", "expired or falsified token")
		}

		// Get userID from JSON Web Token
		userID := ah.accessTokenUseCase.GetUserID(jwtToken)
		c.Set("userID", userID)
		colog.FixedValue("userID", userID)

		// Extract userInfo with JSON object
		userInfo, err := ah.accessTokenUseCase.GetUserInfo(userID)
		if err != nil {
			return response.ErrorResponse(c, "CACHE_SERVER_ERROR", "expired or falsified token")
		}
		if userInfo == "" {
			return response.ErrorResponse(c, "LOGIN_USER_CACHE_EXPIRED", "user cache expired.")
		}

		// Convert JSON to map object
		var userInfoMap map[string]interface{}
		err = json.Unmarshal([]byte(userInfo), &userInfoMap)
		if err != nil {
			log.Println(fmt.Sprintf("%v: [%v] %v", "warn", userID, err.Error()))
		}

		return next(c)
	}
}
