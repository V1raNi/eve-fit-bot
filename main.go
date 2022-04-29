package main

import (
	"eve-fit-bot/pkg/db"
	"eve-fit-bot/pkg/util"
	"eve-fit-bot/pkg/util/keyboard"
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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

	bot, err := telegramBotApi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegramBotApi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		//if !update.Message.IsCommand() { // ignore any non-command Messages
		//	continue
		//}
		msg := telegramBotApi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		numericKeyboard := keyboard.New([][]string{
			{"sky", "ocean"},
			{"red", "blue"},
		})

		numericKeyboard1 := keyboard.New([][]string{
			{"5", "12"},
			{"2", "5"},
		})
		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /open, /close, and /status."
		case "open":
			msg.Text = "Opening keyboard"
			msg.ReplyMarkup = numericKeyboard
		case "open1":
			msg.Text = "Opening keyboard1"
			msg.ReplyMarkup = numericKeyboard1
		case "close":
			msg.Text = "Closing keyboard"
			msg.ReplyMarkup = keyboard.Close()
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
