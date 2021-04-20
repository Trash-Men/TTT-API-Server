package configs

import (
	"fmt"
	"log"
	"os"
)

var env environment

var (
	mode string

	dbName     string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string

	jwtSecretKey string

	s3BucketName string

	serverPort         string
	iamAccessKey       string
	iamSecretAccessKey string
)

type environment struct {
	MODE                  string
	DB_NAME               string
	DB_HOST               string
	DB_PORT               string
	DB_USER               string
	DB_PASSWORD           string
	JWT_SECRET_KEY        string
	S3_BUCKET_NAME        string
	SERVER_PORT           string
	IAM_ACCESS_KEY        string
	IAM_SECRET_ACCESS_KEY string
}

func GetEnvironments() environment {
	if env != (environment{}) {
		return env
	}

	if mode == "" {
		mode = os.Getenv("MODE")

		dbName = os.Getenv("DB_NAME")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")

		jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

		s3BucketName = os.Getenv("S3_BUCKET_NAME")

		serverPort = os.Getenv("SERVER_PORT")

		iamAccessKey = os.Getenv("IAM_ACCESS_KEY")
		iamSecretAccessKey = os.Getenv("IAM_SECRET_ACCESS_KEY")
	}

	if mode != "development" && mode != "production" {
		panic("environment variable 'MODE' must be required with command(e.g., MODE=production go run main.go).")
	}

	if mode == "production" && s3BucketName == "" {
		panic("when in 'production' mode, must be required s3BucketName")
	}

	if dbName == "" {
		environmentRequiredError("DB_NAME")
	}

	if dbHost == "" {
		environmentRequiredError("DB_HOST")
	}

	if dbPort == "" {
		environmentRequiredError("DB_PORT")
	}

	if dbUser == "" {
		environmentRequiredError("DB_USER")
	}

	if dbPassword == "" {
		environmentRequiredError("DB_PASSWORD")
	}

	if jwtSecretKey == "" {
		environmentRequiredError("jwtSecretKey")
	}

	if serverPort == "" {
		environmentRequiredError("SERVER_PORT")
	}

	if iamAccessKey == "" {
		environmentRequiredError("IAM_ACCESS_KEY")
	}

	if iamSecretAccessKey == "" {
		environmentRequiredError("IAM_SECRET_ACCESS_KEY")
	}

	env = environment{
		MODE:                  mode,
		DB_NAME:               dbName,
		DB_HOST:               dbHost,
		DB_PORT:               dbPort,
		DB_USER:               dbUser,
		DB_PASSWORD:           dbPassword,
		JWT_SECRET_KEY:        jwtSecretKey,
		S3_BUCKET_NAME:        s3BucketName,
		SERVER_PORT:           serverPort,
		IAM_ACCESS_KEY:        iamAccessKey,
		IAM_SECRET_ACCESS_KEY: iamSecretAccessKey,
	}

	return env
}

func environmentRequiredError(environmentKey string) {
	log.Panic(fmt.Errorf("<%s> environment key must be required!", environmentKey))
}
