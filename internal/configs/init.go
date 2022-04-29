package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitializeViper() {
	// set files name of config file
	viper.SetConfigName("config.dev")

	// set path for config file
	// in this case "." is the root
	viper.AddConfigPath(".")

	// enables Viper to read env variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
