package ship

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strconv"
	"strings"
)

func (s ShipCommander) List(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	first, rest, found := strings.Cut(args, " ")
	if !found {
		errorText := "Неверный формат команды. Ожидается два аргумента. Пример корректной команды:" +
			"/list__logistic__ship 0 3"
		s.sendMessage(inputMsg.Chat.ID, errorText)
		return
	}

	cursor, err := strconv.ParseUint(first, 10, 64)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID,
			fmt.Sprintf("Ошибка: 'cursor' должен быть числом, получено: %s", first))
		return
	}

	limit, err := strconv.ParseUint(rest, 10, 64)
	if err != nil {
		s.sendMessage(inputMsg.Chat.ID,
			fmt.Sprintf("Ошибка: 'limit' должен быть числом, получено: %s", rest))
		return
	}

	outputMsgText := "Корабли по запросу: \n"
	ships, err := s.service.List(cursor, limit)

	if err != nil {
		s.sendMessage(inputMsg.Chat.ID, err.Error())
		return
	}

	var builder strings.Builder
	for _, p := range ships {
		builder.WriteString(p.Title)
		builder.WriteString("\n")
	}
	outputMsgText = builder.String()

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: int(cursor + limit),
		Limit:  int(limit),
	})

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "ship",
		CallbackName: "list__logistic__ship",
		CallbackData: string(serializedData),
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = s.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}

}
