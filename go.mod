module onmyoji

go 1.15

require (
	github.com/TraiOi/notify v0.0.0
	github.com/TraiOi/shikigami v0.0.0
	github.com/TraiOi/summoning v0.0.0
	github.com/TraiOi/translator v0.0.0
	github.com/TraiOi/utils v0.0.0
	github.com/TraiOi/validator v0.0.0
	github.com/bwmarrin/discordgo v0.22.0
	github.com/ozankasikci/go-image-merge v0.2.2 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/viper v1.7.1
)

replace github.com/TraiOi/notify => ./notify

replace github.com/TraiOi/utils => ./utils

replace github.com/TraiOi/translator => ./translator

replace github.com/TraiOi/validator => ./validator

replace github.com/TraiOi/shikigami => ./shikigami

replace github.com/TraiOi/summoning => ./summoning
