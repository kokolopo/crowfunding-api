package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

//function yang menyediakan configurasi DB
func GetConfig() Configuration {
	conf := Configuration{}

	//GetConf -> mengambil seluruh env.json dan menaruhnya pada interface
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
