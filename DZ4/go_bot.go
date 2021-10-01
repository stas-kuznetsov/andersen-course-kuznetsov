package main

import (
	"flag"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)

func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		reply := "Не знаю что сказать. Понимаю только команды /start, /Git, /Tasks и /Task#, где # - это номер задания"
		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот"

		case "Git":
			reply = "github.com/stas-kuznetsov/andersen-course-kuznetsov/"
		case "Tasks":
			reply = "DZ1, DZ2, DZ3, DZ4"
		case "Task1":
			reply = "github.com/stas-kuznetsov/andersen-course-kuznetsov/tree/main/DZ1"
		case "Task2":
			reply = "github.com/stas-kuznetsov/andersen-course-kuznetsov/tree/main/DZ2"
		case "Task3":
			reply = "github.com/stas-kuznetsov/andersen-course-kuznetsov/tree/main/DZ3"
		case "Task4":
			reply = "github.com/stas-kuznetsov/andersen-course-kuznetsov/tree/main/DZ4"


		}

		// создаем ответное сообщение
		ChatID := update.Message.Chat.ID 
		msg := tgbotapi.NewMessage(ChatID, reply)
		// отправляем
		bot.Send(msg)
	}
}


