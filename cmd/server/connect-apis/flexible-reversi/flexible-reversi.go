package flexible_reversi_servce

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"slash_mochi/cmd/server/server_common"
	flexible_reversi_store "slash_mochi/cmd/server/store/flexible-reversi"
	flexible_reversiv1 "slash_mochi/gen/go/slash_mochi/v1/flexible_reversi"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FlexibleReversiService struct {
	storeInterface *flexible_reversi_store.FlexibleReversiStoreInterfaces
}

func NewFlexibleReversiService(channels *flexible_reversi_store.FlexibleReversiStoreInterfaces) *FlexibleReversiService {
	return &FlexibleReversiService{
		storeInterface: channels,
	}
}

// SubscribeRoomList implements flexible_reversiv1connect.FlexibleReversiServiceHandler.
func (*FlexibleReversiService) SubscribeRoomList(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
	ss *connect.ServerStream[flexible_reversiv1.RoomList],
) error {
	panic("unimplemented")
}

// GlobalChat implements flexible_reversiv1connect.FlexibleReversiServiceHandler.
func (s *FlexibleReversiService) GlobalChat(
	ctx context.Context,
	bs *connect.BidiStream[flexible_reversiv1.ChatToSend, flexible_reversiv1.ChatToReceive],
) error {
	// get a user ID from a context
	userId, err := server_common.GetUserIdFromContext(ctx)
	if err != nil {
		return connect.NewError(connect.CodeAborted, fmt.Errorf("cannot get user ID according to current context:%w", err))
	}

	// validate a user ID
	validationResChan := make(chan bool)
	s.storeInterface.ValidateUserIdRequest <- &server_common.GetSetRequest[bool, string]{
		DataToSet: userId,
		ResChan:   validationResChan,
	}
	if isCorrectUserId := <-validationResChan; !isCorrectUserId {
		log.Println("invalid user ID @ GlobalChat")
		return connect.NewError(connect.CodeAborted, fmt.Errorf("invalid user ID"))
	}

	// add a stream to the stream map
	addGlobalChatStreamResChan := make(chan bool)
	s.storeInterface.AddGlobalChatStreamRequest <- &server_common.SetRequest[flexible_reversi_store.GlobalChatStreamItem]{
		Data:    *flexible_reversi_store.NewGlobalChatStreamItem(userId, bs),
		ResChan: addGlobalChatStreamResChan,
	}
	if isAdded := <-addGlobalChatStreamResChan; !isAdded {
		log.Println("failed to add a stream to the store")
		return connect.NewError(connect.CodeInternal, fmt.Errorf("failed to add a stream to the store. user ID=%s", userId))
	}

	for {
		msg, err := bs.Receive()
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			if ctx.Err() == context.Canceled {

			}
			return connect.NewError(connect.CodeInternal, fmt.Errorf("failed to receive global chat:%w", err))
		}

		newChatItemResChan := make(chan flexible_reversi_store.ChatItem)
		s.storeInterface.GlobalChatRequest <- &server_common.GetSetRequest[flexible_reversi_store.ChatItem, flexible_reversi_store.ChatHistoryItem]{
			DataToSet: *flexible_reversi_store.NewChatHistoryItem(
				msg.GetMessage(),
				msg.GetUserId(),
			),
			ResChan: newChatItemResChan,
		}
		newChatItem := <-newChatItemResChan
		if newChatItem.UserName == "" {
			return connect.NewError(connect.CodeAborted, fmt.Errorf("user info not found"))
		}

		// send to ALL clients who connect to Flexible Reversi
		if chatResult := s.BloadcastGlobalChat(&flexible_reversiv1.ChatToReceive{
			Message:  newChatItem.Message,
			UserName: newChatItem.UserName,
		}); chatResult != nil {
			log.Println("failed to send a global chat to someone")
		}
	}
}

/////////// private methods

func (s *FlexibleReversiService) BloadcastGlobalChat(chatData *flexible_reversiv1.ChatToReceive) error {
	chatResultResChan := make(chan bool)

	// send a chat to all clients who is connecting to this game.
	s.storeInterface.BroadcastGlobalChatRequest <- &server_common.SetRequest[*flexible_reversiv1.ChatToReceive]{
		Data:    chatData,
		ResChan: chatResultResChan,
	}

	// if failed to send a chat to someone,
	if chatResult := <-chatResultResChan; !chatResult {
		return connect.NewError(connect.CodeUnknown, fmt.Errorf("failed to broadcast global chat"))
	}

	return nil
}
