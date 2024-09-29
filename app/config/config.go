package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type (
	Config struct {
		Server *Server
		Db     *Db
	}

	Server struct {
		Port            int
		CORS            []string
		ServiceUsername string
		ServicePassword string
		OAuthKey        string
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func NewConfig() *Config {

	server := &Server{
		Port:            8081,
		CORS:            strings.Split(os.Getenv("CORS"), ","),
		ServiceUsername: os.Getenv("SERVICE_USER"),
		ServicePassword: os.Getenv("SERVICE_PWD"),
		OAuthKey:        os.Getenv("OAUTH_KEY"),
	}

	DBPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("DBPort is not number")
	}
	db := &Db{
		Host:     os.Getenv("DB_HOST"),
		Port:     DBPort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return &Config{
		Server: server,
		Db:     db,
	}
}

func GetConfig() *Config {
	once.Do(func() {

		configInstance = NewConfig()
	})

	return configInstance
}
