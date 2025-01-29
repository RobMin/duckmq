package dispatcher

import (
	"fmt"

	"github.com/RobMin/duckmq/pkg/common"
)

func Init(message_dispatch_channel chan common.MessageRequest) {
	for true {
		message := <-message_dispatch_channel
		dispatchMessage(message)
	}
}

func dispatchMessage(message common.MessageRequest) {
	log := fmt.Sprintf("Dispatching: Id=%s, Timestamp=%s, Message=%s", message.Id, message.Timestamp.String(), message.Message)
	fmt.Println(log)
}
