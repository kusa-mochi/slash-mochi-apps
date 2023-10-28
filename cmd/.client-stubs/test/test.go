package test

import (
	"context"
	"log"
	testv1 "slash_mochi/gen/go/slash_mochi/v1/test"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"connectrpc.com/connect"
)

type TestStub struct {
	client *testv1connect.TestServiceClient
}

func NewTestStub(client *testv1connect.TestServiceClient) *TestStub {
	return &TestStub{
		client: client,
	}
}

func (s *TestStub) TestLoopback() (string, bool) {
	res, err := s.client.Loopback(
		context.Background(),
		connect.NewRequest(&testv1.LoopbackRequest{
			Message: "hoge-",
		}),
	)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(res.Msg.GetMessage())
	return "abc", true
}
