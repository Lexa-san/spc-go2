package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Lexa-san/spc-go2/10.JWTAuth/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file (.toml or .env file)")
}

func main() {
	//Flag parsing and add value to configPath var
	flag.Parse()
	//config instance
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find path to config, app will use default confs:", err)
	}

	//server instance
	s := apiserver.New(config)

	//server start
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
