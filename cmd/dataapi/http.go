package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shuufujita/data-api/infrastructure/persistance"
	"github.com/shuufujita/data-api/interfaces/handler"
	"github.com/shuufujita/data-api/usecases"
)

// RunServer launch and run server.
func RunServer(port int64) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderCookie,
			echo.HeaderSetCookie,
		},
		AllowMethods: []string{
			echo.GET,
			echo.PUT,
			echo.POST,
			echo.DELETE,
		},
		AllowCredentials: true,
	}))

	trainingRepository := persistance.NewTrainingPersistance()
	trainingUseCase := usecases.NewTrainingUseCase(trainingRepository)
	trainingHandler := handler.NewTrainingHandler(trainingUseCase)

	accessTokenRepository := persistance.NewAccessTokenPersistance()
	accessTokenUseCase := usecases.NewAccessTokenUseCase(accessTokenRepository)
	accessTokenHandler := handler.NewAuthenticationHandler(accessTokenUseCase)

	g := e.Group("/v1", customLogger)
	g.Use(accessTokenHandler.Authentication)
	g.GET("/recorder/training", trainingHandler.RetrieveLogs)
	g.POST("/recorder/training", trainingHandler.CreateLog)
	g.PUT("/recorder/training", trainingHandler.UpdateLog)
	g.DELETE("/recorder/training", trainingHandler.DeleteLog)

	return e.Start(":" + strconv.FormatInt(port, 10))
}

func customLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path, err := getLogFilePath()
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		colog.SetOutput(io.MultiWriter(file, os.Stdout))
		colog.SetFormatter(&colog.StdFormatter{
			Flag: log.Ldate | log.Ltime | log.Lshortfile,
		})
		colog.FixedValue("remoteAddr", c.Request().RemoteAddr)

		return next(c)
	}
}

func getLogFilePath() (string, error) {
	if _, err := os.Stat(os.Getenv("LOG_DIR_PATH")); os.IsNotExist(err) {
		err = os.Mkdir(os.Getenv("LOG_DIR_PATH"), 0777)
		if err != nil {
			return "", err
		}
		log.Println(fmt.Sprintf("%v: [%v] %v", "info", "http", "mkdir with "+os.Getenv("LOG_DIR_PATH")))
	}
	return os.Getenv("LOG_DIR_PATH") + "/" + time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("20060102") + ".log", nil
}
