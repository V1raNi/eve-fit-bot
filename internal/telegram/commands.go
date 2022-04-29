package telegram

import (
	"eve-fit-bot/internal/telegram/keyboard"
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (s *Service) handleCommand(m *telegramBotApi.Message) {
	msg := telegramBotApi.NewMessage(m.Chat.ID, m.Text)

	// Create a new MessageConfig. We don't have text yet,
	// so we leave it empty.
	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	numericKeyboard := keyboard.New([][]string{
		{"sky", "ocean"},
		{"red", "blue"},
	})

	numericKeyboard1 := keyboard.New([][]string{
		{"5", "12"},
		{"2", "5"},
	})
	// Extract the command from the Message.
	switch m.Command() {
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

	if _, err := s.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
