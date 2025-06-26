package ship

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (s ShipCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, "Некорректный идентификатор. Введите целое число. Например: 1")
		return
	}

	_, err = s.service.Remove(uint64(idx))
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, err.Error())
		return
	}
	s.sendMessage(inputMsg.Chat.ID, fmt.Sprintf("Корабль с id %d успешно удалён.", idx))
}
