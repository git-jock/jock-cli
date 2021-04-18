package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type FolderConfig struct {
	Location string                       `mapstructure:"location"`
	Plugins  map[string]map[string]string `mapstructure:"plugins"`
}

// ReadConfig reads the config file from ~/.jockrc
func ReadConfig() {
	viper.SetConfigName(".jockrc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println(err)
		} else {
			// Config file was found but another error was produced
			fmt.Println(err)
		}
	}
}

const FOLDERS = "folders"

// GetFolderConfig checks that the config contains the specified folder, and returns marshalled FolderConfig if so.
func GetFolderConfig(name string) FolderConfig {
	folder := FolderConfig{}

	if !viper.InConfig(FOLDERS) {
		fmt.Print("Config does not contain folders")
		os.Exit(1)
	}

	if !viper.Sub(FOLDERS).InConfig(name) {
		fmt.Printf("Config folders do not contain [%s]", name)
		os.Exit(1)
	}

	err := viper.Sub(FOLDERS).UnmarshalKey(name, &folder)
	if err != nil {
		fmt.Printf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}

	return folder
}