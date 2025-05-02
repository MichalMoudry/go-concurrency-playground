package main

import (
	"fmt"
	"log"
	"net/http"
	"section-04/internal/config"
	"section-04/internal/database"
	"section-04/internal/transport"
)

func main() {
	// set up the app config
	cfg, err := config.LoadConfigFromFile("config.toml")
	if err != nil {
		log.Fatalln(err)
	}

	// connect to DB
	_, err = database.OpenDb(cfg.DbConnStr)
	if err != nil {
		log.Fatalln(err)
	}

	// create sessions
	redis := transport.NewRedisPool(cfg.RedisConnStr)
	sessionMgr := transport.NewSessionManager(redis)

	// create some channels

	// waitgroup

	// set up mail

	// listen for web connections
	handler := transport.NewHandler(cfg.Port)
	log.Printf("Launching server on port: %d\n", handler.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", handler.Port), sessionMgr.LoadAndSave(handler.Mux)); err != nil {
		log.Fatalln(err)
	}
}
