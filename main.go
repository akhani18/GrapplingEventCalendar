package main

import (
	"github.com/akhani18/GrapplingEventCalendar/alexa"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleUpcomingEventsIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Upcoming Events", "Upcoming events data here")
}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("Here are some of the things you can ask.")
	builder.Pause("1000")
	builder.Say("Give me the upcoming events.")
	builder.Pause("1000")
	builder.Say("What are the upcoming jiujitsu events.")
	builder.Pause("1000")
	builder.Say("What are the upcoming grappling events.")
	return alexa.NewSSMLResponse("Slick Dealer Help", builder.Build())
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
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
