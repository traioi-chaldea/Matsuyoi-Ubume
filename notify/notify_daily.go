package notify

import (
	"bytes"
	"fmt"
	"text/template"
	"reflect"
	vp "github.com/spf13/viper"

	"github.com/TraiOi/utils"
	"github.com/TraiOi/translator"
)

type DailyFunc struct {
	Config *vp.Viper
}

func (this *DailyFunc) Init(wDay string) {
	this.Config = vp.New()
	this.Config.SetConfigName(wDay + ".yaml")
	this.Config.SetConfigType("yaml")
	this.Config.AddConfigPath("data/daily/")

	err := this.Config.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func parseFrom7H(content Daily7h) Daily7h {
	content = Daily7h {
		content.NetherworldGate,
		translator.Soul(content.DemonEncounter),
	}
	return content
}

func (this *DailyFunc) getFrom7H() string {
	var content Daily7h
	var hTime = "7"
	info := this.Config.Sub(hTime)

	content.NetherworldGate = info.GetBool("netherworld_gate")
	content.DemonEncounter = info.GetString("demon_encounter")

	content = parseFrom7H(content)
	return this.getTemplate(hTime, content)
}

func parseFrom8H(content Daily8h) Daily8h {
	content = Daily8h{
		translator.Kirin(content.GuildRaid),
	}
	return content
}

func (this *DailyFunc) getFrom8H() string {
	var content Daily8h
	var tpl string
	var hTime = "8"
	info := this.Config.Sub(hTime)

	content.GuildRaid = info.GetString("guild_raid")

	if reflect.ValueOf(content).IsZero() {
		tpl = "empty"
	} else {
		content = parseFrom8H(content)
		tpl = this.getTemplate(hTime, content)
	}

	return tpl
}

func parseFrom12H(content Daily12h) Daily12h {
	for index, _ := range content.SoulZone10 {
		tmp := content.SoulZone10[index]
		content.SoulZone10[index] = translator.Soul(tmp)
	}
	content.SoulZone11 = translator.Soul(content.SoulZone11)
	return content
}

func (this *DailyFunc) getFrom12H() string {
	var content Daily12h
	var hTime = "12"
	info := this.Config.Sub(hTime)

	content.Maintenance = info.GetBool("maintenance")
	content.SoulZone10 = info.GetStringSlice("soul_zone_10")
	content.SoulZone11 = info.GetString("soul_zone_11")
	content.TotemZone = info.GetString("totem_zone")

	content = parseFrom12H(content)
	return this.getTemplate(hTime, content)
}

func parseFrom18H(content Daily18h) Daily18h {
	content.DemonEncounter = translator.Soul(content.DemonEncounter)
	return content
}

func (this *DailyFunc) getFrom18H() string {
	var content Daily18h
	var hTime = "18"
	info := this.Config.Sub(hTime)

	content.DemonEncounter = info.GetString("demon_encounter")
	content.NetherworldGate = info.GetBool("netherworld_gate")

	content = parseFrom18H(content)
	return this.getTemplate(hTime, content)
}

func (this *DailyFunc) getFrom21H() string {
	var content Daily21h
	var tpl string
	var hTime = "21"
	info := this.Config.Sub(hTime)

	content.GuildFeast = info.GetBool("guild_feast")
	if reflect.ValueOf(content).IsZero() {
		tpl = "empty"
	} else {
		tpl = this.getTemplate(hTime, content)
	}

	return tpl
}

func (this *DailyFunc) getTemplate(hTime string,
				       content interface{}) string {
	var td DailyEntry
	var result bytes.Buffer
	tplName := fmt.Sprintf("daily%sh.tpl", hTime)

	td.Greeting = utils.GetGreeting()
	td.Time = hTime
	td.Content = content

	t := template.Must(template.New(tplName).ParseFiles("data/template/" + tplName))
	err := t.Execute(&result, td)
	if err != nil {
		panic(err)
	}

	return result.String()
}
