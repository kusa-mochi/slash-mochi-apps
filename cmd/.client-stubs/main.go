package main

import (
	"context"
	"log"
	"net/http"
	testv1 "slash_mochi/gen/go/slash_mochi/v1/test"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"connectrpc.com/connect"
)

func main() {
	serverAddrPort := "http://localhost:3081"
	testClient := testv1connect.NewTestServiceClient(
		http.DefaultClient,
		serverAddrPort,
	)

	res, err := testClient.Loopback(
		context.Background(),
		connect.NewRequest(&testv1.LoopbackRequest{
			Message: "hoge-",
		}),
	)
	if err != nil {
		log.Println(err)
	}

	log.Println(res.Msg.GetMessage())
}
