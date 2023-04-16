package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/maximys963/pizza-order/internal/app/apiserver"
	"github.com/maximys963/pizza-order/pkg/config"
	"github.com/maximys963/pizza-order/util"
	"log"
)

var (
	configPath string
)

func initDBConnection() {
	var dbConfig config.DBConfig

	dbConfig.Host = util.GetEnvOrFail("HOST")
	dbConfig.Port = util.GetEnvOrFail("PORT")
	dbConfig.User = util.GetEnvOrFail("DB_USER")
	dbConfig.Password = util.GetEnvOrFail("DB_PASS")
	dbConfig.DbName = "films_db"

	config.Connect(dbConfig)
}

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

	initDBConnection()

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
