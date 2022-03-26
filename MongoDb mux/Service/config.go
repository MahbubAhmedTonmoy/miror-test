package service

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connectionstring"`
	dbName           string `mapstructure:"dbname"`
	collcetionName   string `mapstructure:"collectionname"`
}

var AppConfig *Config

func LoadAppConfig() {
	fmt.Println("Loading Server Configurations...")
	viper.AddConfigPath("C:/Users/Admin/OneDrive/Desktop/GoWorkPlace/MongoDb mux")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
