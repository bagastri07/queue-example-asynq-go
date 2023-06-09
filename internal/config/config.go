package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetConf() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warningf("%v", err)
	}
}

func parseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}

func RedisHost() string {
	return viper.GetString("redis.host")
}

func RedisPassword() string {
	return viper.GetString("redis.password")
}

func RedisCacheTTL() time.Duration {
	cfg := viper.GetString("redis.cache_ttl")
	return parseDuration(cfg, DefaultRedisCacheTTL)
}
