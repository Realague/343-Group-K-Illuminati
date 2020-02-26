package config

import (
	"os"
	"time"

	"github.com/lbrulet/gonfig"
)

type Configuration struct {
	Port                             string        `json:"app_port" env:"APP_PORT"`
	TokenEncryptionKey               string        `json:"token_encryption_key" env:"TOKEN_ENCRYPTION_KEY"`
	TokenAuthenticationKey           string        `json:"token_authentication_key" env:"TOKEN_AUTHENTICATION_KEY"`
	Email                            string        `json:"email" env:"EMAIL"`
	EmailPassword                    string        `json:"email_password" env:"EMAIL_PASSWORD"`
	EmailConfirmationEndpoint        string        `json:"email_confirmation_endpoint" env:"EMAIL_CONFIRMATION_ENDPOINT"`
	LoginRedirectionUrl              string        `json:"login_redirection_url" env:"LOGIN_REDIRECTION_URL"`
	RecoverPasswordUrl               string        `json:"recover_password_url" env:"RECOVER_PASSWORD_URL"`
	DatabaseName                     string        `json:"database_name" env:"DATABASE_NAME"`
	DatabaseHost                     string        `json:"database_host" env:"DATABASE_HOST"`
	IsUnitTest                       bool          `json:"is_unit_test" env:"IS_UNIT_TEST"`
	RecoverPasswordTokenValidityTime time.Duration `json:"recover_password_token_validity_time" env:"RECOVER_PASSWORD_TOKEN_VALIDITY_TIME"`
	AccessTokenValidityTime          time.Duration `json:"access_token_validity_time" env:"ACCESS_TOKEN_VALIDITY_TIME"`
	RefreshTokenValidityTime         time.Duration `json:"refresh_token_validity_time" env:"REFRESH_TOKEN_VALIDITY_TIME"`
}

var Config Configuration

func InitConfig() {
	var err error
	var file = func() string {
		if os.Getenv("IS_UNIT_TEST") == "true" {
			return "../config"
		}
		return "config"
	}()
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		err = gonfig.GetConf(file+"/production/config.json", &Config)
	} else if os.Getenv("ENVIRONMENT") == "LOCAL" {
		err = gonfig.GetConf(file+"/local/config.json", &Config)
	} else {
		err = gonfig.GetConf(file+"/dev/config.json", &Config)
	}
	if err != nil {
		panic(err)
	}

}
