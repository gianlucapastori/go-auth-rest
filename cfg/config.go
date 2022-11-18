package cfg

import (
	"github.com/spf13/viper"
)

type Config struct {
	SERVER   Server
	DATABASE Database
}

type Server struct {
	PORT    string
	VERSION string
	ENV     string
	JWT     Jwt
}

type Jwt struct {
	REFRESH_COOKIE_NAME string
	REFRESH_EXPIRES_AT  string
	ACCESS_EXPIRES_AT   string
	SECRET_KEY          string
}

type Database struct {
	PORT   string
	HOST   string
	NAME   string
	PSWD   string
	USER   string
	DRIVER string
}

func New(name string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("cfg")
	v.SetConfigName(name)
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}

func Parse(v *viper.Viper) (*Config, error) {
	var cfg *Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
