package flexible_reversi_store

import (
	"fmt"
	"slash_mochi/cmd/server/server_common"
	flexible_reversiv1 "slash_mochi/gen/go/slash_mochi/v1/flexible_reversi"

	"connectrpc.com/connect"
)

type FlexibleReversiStoreControllers struct {
	data    *FlexibleReversiStoreStructures
	streams *FlexibleReversiStreams
}

func NewFlexibleReversiControllers() *FlexibleReversiStoreControllers {
	return &FlexibleReversiStoreControllers{
		data:    NewFlexibleReversiStructures(),
		streams: NewFlexibleReversiStreams(),
	}
}

type GlobalChatStreamItem struct {
	UserId string
	Stream *connect.BidiStream[flexible_reversiv1.ChatToSend, flexible_reversiv1.ChatToReceive]
}

func NewGlobalChatStreamItem(userId string, stream *connect.BidiStream[flexible_reversiv1.ChatToSend, flexible_reversiv1.ChatToReceive]) *GlobalChatStreamItem {
	return &GlobalChatStreamItem{
		UserId: userId,
		Stream: stream,
	}
}

func (c *FlexibleReversiStoreControllers) AddGlobalChatStream(req *server_common.SetRequest[GlobalChatStreamItem]) {
	globalChatStreamItem := req.Data
	c.streams.AddGlobalChatStream(globalChatStreamItem.UserId, globalChatStreamItem.Stream)
}

func (c *FlexibleReversiStoreControllers) BroadcastGlobalChat(req *server_common.SetRequest[*flexible_reversiv1.ChatToReceive]) {
	chatData := req.Data
	ret := true
	for _, s := range c.streams.globalChatStreams {
		err := s.Send(&flexible_reversiv1.ChatToReceive{
			Message:  chatData.GetMessage(),
			UserName: chatData.GetUserName(),
		})
		if err != nil {
			ret = false
		}
	}
	req.ResChan <- ret
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
