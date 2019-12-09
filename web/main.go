package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Listening on http://localhost:8088/")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to my API."))
	})
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World"))
	})
	http.HandleFunc("/connectDB", connectDBHandler)
	http.ListenAndServe(":80", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func connectDBHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Connect to DB\n"))
	const DRIVER = "mysql"
	const DSN = "root:admin@tcp(mysql:3306)/data_base"
	db, err := sql.Open(DRIVER, DSN)
	if err != nil {
		w.Write([]byte("Openエラーー\n"))
	} else {
		w.Write([]byte("OpenOK\n"))
	}
	err = db.Ping()
	if err != nil {
		w.Write([]byte("接続失敗\n"))
	} else {
		w.Write([]byte("接続OK\n"))
	}
	db.Close()
}
