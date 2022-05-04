package commands

import (
	"github.com/Khalturin/telegram-bot-demo/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products := product.Service{}.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
