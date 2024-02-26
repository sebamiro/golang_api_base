package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	TemplateExt  = ".html"
	StaticDir    = "static"
	StaticPrefix = "files"
)

type environment string

const (
	EnvLocal      environment = "local"
	EnvTest       environment = "test"
	EnvDevelop    environment = "dev"
	EnvStaging    environment = "staging"
	EnvQA         environment = "qa"
	EnvProduction environment = "prod"
)

type (
	Config struct {
		HTTP HTTPConfig
		App  AppConfig
		Data DatabaseConfig
	}

	HTTPConfig struct {
		Hostname     string
		Port         uint16
		ReadTimeout  time.Duration
		WrtieTimeout time.Duration
		IdleTimeout  time.Duration
		TLS          struct {
			Enabled     bool
			Certificate string
			Key         string
		}
	}

	AppConfig struct {
		Name          string
		Environment   environment
		EncryptionKey string
		Timeout       time.Duration
		PasswordToken struct {
			Expiration time.Duration
			Lenght     int
		}
		EmailVerificationTokenExpiration time.Duration
	}

	DatabaseConfig struct {
		Hostname     string
		Port         uint16
		User         string
		Password     string
		Database     string
		TestDatabase string
	}

	// MailConfig stores the mail configuration
	// MailConfig struct {
	// 	Hostname    string
	// 	Port        uint16
	// 	User        string
	// 	Password    string
	// 	FromAddress string
	// }
)

func SwitchEnvironment(env environment) {
	if err := os.Setenv("GOLANG_APP_ENVIRONMENT", string(env)); err != nil {
		panic(err)
	}
}

func GetConfig() (Config, error) {
	var c Config
	log.Println("[TRACE] GetConfig")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")

	viper.SetEnvPrefix("golang")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Println("[FATAL] GetConfig: viper.ReadInConfig: ", err)
		return c, err
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Println("[FATAL] GetConfig: viper.Unmarshal: ", err)
		return c, err
	}
	return c, nil
}
