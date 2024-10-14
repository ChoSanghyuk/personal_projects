package bot

import (
	"encoding/json"
	"io"
	"net/http"

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

func (t TeleBot) Listen(ch chan string) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			txt := update.Message.Text
			if txt[0] != '/' {
				continue
			}

			switch txt {
			case "/help":
				ch <- `
				조회 API 목록
				/funds
				/funds/{id}/hist
				/funds/:{id}/assets
				/assets/list
				/assets/{id}
				/assets/{id}/hist
				/market
				/market/indicators/{date?}
				`
			default:
				rtn, err := httpsend(txt)
				if err != nil {
					ch <- err.Error()
				} else {
					ch <- rtn
				}
			}

		}
	}
}

func httpsend(path string) (string, error) {

	url := "http://localhost:3000" + path
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var jsonData interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return "", err
	}

	pretty, err := json.MarshalIndent(jsonData, "", "\t")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}
