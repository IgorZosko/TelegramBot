package main

import (
	"flag"
	"log"
	"TelegramBot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main()  {
	tgClient := telegram.New(tgBotHost,mustToken())
}

func mustToken () string{
	token := flag.String("token-bot-token", "", "token for access to telegram bot ")
	flag.Parse()

	if *token == ""{
		log.Fatal("token in not specified")
	}
	return *token
}