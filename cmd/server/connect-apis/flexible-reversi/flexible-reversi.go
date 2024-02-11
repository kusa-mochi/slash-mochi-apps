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

		// TODO: send to ALL clients who connect to Flexible Reversi
		if err := bs.Send(&flexible_reversiv1.ChatToReceive{
			Message:  newChatItem.Message,
			UserName: newChatItem.UserName,
		}); err != nil {
			return connect.NewError(connect.CodeInternal, fmt.Errorf("failed to broadcast global chat:%w", err))
		}
	}
}

/////////// private methods

func (s *FlexibleReversiService) 
