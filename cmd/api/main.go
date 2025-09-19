package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TillBurdorf/Kinoprojekt/internal/database"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

const version = "1.0.0"

type config struct {
	port  int
	env   string
	dbdsn string
}

type application struct {
	config  config
	logger  *log.Logger
	db      *gorm.DB
	storage database.Storage
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.dbdsn, "dbdsn", "postgres://postgres:12345@localhost:5432/kinoprojekt_db?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.Connect(cfg.dbdsn)
	if err != nil {
		logger.Fatal(err)
	}

	app := &application{
		config: cfg,
		logger: logger,
		db:     db,
	}

	addr := fmt.Sprintf(":%d", cfg.port)
	app.storage = database.NewStorage(db)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
