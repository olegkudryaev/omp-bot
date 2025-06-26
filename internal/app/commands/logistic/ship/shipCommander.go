package ship

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)
import "github.com/ozonmp/omp-bot/internal/service/logistic/ship"

type ShipCommanderIntergace interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type Service interface {
	Describe(shipID uint64) (*ship.Ship, error)
	List(cursor uint64, limit uint64) ([]ship.Ship, error)
	Create(s *ship.Ship) (uint64, error)
	Update(ShipID uint64, ship ship.Ship) error
	Remove(ShipID uint64) (bool, error)
}

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ShipCommander struct {
	bot     *tgbotapi.BotAPI
	service Service
}

func NewShipCommander(bot *tgbotapi.BotAPI) *ShipCommander {
	shipService := ship.NewShipService()
	return &ShipCommander{
		bot:     bot,
		service: shipService,
	}
}

func (s *ShipCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		s.CallbackList(callback, callbackPath)
	default:
		log.Printf("Ship.HandleCallback: неизвестный колбеэк: %s", callbackPath.CallbackName)
	}
}

func (s *ShipCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		s.Help(msg)
	case "list":
		s.List(msg)
	case "get":
		s.Get(msg)
	case "delete":
		s.Delete(msg)
	case "new":
		s.New(msg)
	case "edit":
		s.Edit(msg)
	default:
		s.Default(msg)
	}
}

func (s *ShipCommander) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := s.bot.Send(msg)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения в чат - %v", err)
	}
}
