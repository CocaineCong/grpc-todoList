package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
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