package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
		log.Fatal("DATABASE_URL tidak ditemukan")
	}

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database tidak dapat diakses:", err)
	}

	log.Println("Berhasil terhubung ke database")
}
