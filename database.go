package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB{
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/golang_database")	
	if err != nil{
		panic(err)
	}

	//berapa jumlah koneksi minimal yang dibuat
	db.SetMaxIdleConns(10)
	//berapa jumlah koneksi maksimal yang dibuat
	db.SetMaxOpenConns(100)
	// berapa lama koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxIdleTime(5*time.Minute)
	// berapa lama koneksi boleh digunakan
	db.SetConnMaxLifetime(60*time.Minute)

	return db
}