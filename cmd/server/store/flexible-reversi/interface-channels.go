package flexible_reversi_store

import (
	"slash_mochi/cmd/server/server_common"
	flexible_reversiv1 "slash_mochi/gen/go/slash_mochi/v1/flexible_reversi"
)

type FlexibleReversiStoreInterfaces struct {
	AddGlobalChatStreamRequest chan *server_common.SetRequest[GlobalChatStreamItem]
	BroadcastGlobalChatRequest chan *server_common.SetRequest[*flexible_reversiv1.ChatToReceive]
	GlobalChatRequest          chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem]
	ValidateUserIdRequest      chan *server_common.GetSetRequest[bool, string]
}

func NewFlexibleReversiStoreInterfaces() *FlexibleReversiStoreInterfaces {
	return &FlexibleReversiStoreInterfaces{
		AddGlobalChatStreamRequest: make(chan *server_common.SetRequest[GlobalChatStreamItem], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		BroadcastGlobalChatRequest: make(chan *server_common.SetRequest[*flexible_reversiv1.ChatToReceive], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		GlobalChatRequest:          make(chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		ValidateUserIdRequest:      make(chan *server_common.GetSetRequest[bool, string], server_common.STORE_INTERFACE_CHANNELS_SIZE),
	}
}
