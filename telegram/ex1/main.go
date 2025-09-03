package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken := "8139701471:AAGEKavSrGEkCgIyxJmeZjp6KX7wZBo8dKA" // Замените на ваш API токен
	chatID := int64(-1002768113061)                              // Замените на ID чата
	messageText := "Привет из Go!"                               // Текст сообщения

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // Включите отладку для вывода информации в консоль

	log.Printf("Авторизован как %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(chatID, messageText)
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
