package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

//LineApp :
type LineApp struct {
	bot *linebot.Client
}

//NewLineApp : New LineApp
func NewLineApp(channelSecret, channelToken string) (*LineApp, error) {
	bot, err := linebot.New(
		channelSecret,
		channelToken,
	)
	if err != nil {
		return nil, err
	}
	return &LineApp{
		bot: bot,
	}, nil
}

func main() {
	log.Println("Channel Secret : " + os.Getenv("CHANNEL_SECRET"))
	log.Println("Channel Token : " + os.Getenv("CHANNEL_TOKEN"))

	app, err := NewLineApp(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/linebot", app.callbackHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	log.Println("Server is Running...")
}

func (app *LineApp) callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
			log.Println("Invalid Signature")
			log.Println("X-Line-Signature: " + r.Header.Get("X-Line-Signature"))
		} else {
			w.WriteHeader(500)
			log.Println("Unknow error")
		}
		return
	}

	log.Printf("Got events %v", events)
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = app.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func (app *LineApp) replyText(replyToken, text string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}
