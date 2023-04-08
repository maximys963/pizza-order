package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/maximys963/pizza-order/util"
	"log"
)

func Init() {
	var (
		host     = util.GetEnvOrFail("HOST")
		port     = util.GetEnvOrFail("PORT")
		user     = util.GetEnvOrFail("DB_USER")
		password = util.GetEnvOrFail("DB_PASS")
		dbname   = "films_db"
	)

	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(user)
	fmt.Println(password)
	fmt.Println(dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Successfully connected!")
}
