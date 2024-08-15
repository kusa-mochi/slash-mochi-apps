package can_i_use_up_food_store

import (
	"slash_mochi/cmd/server/server_common"
	can_i_use_up_foodv1 "slash_mochi/gen/go/slash_mochi/v1/can_i_use_up_food"
)

type CanIUseUpFoodStoreInterfaces struct {
	// BroadcastGlobalChatRequest chan *server_common.SetRequest[*flexible_reversiv1.ChatToReceive]
	// GlobalChatRequest          chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem]
	AddFoodToList          chan *server_common.SetRequest[can_i_use_up_foodv1.Food]
	LoadProjectDataRequest chan *server_common.GetSetRequest[ProjectData, can_i_use_up_foodv1.LoadProjectDataRequest]
	NewProjectRequest      chan *server_common.GetSetRequest[string, can_i_use_up_foodv1.NewProjectRequest]
}

func NewCanIUseUpFoodStoreInterfaces() *CanIUseUpFoodStoreInterfaces {
	return &CanIUseUpFoodStoreInterfaces{
		// BroadcastGlobalChatRequest: make(chan *server_common.SetRequest[*flexible_reversiv1.ChatToReceive], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		// GlobalChatRequest:          make(chan *server_common.GetSetRequest[ChatItem, ChatHistoryItem], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		AddFoodToList:          make(chan *server_common.SetRequest[can_i_use_up_foodv1.Food], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		LoadProjectDataRequest: make(chan *server_common.GetSetRequest[ProjectData, can_i_use_up_foodv1.LoadProjectDataRequest], server_common.STORE_INTERFACE_CHANNELS_SIZE),
		NewProjectRequest:      make(chan *server_common.GetSetRequest[string, can_i_use_up_foodv1.NewProjectRequest], server_common.STORE_INTERFACE_CHANNELS_SIZE),
	}
}
