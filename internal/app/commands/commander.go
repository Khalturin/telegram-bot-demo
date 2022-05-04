package commands

import (
	"github.com/Khalturin/telegram-bot-demo/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil { // ignore any non-Message Updates
		return
	}

	command, ok := registeredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}

	//switch update.Message.Command() {
	//case "help":
	//	c.Help(update.Message)
	//case "list":
	//	c.List(update.Message)
	//default:
	//	c.Default(update.Message)
	//}
}
