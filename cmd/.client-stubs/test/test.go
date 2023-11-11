package test_service_tester

import (
	"context"
	"net/http"
	"slash_mochi/cmd/.client-stubs/test_common"
	testv1 "slash_mochi/gen/go/slash_mochi/v1/test"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"connectrpc.com/connect"
)

type TestServiceTester struct {
	connectClient testv1connect.TestServiceClient
}

func NewTestServiceTester() *TestServiceTester {
	return &TestServiceTester{
		connectClient: testv1connect.NewTestServiceClient(
			http.DefaultClient,
			"http://localhost:3081",
		),
	}
}

func (t *TestServiceTester) Test() []test_common.TestResult {
	ret := make([]test_common.TestResult, 0)
	ret = append(ret, t.testLoopback_Normal1())

	return ret
}

func (t *TestServiceTester) testLoopback_Normal1() test_common.TestResult {
	methodName := "testLoopback_Normal1"

	res, err := t.connectClient.Loopback(
		context.Background(),
		connect.NewRequest(&testv1.LoopbackRequest{
			Message: "test",
		}),
	)
	if err != nil {
		return test_common.TestResult{
			methodName,
			false,
		}
	}
	result := res.Msg.Message == "response:test"

	return test_common.TestResult{
		methodName,
		result,
	}
}

// type TestStub struct {
// 	client *testv1connect.TestServiceClient
// }

// func NewTestStub(client *testv1connect.TestServiceClient) *TestStub {
// 	return &TestStub{
// 		client: client,
// 	}
// }

// func (s *TestStub) TestLoopback() (string, bool) {
// 	res, err := s.client.Loopback(
// 		context.Background(),
// 		connect.NewRequest(&testv1.LoopbackRequest{
// 			Message: "hoge-",
// 		}),
// 	)
// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}

// 	log.Println(res.Msg.GetMessage())
// 	return "abc", true
// }
