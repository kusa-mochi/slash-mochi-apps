package test_service

import (
	"context"
	"errors"
	"fmt"
	testv1 "slash_mochi/gen/go/slash_mochi/v1/test"

	"connectrpc.com/connect"
)

type TestService struct {
}

func NewTestService() *TestService {
	return &TestService{}
}

// Loopback implements testv1connect.TestServiceHandler.
func (s *TestService) Loopback(
	ctx context.Context,
	req *connect.Request[testv1.LoopbackRequest],
) (*connect.Response[testv1.LoopbackResponse], error) {
	testMessage := req.Msg.GetMessage()
	if testMessage == "" {
		return nil, errors.New("message is empty")
	}
	return connect.NewResponse[testv1.LoopbackResponse](&testv1.LoopbackResponse{
		Message: fmt.Sprintf("response:%s", testMessage),
	}), nil
}
