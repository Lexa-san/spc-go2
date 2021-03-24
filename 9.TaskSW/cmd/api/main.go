package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Lexa-san/spc-go2/9.TaskSW/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	//В этот момент происходит инициализация переменной configPath значением
	flag.Parse()
	log.Println("It works")
	//server instance initialization
	config := api.NewConfig()
	// Десериалзиуете содержимое .toml файла
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	server := api.New(config)

	//api server start
	log.Fatal(server.Start())
}
