package ship

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (s ShipCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	if args == "" {
		s.sendMessage(inputMsg.Chat.ID,
			"Пожалуйста, укажите идентификатор корабля. Пример: /get__logistic__ship 1")
		return
	}

	idx, err := strconv.Atoi(args)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, "Некорректный идентификатор. Введите целое число. Например: 1")
		return
	}

	ship, err := s.service.Describe(uint64(idx))
	if err != nil {
		log.Printf("Ошибка при получении корабля по идентификатору %d: %v", idx, err)
		s.sendMessage(inputMsg.Chat.ID, err.Error())
		return
	}

	s.sendMessage(inputMsg.Chat.ID, fmt.Sprintf("Корабль успешно найден: %s (id: %d)", ship.Title, ship.Id))
}
