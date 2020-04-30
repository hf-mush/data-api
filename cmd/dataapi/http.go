package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/comail/colog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	return e.Start(":" + strconv.FormatInt(port, 10))
}

func customLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := getLogFilePath()
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

func getLogFilePath() string {
	return os.Getenv("LOG_DIR_PATH") + "/" + time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("20060102") + ".log"
}
