package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func loadDotEnv() error {
	env := os.Getenv("EXEC_ENV")
	if "" == env {
		env = "local"
	}

	path := getRootPath() + "/.env.local"
	if env == "production" {
		path = getRootPath() + "/.env"
	}

	err := godotenv.Load(path)
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "LoadDotEnv", err.Error()))
		return err
	}
	return nil
}

func getRootPath() string {
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(fmt.Sprintf("%v: [%v] %v", "error", "filepath", err.Error()))
	}
	return root + "/.."
}
