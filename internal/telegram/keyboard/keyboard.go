package keyboard

import "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func newReplyKeyboard(k [][]tgbotapi.KeyboardButton) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       k,
	}
}

func New(rows [][]string) tgbotapi.ReplyKeyboardMarkup {
	var result [][]tgbotapi.KeyboardButton
	for _, row := range rows {
		var keyboardButtons []tgbotapi.KeyboardButton
		for _, buttonText := range row {
			keyboardButtons = append(keyboardButtons, tgbotapi.NewKeyboardButton(buttonText))
		}
		result = append(result, keyboardButtons)
	}

	return newReplyKeyboard(result)
}

func Close() tgbotapi.ReplyKeyboardRemove {
	return tgbotapi.NewRemoveKeyboard(true)
}
