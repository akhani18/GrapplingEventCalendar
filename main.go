package main

import (
	"fmt"

	"github.com/akhani18/GrapplingEventCalendar/alexa"
	"github.com/aws/aws-lambda-go/lambda"
)

// Upcoming events logic goes here.
func HandleUpcomingEventsIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder

	builder.Say("Here are some of the upcoming grappling events.")
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

	return alexa.NewSSMLResponse("Upcoming Events", builder.Build())
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("Here are some of the things you can ask.")
	builder.Pause("1000")
	builder.Say("Give me the upcoming events.")
	builder.Pause("1000")
	builder.Say("What are the upcoming jiu jitsu events.")
	builder.Pause("1000")
	builder.Say("What are the upcoming grappling events.")
	return alexa.NewSSMLResponse("GrapplingEventCalendar Help", builder.Build())
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("About", "An alexa voice skill that provides information about upcoming grappling events.")
}

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response

	switch request.Body.Intent.Name {
	case "UpcomingEventsIntent":
		response = HandleUpcomingEventsIntent(request)
	case "HelpIntent":
		response = HandleHelpIntent(request)
	case "AboutIntent":
		response = HandleAboutIntent(request)
	default:
		response = HandleHelpIntent(request)
	}

	return response
}

func Handler(request alexa.Request) (alexa.Response, error) {
	fmt.Printf("Request: %+v\n", request)
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
