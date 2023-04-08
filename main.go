package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/maximys963/pizza-order/internal/app/apiserver"
	"github.com/maximys963/pizza-order/util"
	"gorm.io/gorm"
	"log"
)

var (
	configPath string
)

func onceInitDBOnboardingConnection() *gorm.DB {
	var (
		host     = util.GetEnvOrFail("HOST")
		port     = util.GetEnvOrFail("PORT")
		user     = util.GetEnvOrFail("DB_USER")
		password = util.GetEnvOrFail("DB_PASS")
		dbname   = "films_db"
	)
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	connection, err := util.CreateDBConnection(
		dbUrl,
	)
	if err != nil {
		log.Panic(err)
	}

	return connection
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

	onceInitDBOnboardingConnection()

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
