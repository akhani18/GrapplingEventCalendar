package main

import (
	"fmt"

	"github.com/akhani18/GrapplingEventCalendar/alexa"
	"github.com/aws/aws-lambda-go/lambda"
)

// Upcoming events logic goes here.
func HandleUpcomingEventsIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder

	stateName := request.Body.Intent.Slots["state"].Value

	builder.Say(fmt.Sprintf("Here are some of the upcoming competitions in %s.", stateName))
	builder.Pause("500")

	builder.Say("Grappling Industries Pheonix.")
	builder.Pause("500")
	builder.Say("In Scotsdale, Arizona.")
	builder.Pause("500")
	builder.Say("On January 4, 2020.")

	builder.Pause("1000")

	builder.Say("NAGA San Diego Grappling Championship.")
	builder.Pause("500")
	builder.Say("In San Diego, California.")
	builder.Pause("500")
	builder.Say("On February 22, 2020.")

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
