package ship

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (s ShipCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__logistic__ship - выводит список всех команд\n"+
			"/get__logistic__ship — get a entity\n"+
			"/list__logistic__ship — get a list of your entity\n"+
			"/delete__logistic__ship — delete an existing entity\n"+
			"/new__logistic__ship — create a new entity\n"+
			"/edit__logistic__ship — edit a entity\n",
	)

	_, err := s.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
