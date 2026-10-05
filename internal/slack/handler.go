package slack

import (
	"fmt"
	"time"

	"github.com/shomali11/slacker"

	"github.com/StezX433/slack-age-bot/internal/age"
)

func RegisterCommands(bot *slacker.Slacker) {
	bot.Command("my yob is {year}", &slacker.CommandDefinition{
		Description: "Calculate age from year of birth",
		Examples:    []string{"my yob is 2004"},

		Handler: func(
			botCtx slacker.BotContext,
			request slacker.Request,
			response slacker.ResponseWriter,
		) {
			year := request.Param("year")

			dob, err := time.Parse("2006", year)
			if err != nil {
				response.Reply(
					"Invalid year. Please use a format like: my yob is 2004",
				)
				return
			}

			now := time.Now()

			currentAge := age.CalculateAge(dob, now)

			response.Reply(
				fmt.Sprintf("You are %d years old.", currentAge),
			)
		},
	})

	bot.Command("my dob is {date}", &slacker.CommandDefinition{
		Description: "Calculate detailed age from date of birth",
		Examples:    []string{"my dob is 15-07-2004"},

		Handler: func(
			botCtx slacker.BotContext,
			request slacker.Request,
			response slacker.ResponseWriter,
		) {
			date := request.Param("date")

			dob, err := parseDOB(date)
			if err != nil {
				response.Reply(
					"Invalid date. Please use DD-MM-YYYY, for example: my dob is 15-07-2004",
				)
				return
			}

			now := time.Now()

			currentAge := age.CalculateAge(dob, now)
			daysAlive := age.DaysAlive(dob, now)
			nextBirthday := age.NextBirthday(dob, now)
			nextMilestone := age.NextMilestone(currentAge)

			message := fmt.Sprintf(
				"🎂 Age: %d years\n"+
					"🌍 Days alive: %d\n"+
					"🎉 Next birthday: %s\n"+
					"🎯 Next milestone: %d years",
				currentAge,
				daysAlive,
				nextBirthday.Format("02 January 2006"),
				nextMilestone,
			)

			response.Reply(message)
		},
	})
}

func parseDOB(value string) (time.Time, error) {
	return time.Parse("02-01-2006", value)
}
