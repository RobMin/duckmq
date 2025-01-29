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

	go message_store.Init(message_recieve_channel, message_dispatch_channel)
	go dispatcher.Init(message_dispatch_channel)

	reciever.Init(message_recieve_channel)

	fmt.Println("Finished")
}
