package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/arienmalec/alexa-go"
	"time"
)

// function to delegate the voice intent to appropriate handler
func IntentDelegator(request alexa.Request) alexa.Response {
	var response alexa.Response
	if request.Body.Type == "LaunchRequest" {
		response = lauchPrompt(request)
	}
	switch request.Body.Intent.Name {
	case "DateIntent":
		response = handleDate(request)
	case "TimeIntent":
		response = handleTime(request)
	case alexa.HelpIntent:
		response = handleHelp()
	case alexa.StopIntent:
		response = handleStop()
	}

	return response
}
// function invoke during intial skill launch
func lauchPrompt(request alexa.Request) alexa.Response {
	title := "Launch golang Alexa"
	var text string
	//text = "A simple alexa skill using golang, say hello to receive a greeting new response"
	text = "Welcome to Simple date and time skill"
	return SimpleResponse(title, text)
}
// function invoke during Date Intent
func handleDate(request alexa.Request) alexa.Response {
	title := "Current date"
	var text string
	t := time.Now()
	text = "Current date is "+t.Format("Mon Jan _2 2006")
	return SimpleResponse(title, text)
}
// function invoke during Time Intent
func handleTime(request alexa.Request) alexa.Response {
	title := "Current time"
	var text string
	t := time.Now()
	text = "Current time is "+t.Format("3:04PM")
	return SimpleResponse(title, text)
}

// function invoke during help Intent, with options to use the skills.
func handleHelp() alexa.Response {
	return SimpleResponse("Help for Current Date & Time", "Simple date and time skill, ask current date and time")
}

// function invoke during close the skill session, with thanks note.
func handleStop() alexa.Response {
	//close the session.
	return alexa.NewSimpleResponse("Thanks", "Thanks for using date and time skill")
}

// lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDelegator(request), nil
}

// Simple response keep the session live
func SimpleResponse(title string, text string) alexa.Response {
	r := alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "PlainText",
				Text: text,
			},
			Card: &alexa.Payload{
				Type:    "Simple",
				Title:   title,
				Content: text,
			},
			ShouldEndSession: false,
		},
	}
	return r
}

// Trigger
func main() {
	lambda.Start(Handler)
}
