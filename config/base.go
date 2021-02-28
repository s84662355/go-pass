package config

import (
	"github.com/joho/godotenv"
	configenv "GoPass/config/env"
	   "github.com/timest/env"
)

var RedisConfig = &configenv.Redis{}

func init(){
	godotenv.Load(".env")
    env.Fill(RedisConfig)
}
