package env

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

var Env *config

type config struct {
	GoEnv       string `mapstructure:"GO_ENV"`
	GoPort      string `mapstructure:"GO_PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JwtSecret 	string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn int `mapstructure:"JWT_EXPIRES_IN"`
	TokenAutn *jwtauth.JWTAuth
}

func LoadingConfig(path string) (*config, error) {
	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		return nil, err
	}

	Env.TokenAutn = jwtauth.New("HS256", []byte(Env.JwtSecret), nil)
	return Env, nil
}