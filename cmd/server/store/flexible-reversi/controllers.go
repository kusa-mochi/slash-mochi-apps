package flexible_reversi_store

import (
	"fmt"
	"slash_mochi/cmd/server/server_common"
)

type FlexibleReversiStoreControllers struct {
	data FlexibleReversiStoreStructures
}

func NewFlexibleReversiControllers() *FlexibleReversiStoreControllers {
	return &FlexibleReversiStoreControllers{
		data: *NewFlexibleReversiStructures(),
	}
}

func (c *FlexibleReversiStoreControllers) ValidateUserId(req *server_common.GetSetRequest[bool, string]) {
	userId := req.DataToSet
	isCorrectUserId := c.data.ValidateUserId(userId)
	req.ResChan <- isCorrectUserId
}

func (c *FlexibleReversiStoreControllers) GlobalChat(req *server_common.GetSetRequest[ChatItem, ChatHistoryItem]) {
	chatHistoryItem := req.DataToSet
	messageToBroadCast := fmt.Sprintf("hoge %s", chatHistoryItem.message)
	chatHistoryItemToBroadcast := NewChatHistoryItem(messageToBroadCast, chatHistoryItem.userId)

	// add chat history to the store
	c.data.AppendGlobalChatHistory(*chatHistoryItemToBroadcast)

	// get a user name from the store
	userInfo, err := c.data.GetUserInfo(chatHistoryItem.userId)
	if err != nil {
		req.ResChan <- *NewChatItem(
			"",
			"",
		)
		return
	}

	req.ResChan <- *NewChatItem(
		chatHistoryItem.message,
		userInfo.name,
	)
}
