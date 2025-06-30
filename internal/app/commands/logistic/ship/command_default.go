package ship

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *ShipCommander) Default(msg *tgbotapi.Message) {
	s.sendMessage(msg.Chat.ID, fmt.Sprintf("Не корректная команда. Попробуйте вызвать команду help, "+
		"для получения списка комманд: /help__logistic__ship"))

}
