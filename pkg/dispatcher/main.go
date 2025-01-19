package dispatcher

import (
	"fmt"

	"github.com/RobMin/duckmq/pkg/common"
)

func Init(message_dispatch_channel chan common.MessageRequest) {
	message := <-message_dispatch_channel
	fmt.Printf("Dispatching: Id=%s, Timestamp=%s, Message=%s", message.Id, message.Timestamp.String(), message.Message)
}
