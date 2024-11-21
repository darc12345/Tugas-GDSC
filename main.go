package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/controller"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	file, err := os.OpenFile("log/BACKENDLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	cfg := mysql.Config{
		User:   "bob",
		Passwd: os.Getenv("PASSWRD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "gdsc_bookdb",
	}
	fmt.Println(os.Getenv("PASSWRD"))
	// postgresURI := "mysql://root:password@(127.0.0.1:3306)/projectdb"
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Print("Failed to connect ot the database")
		log.Print(err)
		os.Exit(0)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Print("ping db gagal")
		log.Print(err)
		os.Exit(0)
	}

	r := gin.Default()
	server := &http.Server{
		Addr:              ":8000",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           r,
	}
	ctrlHandler := controller.NewControllerDB(db)

	r.GET("/api/books", ctrlHandler.GetBooksHandler)
	r.POST("/api/books", ctrlHandler.PostBooksHandler)
	r.PUT("/api/books/:id", ctrlHandler.PutBooksHandler)
	r.DELETE("/api/books/:id", ctrlHandler.DeleteBooksHandler)
	r.GET("/api/books/:id", ctrlHandler.GetBookByIDHandler)
	log.Fatal(server.ListenAndServe())
}
