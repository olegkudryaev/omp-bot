package ship

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strings"
)

type CallbackListData struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (s *ShipCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	raw := callbackPath.CallbackData
	jsonStart := strings.Index(raw, "{")
	if jsonStart == -1 {
		s.sendMessage(callback.Message.Chat.ID, "Техническая ошибка при выводе следующей страницы")
		log.Printf("Некорректный формат callbackData: %s", raw)
		return
	}
	jsonStr := raw[jsonStart:]

	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(jsonStr), &parsedData)
	if err != nil {
		s.sendMessage(callback.Message.Chat.ID, "Техническая ошибка при выводе следующей страницы")
		log.Printf("Ship.CallbackList: Ошибка при чтении CallbackListData: %v - %v",
			callbackPath.CallbackData, err)
		return
	}

	cursor := uint64(parsedData.Offset)
	limit := uint64(parsedData.Limit)

	ships, err := s.service.List(cursor, limit)

	if err != nil {
		s.sendMessage(callback.Message.Chat.ID, err.Error())
		return
	}

	outputMsgText := "Корабли по запросу: \n"
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

	newCallbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "ship",
		CallbackName: "list__logistic__ship",
		CallbackData: string(serializedData),
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		outputMsgText,
	)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", newCallbackPath.String()),
		),
	)

	_, err = s.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
