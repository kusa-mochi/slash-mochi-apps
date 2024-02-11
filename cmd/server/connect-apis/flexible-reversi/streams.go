package flexible_reversi_servce

import (
	"fmt"
	"slash_mochi/cmd/server/server_common"
	flexible_reversiv1 "slash_mochi/gen/go/slash_mochi/v1/flexible_reversi"
	"sync"

	"connectrpc.com/connect"
)

type FlexibleReversiStreams struct {
	mu                sync.Mutex
	globalChatStreams map[server_common.UserId]*connect.BidiStream[flexible_reversiv1.ChatToSend, flexible_reversiv1.ChatToReceive]
}

func NewFlexibleReversiStreams() *FlexibleReversiStreams {
	return &FlexibleReversiStreams{
		globalChatStreams: make(map[string]*connect.BidiStream[flexible_reversiv1.ChatToSend, flexible_reversiv1.ChatToReceive]),
	}
}

func (s *FlexibleReversiStreams) AddGlobalChatStream(
	key string,
	stream *connect.BidiStream[flexible_reversiv1.ChatToSend, flexible_reversiv1.ChatToReceive],
) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.globalChatStreams[key] = stream
}

func (s *FlexibleReversiStreams) RemoveGlobalChatStream(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.globalChatStreams, key)
}

func (s *FlexibleReversiStreams) GetGlobalChatStream(key string) (*connect.BidiStream[flexible_reversiv1.ChatToSend, flexible_reversiv1.ChatToReceive], error) {
	if stream, ok := s.globalChatStreams[key]; ok {
		return stream, nil
	}

	return nil, fmt.Errorf("stream [\"%s\"] not found", key)
}
