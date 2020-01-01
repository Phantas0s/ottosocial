package main

import (
	"fmt"

	"github.com/Phantas0s/tweetwee/cmd"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	cmd.Execute()
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
	}

	viper.AddConfigPath(home)
	viper.AddConfigPath("/home/hypnos")
	viper.SetConfigName("tuit")

	viper.AutomaticEnv()
	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
