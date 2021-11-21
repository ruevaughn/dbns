package repositories

import (
	"database/sql"
	"fmt"
	"log"
)

type PsqlNucleiRepository struct {
	DB    *sql.DB
	Table string
}

func (n PsqlNucleiRepository) GetSubdomains() error {
	records, err := n.DB.Query(`select host from "nucleires"`)
	if err != nil {
		log.Fatal("ERR", err)
	}
	for records.Next() {
		row := ""
		err = records.Scan(&row)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(row)
	}
	return nil
}

func (n PsqlNucleiRepository) AddSubdomain() error {
	fmt.Println("Add to db repository")
	return nil
}
