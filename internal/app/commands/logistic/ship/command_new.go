package ship

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/logistic/ship"
	"log"
)

func (s ShipCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	var shipInput ship.Ship
	err := json.Unmarshal([]byte(args), &shipInput)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, "Ошибка: не удалось разобрать JSON. "+
			"Правильный формат: {\"id\":2,\"title\":\"Aurora\"}")
		log.Printf("Ошибка: не удалось разобрать JSON: " + err.Error())
		return
	}

	_, err = s.service.Create(&shipInput)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, err.Error())
		log.Printf(err.Error())
		return
	}

	s.sendMessage(inputMsg.Chat.ID, fmt.Sprintf("Корабль успешно создан с id: %d", shipInput.Id))
}
