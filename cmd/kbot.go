/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var (
	//teletoken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// choise for bot
var choices = map[int]string{
	0: "rock",
	1: "paper",
	2: "scissors",
}

var (
	// player answer
	playerAnswer = ""
	// bot answer
	botAnswer   = ""
	playerScore = 0
	botScore    = 0
)

// menu builder
var menu = &telebot.ReplyMarkup{ResizeKeyboard: true}

// buttons
var (
	btnRock     = menu.Text("rock")
	btnPaper    = menu.Text("paper")
	btnScissors = menu.Text("scissors")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started ", appVersion)

		menu.Reply(
			menu.Row(btnRock, btnPaper, btnScissors),
		)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please, check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(message telebot.Context) error {
			log.Print(message.Message().Payload, message.Text())
			payload := message.Message().Payload

			switch payload {
			case "play":
				err = message.Send("Here is a menu", menu)
			case btnRock.Text, btnPaper.Text, btnScissors.Text:
				playerAnswer := payload
				botAnswer := getBotAnswer()
				playerScore, botScore = calculateScore(playerScore, botScore, playerAnswer, botAnswer)
				err = message.Send(fmt.Sprintf("I got %s.\nOur Score: You %d, Bot %d", botAnswer, playerScore, botScore))
			case "hello":
				err = message.Send(fmt.Sprintf("Hello, I am kbot %s", appVersion))
			default:
				err = message.Send(fmt.Sprintf("Your messge recieved %s!", appVersion))
			}

			return err
		})
		kbot.Start()
	},
}

func getBotAnswer() string {
	randomAnswer := rand.Intn(len(choices))

	return choices[randomAnswer]
}

func calculateScore(playerScore int, botScore int, playerAnswer string, botAnswer string) (int, int) {
	if playerAnswer == botAnswer {
		fmt.Print("No one")
	} else if playerAnswer == "paper" && botAnswer == "rock" {
		playerScore += 1
	} else if playerAnswer == "scissors" && botAnswer == "rock" {
		botScore += 1
	} else if playerAnswer == "rock" && botAnswer == "paper" {
		botScore += 1
	} else if playerAnswer == "rock" && botAnswer == "scissors" {
		playerScore += 1
	} else if playerAnswer == "scissors" && botAnswer == "paper" {
		playerScore += 1
	} else if playerAnswer == "paper" && botAnswer == "scissors" {
		botScore += 1
	}
	return playerScore, botScore
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
