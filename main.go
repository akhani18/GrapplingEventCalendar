package main

import (
	"fmt"
	"time"

	"github.com/akhani18/GrapplingEventCalendar/alexa"
	"github.com/akhani18/GrapplingEventCalendar/dao"
	"github.com/aws/aws-lambda-go/lambda"
)

// Upcoming events logic goes here.
func HandleUpcomingEventsIntent(request alexa.Request) alexa.Response {
	state := request.Body.Intent.Slots["state"].Value
	comps := dao.GetUpcomingCompetitionsInState(state)

	var builder alexa.SSMLBuilder

	if len(comps) == 0 {
		builder.Say(fmt.Sprintf("I couldn't find any upcoming competitions in the state of %s.", state))
	} else {
		builder.Say(fmt.Sprintf("I found %d upcoming competitions in %s.", len(comps), state))
		builder.Pause("500")

		for _, comp := range comps {
			builder.Say(comp.Name)
			builder.Pause("200")
			builder.Say(fmt.Sprintf("In %s,", comp.City))
			builder.Pause("200")

			d, err := time.Parse("2006-01-02", comp.Date)
			if err != nil {
				fmt.Printf("Error parsing competition date: %+v", comp)
				continue
			}

			builder.Say(fmt.Sprintf("On %s %d, %d", d.Month().String(), d.Day(), d.Year()))
			builder.Pause("500")
		}
	}

	return alexa.NewSSMLResponse("Upcoming Events", builder.Build(), true)
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder

	builder.Say("You can ask.")
	builder.Pause("500")
	builder.Say("Give me the upcoming jiu jitsu competitions.")
	builder.Pause("500")
	builder.Say("Give me the upcoming grappling competitions.")

	return alexa.NewSSMLResponse("FightCalender Help", builder.Build(), false)
}

func HandleUnknownIntent(request alexa.Request) alexa.Response {
	fmt.Printf("Errored request: %+v\n", request)
	return alexa.NewSimpleResponse("Error", "Sorry, I had trouble doing what you asked. Please try again.", true)
}

func LaunchRequestHandler(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder

	builder.Say("Welcome to fight calendar. Which state do you want to compete in?")

	return alexa.NewSSMLResponse("FightCalender Launch", builder.Build(), false)
}

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response

	switch request.Body.Intent.Name {
	case "UpcomingEventsIntent":
		response = HandleUpcomingEventsIntent(request)
	case alexa.HelpIntent:
		response = HandleHelpIntent(request)
	default:
		response = HandleUnknownIntent(request)
	}

	return response
}

func Handler(request alexa.Request) (alexa.Response, error) {
	fmt.Printf("Request: %+v\n", request)

	if request.Body.Type == "LaunchRequest" {
		return LaunchRequestHandler(request), nil
	}

	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
