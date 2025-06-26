package ship

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/logistic/ship"
	"strconv"
	"strings"
)

func (s ShipCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	firstSpace := strings.Index(args, " ")
	if firstSpace == -1 {
		s.sendMessage(inputMsg.Chat.ID, fmt.Sprintf("Ошибка: ожидается два аргумента: идентификатор и JSON с "+
			"моделью Ship. Пример корректного запроса: 1 {Id:1 Title:Aurora}"))
		return
	}

	idStr := args[:firstSpace]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, "Ошибка: идентификатор должен быть числом")
		fmt.Printf("Ошибка: идентификатор должен быть числом: %v\n", err)
		return
	}

	jsonStr := strings.TrimSpace(args[firstSpace+1:])
	var shipInput ship.Ship
	err = json.Unmarshal([]byte(jsonStr), &shipInput)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, "Ошибка: передан не корректный Json. "+
			"Пример корректного Json: {Id:1 Title:Aurora}")
		fmt.Printf("Ошибка: не удалось разобрать JSON: " + err.Error())
		return
	}

	err = s.service.Update(id, shipInput)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, err.Error())
		return
	}
	s.sendMessage(inputMsg.Chat.ID, fmt.Sprintf("Корабль обновлен. Текущее значение полей: "+
		"Ship: %+v", shipInput))
}
