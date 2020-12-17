package translator

import (
	"github.com/spf13/viper"
)

func translate(filename string) *viper.Viper {
	vp := viper.New()
	vp.SetConfigName(filename + ".yaml")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("data/dics/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return vp
}
