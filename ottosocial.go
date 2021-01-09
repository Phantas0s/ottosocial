package main

import (
	"fmt"

	"github.com/Phantas0s/ottosocial/cmd"
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
	viper.AddConfigPath(".")
	viper.SetConfigName("ottosocial")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	fmt.Println("ottosocial running...")
}
