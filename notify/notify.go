package notify

import (
	"time"
	"strings"
	dgo "github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func HandleNotify(s *dgo.Session, m *dgo.MessageCreate) {
	vp := viper.New()
	vp.SetConfigName("discord.yaml")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("config/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	channel := vp.GetString("ChannelNotify")

	if m.Author.ID == s.State.User.ID {
		return
	}

	// Get current datetime
	t := time.Now()
	t.Format("Mon Jan _2 15:04:05 2006")
	weekday := strings.ToLower(t.Weekday().String())

	// Get daily notify feature
	daily := new(DailyFunc)
	daily.Init(weekday)

	if m.Content == "test" {
		tmp := daily.getFrom7H()
		if tmp != "empty" {
			s.ChannelMessageSend(channel, tmp)
		}
	}

	c := cron.New()
	c.AddFunc("45 6 * * *", func() {
		tmp := daily.getFrom7H()
		s.ChannelMessageSend(channel, tmp)
	})
	c.AddFunc("45 7 * * *", func() {
		tmp := daily.getFrom8H()
		if tmp != "empty" {
			s.ChannelMessageSend(channel, tmp)
		}
	})
	c.AddFunc("45 11 * * *", func() {
		tmp := daily.getFrom12H()
		s.ChannelMessageSend(channel, tmp)
	})
	c.AddFunc("45 17 * * *", func() {
		tmp := daily.getFrom18H()
		s.ChannelMessageSend(channel, tmp)
	})
	c.AddFunc("15 21 * * *", func() {
		tmp := daily.getFrom21H()
		if tmp != "empty" {
			s.ChannelMessageSend(channel, tmp)
		}
	})
	c.Start()
}
