package providers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alfuhigi/news-ajf-sa/config"
	_ "github.com/lib/pq"
)

// Connect open conntion with postgresql
func Connect() *sql.DB {
	psqlinfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Config("DBUSER"),
		config.Config("DBPASS"),
		config.Config("DBHOST"),
		config.Config("DBPORT"),
		config.Config("DBNAME"))
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Println("DB Open: ", err)
	}
	if err = db.Ping(); err != nil {
		log.Println("Ping: ", err)
	}
	return db
}
