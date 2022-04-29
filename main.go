package main

import (
	"errors"
	"eve-fit-bot/internal/api"
	"eve-fit-bot/internal/db"
	"eve-fit-bot/internal/telegram"
	"eve-fit-bot/internal/util"
	"log"
	"net/http"
	"sync"
)

func main() {
	log.Println("Loading Server Configurations...")

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	log.Printf("%+v\n", config)

	database := db.Init(config)
	_ = db.New(database)
	var wg sync.WaitGroup
	service, err := telegram.New(config.Token, true)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go service.HandleMessages()

	srv := api.NewServer("8088")

	if err := srv.Run(); errors.Is(err, http.ErrServerClosed) {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	wg.Wait()
}
