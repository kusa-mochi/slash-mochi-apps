package flexible_reversi_store

import (
	"slash_mochi/cmd/server/server_common"
)

type FlexibleReversiStoreInterfaces struct {
	GlobalChatRequest     chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem]
	ValidateUserIdRequest chan *server_common.GetSetRequest[bool, string]
}

func NewFlexibleReversiStoreInterfaces() *FlexibleReversiStoreInterfaces {
	return &FlexibleReversiStoreInterfaces{
		GlobalChatRequest:     make(chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		ValidateUserIdRequest: make(chan *server_common.GetSetRequest[bool, string], server_common.STORE_INTERFACE_CHANNELS_SIZE),
	}
}
