package dataanalyzer

import (
	"database/sql"
	"log"
)

type Record struct {
	Data []byte
}

type DataGateway struct {
	DB *sql.DB
}

func (d *DataGateway) Find() ([]Record, error) {
	row := d.DB.QueryRow("select 'data' as data")
	var name string
	err := row.Scan(&name)
	log.Println("Found database records.")
	return []Record{{[]byte(name)}}, err
}

func (d *DataGateway) Save(_ []byte) error {
	row := d.DB.QueryRow("select 'data' as data")
	var name string
	_ = row.Scan(&name)
	log.Println("Updated database records.")
	return nil
}
