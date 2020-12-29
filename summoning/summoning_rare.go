package summoning

import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"image/png"
	"image/jpeg"
	dgo "github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	gim "github.com/ozankasikci/go-image-merge"
)

type Shikigami struct {
	Type string
	Name string
	Image string
}

type RareSummoning struct {
	R   []Shikigami
	SR  []Shikigami
	SSR []Shikigami
	SP  []Shikigami
}

func HandlerRare(s *dgo.Session, m *dgo.MessageCreate) {
	vp := viper.New()
	vp.SetConfigName("discord.yaml")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("config/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	channel := vp.GetString("ChannelSummoning")

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.ChannelID != channel {
		return
	}
	userID := m.Author.ID
	userName := m.Author.Username

	rsum := new(RareSummoning)
	listShiki := rsum.getListShikigami()
	tmp := rsum.sum10(listShiki)
	rsum.mergeShikiImg(tmp, userID)

	fileName := "./tmp/rsum_" + userID + ".png"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ms := &dgo.MessageSend{
		Files: []*dgo.File{
			 &dgo.File{
		            Name:   fileName,
			    Reader: f,
			},
		},
		Content: "Chúc mừng **" + userName + "** đã triệu hồi được các thức thần:",
	}
	if m.Content == "/test" {
		s.ChannelMessageSendComplex("788343042461401118", ms)
	}
	if m.Content == "/rsum" {
		s.ChannelMessageSendComplex(channel, ms)
	}
}

func (this *RareSummoning) getListShikigami() []Shikigami {
	vp := viper.New()
	vp.SetConfigName("shikigami.yaml")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("data/dics/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	rShiki   := vp.Sub("R").AllSettings()
	srShiki  := vp.Sub("SR").AllSettings()
	ssrShiki := vp.Sub("SSR").AllSettings()
	spShiki  := vp.Sub("SP").AllSettings()

	this.getRShikigami(rShiki)
	this.getSRShikigami(srShiki)
	this.getSSRShikigami(ssrShiki)
	this.getSPShikigami(spShiki)

	listShiki := this.mergeList(this.R, this.SR, this.SSR, this.SP)

	return listShiki
}

func (this *RareSummoning) getRShikigami(content map[string]interface{}) {
	var shiki []Shikigami
	for name, data := range content {
		tmp1 := data.(map[string]interface{})
		tmp2 := Shikigami {
			Type: "R",
			Name: tmp1["vsub"].(string),
			Image: "r_" + name + ".png",
		}
		shiki = append(shiki, tmp2)
	}
	this.R = shiki
}

func (this *RareSummoning) getSRShikigami(content map[string]interface{}) {
	var shiki []Shikigami
	for name, data := range content {
		tmp1 := data.(map[string]interface{})
		tmp2 := Shikigami {
			Type: "SR",
			Name: tmp1["vsub"].(string),
			Image: "sr_" + name + ".png",
		}
		shiki = append(shiki, tmp2)
	}
	this.SR = shiki
}

func (this *RareSummoning) getSSRShikigami(content map[string]interface{}) {
	var shiki []Shikigami
	for name, data := range content {
		tmp1 := data.(map[string]interface{})
		tmp2 := Shikigami {
			Type: "SSR",
			Name: tmp1["vsub"].(string),
			Image: "ssr_" + name + ".png",
		}
		shiki = append(shiki, tmp2)
	}
	this.SSR = shiki
}

func (this *RareSummoning) getSPShikigami(content map[string]interface{}) {
	var shiki []Shikigami
	for name, data := range content {
		tmp1 := data.(map[string]interface{})
		tmp2 := Shikigami {
			Type: "SP",
			Name: tmp1["vsub"].(string),
			Image: "sp_" + name + ".png",
		}
		shiki = append(shiki, tmp2)
	}
	this.SP = shiki
}

func (this *RareSummoning) mergeList(r   []Shikigami,
				     sr  []Shikigami,
				     ssr []Shikigami,
				     sp  []Shikigami) []Shikigami {
	var result []Shikigami
	for i := 0; i < 80; i++ {
		result = append(result, r...)
	}
	for i := 0; i < 20; i++ {
		result = append(result, sr...)
	}
	for i := 0; i < 2; i++ {
		result = append(result, ssr...)
	}
	result = append(result, sp...)

	return result
}

func (this *RareSummoning) sum10(content []Shikigami) []Shikigami {
	var sumList []Shikigami
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(content), func(i, j int) {
		content[i], content[j] = content[j], content[i]
	})

	for i := 0; i < 10; i++ {
		tmp := content[rand.Intn(len(content))]
		sumList = append(sumList, tmp)
	}
	return sumList
}

func (this *RareSummoning) getShikiInfo(content []Shikigami) string {
	var result string
	for _, shiki := range content {
		result += fmt.Sprintf("[%s] %s\n", shiki.Type, shiki.Name)
	}

	return result
}

func (this *RareSummoning) mergeShikiImg(content []Shikigami,
					 userID string) {
	grids := []*gim.Grid{
		{ImageFilePath: "data/img/shikigami/" + content[0].Image},
		{ImageFilePath: "data/img/shikigami/" + content[1].Image},
		{ImageFilePath: "data/img/shikigami/" + content[2].Image},
		{ImageFilePath: "data/img/shikigami/" + content[3].Image},
		{ImageFilePath: "data/img/shikigami/" + content[4].Image},
		{ImageFilePath: "data/img/shikigami/" + content[5].Image},
		{ImageFilePath: "data/img/shikigami/" + content[6].Image},
		{ImageFilePath: "data/img/shikigami/" + content[7].Image},
		{ImageFilePath: "data/img/shikigami/" + content[8].Image},
		{ImageFilePath: "data/img/shikigami/" + content[9].Image},
	}
	rgba, err := gim.New(grids, 5, 2).Merge()
	if err != nil {
		panic(err)
	}
	file, err := os.Create("tmp/rsum_" + userID + ".png")
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		panic(err)
	}
	err = png.Encode(file, rgba)
	if err != nil {
		panic(err)
	}
}
