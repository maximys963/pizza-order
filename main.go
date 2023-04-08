package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/maximys963/pizza-order/internal/app/apiserver"
	"github.com/maximys963/pizza-order/pkg/client/postgresql"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs/apiserver.toml", "path to config file")
}

func initEnv() {
	err := godotenv.Load(".env")

	fmt.Println("here")

	if err != nil {
		log.Println(err)
	}
}

func main() {
	initEnv()
	flag.Parse()

	postgresql.Init()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
