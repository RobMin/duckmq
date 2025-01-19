package message_store

import "github.com/RobMin/duckmq/pkg/common"

func Init(message_recieve_channel chan common.MessageRequest, message_dispatch_channel chan common.MessageRequest) {
	message := <-message_recieve_channel
	message_dispatch_channel <- message
}
