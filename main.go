package main

import (
	//"baymax/log"
	"flag"
	"fmt"
	"forum/api"
	"forum/config"
	"forum/model"
	"github.com/jinzhu/configor"
)

func init() {
	var configPath string

	flag.StringVar(&configPath, "c", "", "config file path")
	flag.Parse()

	if configPath != "" {
		if err := configor.Load(&config.Config, configPath); err != nil {
			fmt.Println(err)
			panic("can't file config file")
		}
	} else {
		configor.Load(&config.Config)
	}
	//
	//log.SetLogrus(config.Config.Logger.Level,
	//	config.Config.Logger.Format,
	//	config.Config.Logger.Out,
	//	config.Config.Logger.FluentdEnable,
	//	config.Config.Logger.FluentdHost,
	//	config.Config.Logger.FluentdPort,
	//	config.Config.Logger.FluentdTag)
}

func main() {
	model.InitDB(config.Config.Database.DSN, config.Config.Debug)

	model.MigrateDB()

	api.Run()

}
