package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	MongoDBURI             string `mapstructure:"MONGODB_URI"`
	MongoDBName            string `mapstructure:"MONGODB_NAME"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	DBLogMode              int    `mapstructure:"DB_LOG_MODE"`
}

var E Env;

func NewEnv() {
	E = Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&E)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if E.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
}
