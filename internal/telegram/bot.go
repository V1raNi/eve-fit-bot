package telegram

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Service struct {
	bot *telegramBotApi.BotAPI
}

func New(token string, debug bool) (*Service, error) {
	bot, err := telegramBotApi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s. Bot debug is: %t", bot.Self.UserName, debug)

	return &Service{
		bot: bot,
	}, nil
}

func (s *Service) HandleMessages() {
	u := telegramBotApi.NewUpdate(0)
	u.Timeout = 30

	updates := s.bot.GetUpdatesChan(u)

	for update := range updates {
		message := update.Message
		if message == nil { // ignore any non-Message updates
			continue
		}

		if message.IsCommand() {
			s.handleCommand(message)
		} else {
			s.handleMessage(message)
		}

		if s.bot.Debug {
			log.Printf("Get message '%s' from user '%s' from chat id %d\n",
				message.Text, message.From.UserName, message.Chat.ID,
			)
		}
	}
}
