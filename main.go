package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	SLACK_APP_TOKEN := os.Getenv("SLACK_APP_TOKEN")
	bot := slacker.NewClient(SLACK_BOT_TOKEN, SLACK_APP_TOKEN)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "Year of Birth Calculator",
		Example:     "my yob is 2000",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			currentYear := time.Now().Year()
			age := currentYear - yob
			r := fmt.Sprintf("Age is %d", age)
			fmt.Println(r)
			fmt.Println("Hello?")
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
