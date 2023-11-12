package omikuji_service_tester

import (
	"context"
	"net/http"
	"slash_mochi/cmd/.client-stubs/test_kit"
	"slash_mochi/gen/go/slash_mochi/v1/omikuji/omikujiv1connect"

	"google.golang.org/protobuf/types/known/emptypb"

	"connectrpc.com/connect"
)

type OmikujiServiceTester struct {
	connectClient omikujiv1connect.OmikujiServiceClient
}

func NewOmikujiServiceTester(targetUrl string) *OmikujiServiceTester {
	return &OmikujiServiceTester{
		connectClient: omikujiv1connect.NewOmikujiServiceClient(
			http.DefaultClient,
			targetUrl,
		),
	}
}

func (t *OmikujiServiceTester) Test() []test_kit.TestResult {
	return test_kit.RunTests(
		t.testOpen100_Omikuji(),
	)
}

func (t *OmikujiServiceTester) testOpen200_Omikuji() test_kit.TestResult {
	ret := test_kit.NewTestResult()

	nTest := 200
	const nResultPattern int = 13
	var resultFlag [nResultPattern]bool

	for iTest := 0; iTest < nTest; iTest++ {
		res, err := t.connectClient.OpenOmikuji(
			context.Background(),
			connect.NewRequest(&emptypb.Empty{}),
		)
		if err != nil {
			return ret
		}
		resultFlag[res.Msg.Result] = true
	}

	for iFlag := 0; iFlag < nResultPattern; iFlag++ {
		if !resultFlag[iFlag] {
			return ret
		}
	}
	ret.IsSucceeded = true

	return ret
}
