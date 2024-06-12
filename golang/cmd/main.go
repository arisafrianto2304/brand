package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "db" // Sesuaikan dengan name Container Database
	port     = 5432 // Sesuaikan dengan port yang Anda gunakan, 5432 dari file konfigurasi Docker
	user     = "useraris"
	password = "passwordaris"
	dbname   = "dbaris"
)

func main() {
	// Membuat string koneksi
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Membuka koneksi ke database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Memeriksa koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	// Menjalankan kueri
	rows, err := db.Query("SELECT id, name, age, email FROM table1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		var email string

		err = rows.Scan(&id, &name, &age, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id = %d, name = %s, age = %d, email = %s\n", id, name, age, email)
	}

	// Memeriksa adanya kesalahan saat iterasi
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	port := "8080"
	log.Printf("server sedang berjalan pada http://localhost:%s/", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
