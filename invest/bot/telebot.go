package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TeleBot struct {
	bot    *tgbotapi.BotAPI
	chatId int64
}

func NewTeleBot(token string, chatId int64) (*TeleBot, error) {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true

	return &TeleBot{
		bot:    bot,
		chatId: chatId,
	}, nil
}

func (t TeleBot) SendMessage(msg string) {
	t.bot.Send(tgbotapi.NewMessage(t.chatId, msg))
}
