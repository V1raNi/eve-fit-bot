package main

import (
	"eve-fit-bot/internal/db"
	"eve-fit-bot/internal/telegram"
	"eve-fit-bot/internal/util"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	log.Println("Loading Server Configurations...")

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	log.Printf("%+v\n", config)

	_ = db.Init(config)
	service, err := telegram.New(config.Token, true)
	if err != nil {
		log.Fatal(err)
	}
	go service.HandleMessages()

	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.Logger.Fatal(app.Start(":8080"))
}
