package utils

import (
	"math/rand"
	"time"
	"github.com/spf13/viper"
)

func GetGreeting() string {
	vp := viper.New()
	vp.SetConfigName("greeting.yaml")
        vp.SetConfigType("yaml")
        vp.AddConfigPath("data/")

        err := vp.ReadInConfig()
        if err != nil {
		panic(err)
        }

	rand.Seed(time.Now().UnixNano())
	greetings := vp.GetStringSlice("greeting")
	randIndex := rand.Intn(len(greetings))
	result := greetings[randIndex]

	return result
}

