package message_store

import (
	"fmt"

	"github.com/RobMin/duckmq/pkg/common"
)

func Init(message_recieve_channel chan common.MessageRequest, message_dispatch_channel chan common.MessageRequest) {
	for true {
		message := <-message_recieve_channel
		storeMessage(message)
		message_dispatch_channel <- message
	}
}

var store = []common.MessageRequest{}

func storeMessage(message common.MessageRequest) {
	store = append(store, message)
	log := fmt.Sprintf("In message store: %s", message.Id)
	fmt.Println(log)
}
