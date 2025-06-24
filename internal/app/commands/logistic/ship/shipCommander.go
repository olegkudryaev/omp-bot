package ship

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)
import "github.com/ozonmp/omp-bot/internal/service/logistic/ship"

type ShipCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type shipCommander struct {
	bot     *tgbotapi.BotAPI
	service *ship.ShipService
}

func NewShipCommander(bot *tgbotapi.BotAPI, service *ship.ShipService) ShipCommander {
	return &shipCommander{
		bot:     bot,
		service: service,
	}
}

func (c *shipCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *shipCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}

func (s shipCommander) Help(inputMsg *tgbotapi.Message) {
	//TODO implement me
	panic("implement me")
}

func (s shipCommander) Get(inputMsg *tgbotapi.Message) {
	//TODO implement me
	panic("implement me")
}

func (s shipCommander) List(inputMsg *tgbotapi.Message) {
	//TODO implement me
	panic("implement me")
}

func (s shipCommander) Delete(inputMsg *tgbotapi.Message) {
	//TODO implement me
	panic("implement me")
}

func (s shipCommander) New(inputMsg *tgbotapi.Message) {
	//TODO implement me
	panic("implement me")
}

func (s shipCommander) Edit(inputMsg *tgbotapi.Message) {
	//TODO implement me
	panic("implement me")
}

func (c *shipCommander) Default(msg *tgbotapi.Message) {
	//TODO implement me
	panic("implement me")
}

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *shipCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
