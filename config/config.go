package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func GetDBConnection() string {
	host := viper.GetString(`dbns.database.host`)
	user := viper.GetString(`dbns.database.user`)
	password := viper.GetString(`dbns.database.password`)
	dbname := viper.GetString(`dbns.database.dbname`)
	port := viper.GetString(`dbns.database.port`)
	sslmode := viper.GetString(`dbns.database.sslmode`)
	firstStringConn := `host=` + host + ` user=` + user + ` password=` + password + ` dbname=` + dbname + ` port=` + port + ` sslmode=` + sslmode
	return firstStringConn
}

func Connect() *sql.DB {
	stringConnection := GetDBConnection()
	db, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
