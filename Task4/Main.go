package main

import (
	"log"
	"os"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		os.Exit(9)
	}
}

func main() {
	// os.Setenv("TOKEN", "1963194779:AAENIjqo8Z-nuZJupalKglhvJEy8DIjQMPk")
	//Создаем бота
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	} else {
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/Git":
			msg := prepareRepoLink()
			bot.Send(setMsgOverlay(msg, update))
		case "/Tasks":
			bot.Send(setMsgOverlay("Sending task status request to Raven", update))
			parseTasks()
			//checkTaskValidity("dogs")

		}
	}
}

func setMsgOverlay(content string, update tgbotapi.Update) tgbotapi.Chattable {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, content)
	message.ReplyToMessageID = update.Message.MessageID
	return message
}
