package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

var rootCmd = &cobra.Command{
	Use:   "kbot",
	Short: "Simple Telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		token := os.Getenv("TELE_TOKEN")
		pref := telebot.Settings{
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}

		bot, err := telebot.NewBot(pref)
		if err != nil {
			log.Fatal(err)
		}

		bot.Handle(telebot.OnText, func(c telebot.Context) error {
			return c.Send("Привіт! Ви написали: " + c.Text())
		})

		log.Println("Бот запущено...")
		bot.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
