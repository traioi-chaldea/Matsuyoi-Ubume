package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"

	"github.com/TraiOi/notify"
	"github.com/TraiOi/summoning"
)

func main() {
	// Load discord config
	vp := viper.New()
	vp.SetConfigName("discord.yaml")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("config/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Get Bot Token
	token := vp.GetString("TOKEN")

	// Discord bot handler
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	// Add bot features here
	// `dg.AddHandler(<feature>)
	notify.HandleNotify(dg)
	dg.AddHandler(summoning.HandlerRare)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open bot connection
	err = dg.Open()
	if err != nil {
		panic(err)
	}

	// Close bot connection
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

