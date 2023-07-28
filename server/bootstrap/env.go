package bootstrap

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AsyncPoolCount         int    `mapstructure:"ASYNC_POOL_COUNT"`
	SubfinderPath          string `mapstructure:"SUBFINDER_PATH"`
	NaabuPath              string `mapstructure:"NAABU_PATH"`
	HttpxPath              string `mapstructure:"HTTPX_PATH"`
}

func NewEnv() *Env {
	_, goMod, _, _ := runtime.Caller(0)
	root := filepath.Dir(filepath.Dir(goMod))
	log.Printf("root: %s", root)
	env := Env{}
	conffile := fmt.Sprintf("%s/.env", root)
	viper.SetConfigFile(conffile)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)

	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
