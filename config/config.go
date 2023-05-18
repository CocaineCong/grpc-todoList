package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	fmt.Println(viper.GetString("datasource.host"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
