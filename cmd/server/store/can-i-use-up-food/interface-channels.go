package can_i_use_up_food_store

type CanIUseUpFoodStoreInterfaces struct {
	// BroadcastGlobalChatRequest chan *server_common.SetRequest[*flexible_reversiv1.ChatToReceive]
	// GlobalChatRequest          chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem]
}

func NewCanIUseUpFoodStoreInterfaces() *CanIUseUpFoodStoreInterfaces {
	return &CanIUseUpFoodStoreInterfaces{
		// BroadcastGlobalChatRequest: make(chan *server_common.SetRequest[*flexible_reversiv1.ChatToReceive], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		// GlobalChatRequest:          make(chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem], server_common.STORE_INTERFACE_CHANNELS_SIZE),
	}
}
