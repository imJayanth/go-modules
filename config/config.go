package config

import (
	"os"
	"time"

	"github.com/imJayanth/go-modules/helpers"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type mysqlDBConfig struct {
	DB          *gorm.DB
	DBUser      string
	DBPassword  string
	DBHost      string
	DBName      string
	DBPort      string
	Log         bool
	Automigrate []helpers.Model
}

type redisConfig struct {
	Host string
	Port string
	Pool *redis.Pool
}

type serverConfig struct {
	APIPort  int
	Timezone string
}

type AuthConfig struct {
	Username  string
	Password  string
	JwtSecret string
}

type loggerConfig struct {
	ZapLogger *zap.Logger
	Filename  string
	File      *os.File
}

type AppConfig struct {
	ServerConfig  serverConfig
	MySQLDBConfig mysqlDBConfig
	LoggerConfig  loggerConfig
	RedisConfig   redisConfig
	AuthConfig    AuthConfig
	Environment   string
}

func SetupConfig() *AppConfig {
	return &AppConfig{
		AuthConfig: AuthConfig{
			Username:  helpers.GetEnv("USERNAME", "internal"),
			Password:  helpers.GetEnv("PASSWORD", "internal"),
			JwtSecret: helpers.GetEnv("JWT_SECRET", "9iyfuyk0uogfus3456t7yu8gf"),
		},
		MySQLDBConfig: mysqlDBConfig{
			DBUser:     helpers.GetEnv("DBUSER", "root"),
			DBPassword: helpers.GetEnv("DBPASSWORD", "root"),
			DBHost:     helpers.GetEnv("DBHOST", "localhost"),
			DBName:     helpers.GetEnv("DBNAME", ""),
			DBPort:     helpers.GetEnv("DBPORT", "3306"),
		},
		ServerConfig: serverConfig{
			APIPort:  helpers.GetEnvAsInt("APIPORT", 8081),
			Timezone: helpers.GetEnv("TIMEZONE", "Asia/Kolkata"),
		},
		RedisConfig: redisConfig{
			Host: helpers.GetEnv("REDISHOST", "localhost"),
			Port: helpers.GetEnv("REDISPORT", "6379"),
		},
		LoggerConfig: loggerConfig{
			Filename: "",
		},
		Environment: helpers.GetEnv("ENVIRONMENT", "dev"),
	}
}

var loc *time.Location

func (s *AppConfig) CurrentTime() time.Time {
	if loc == nil {
		location, err := time.LoadLocation(s.ServerConfig.Timezone)
		if err != nil {
			panic(err)
		}
		loc = location
	}
	t := time.Now()
	return t.In(loc)
}

func (s *AppConfig) Timezone() *time.Location {
	if loc == nil {
		location, err := time.LoadLocation(s.ServerConfig.Timezone)
		if err != nil {
			panic(err)
		}
		loc = location
	}
	return loc
}
