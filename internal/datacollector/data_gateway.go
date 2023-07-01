package datacollector

import (
	"database/sql"
	"log"
)

type Record struct {
	Url string
}

type DataGateway struct {
	DB *sql.DB
}

func (d *DataGateway) Find() ([]Record, error) {
	row := d.DB.QueryRow("select 'https://feed.infoq.com/' as data")
	var url string
	err := row.Scan(&url)
	log.Println("Found database records.")
	return []Record{{url}}, err
}

func (d *DataGateway) Save(_ []byte) error {
	row := d.DB.QueryRow("select 'https://feed.infoq.com/' as data")
	var url string
	_ = row.Scan(&url)
	log.Println("Updated database records.")
	return nil
}
