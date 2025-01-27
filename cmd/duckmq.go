package main

import (
	"fmt"

	"github.com/RobMin/duckmq/pkg/common"
	"github.com/RobMin/duckmq/pkg/dispatcher"
	"github.com/RobMin/duckmq/pkg/message_store"
	"github.com/RobMin/duckmq/pkg/reciever"
)

func main() {
	message_recieve_channel := make(chan common.MessageRequest)
	message_dispatch_channel := make(chan common.MessageRequest)

	reciever.Init(message_recieve_channel)
	message_store.Init(message_recieve_channel, message_dispatch_channel)
	dispatcher.Init(message_dispatch_channel)

	fmt.Print("Finished")
}
