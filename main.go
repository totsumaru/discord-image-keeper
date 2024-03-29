package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/techstart35/discord-image-keeper/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	godotenv.Load(".env")

	location := os.Getenv("TZ")
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	var Token = "Bot " + os.Getenv("APP_BOT_TOKEN")

	session, err := discordgo.New(Token)
	session.Token = Token
	if err != nil {
		log.Fatalln(err)
	}

	session.AddHandler(handler.Basic)
	session.AddHandler(handler.PanelCreate)
	session.AddHandler(handler.SecretBasic)

	if err = session.Open(); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = session.Close(); err != nil {
			log.Fatalln(err)
		}
		return
	}()

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
}
