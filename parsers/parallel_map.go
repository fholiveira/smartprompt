package parsers

import . "github.com/fholiveira/smartprompt/plugins"

type InputMessage struct {
	plugin Plugin
	token  Token
}

type OutputMessage struct {
	token  Token
	result string
	err    error
}

func runParallel(num InputMessage, function func(InputMessage) OutputMessage) <-chan OutputMessage {
	output := make(chan OutputMessage, 10)
	go func() {
		output <- function(num)
	}()

	return output
}

func ParallelMap(messages []InputMessage, function func(InputMessage) OutputMessage) []OutputMessage {
	channels := make([]<-chan OutputMessage, len(messages))

	for index, message := range messages {
		channel := runParallel(message, function)
		channels[index] = channel
	}

	outputMessages := make([]OutputMessage, len(messages))
	for index, channel := range channels {
		outputMessages[index] = <-channel
	}

	return outputMessages
}
