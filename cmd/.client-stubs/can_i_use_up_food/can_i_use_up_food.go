package can_i_use_up_food_service_tester

import (
	"context"
	"log"
	"net/http"
	"slash_mochi/cmd/.client-stubs/test_kit"
	can_i_use_up_foodv1 "slash_mochi/gen/go/slash_mochi/v1/can_i_use_up_food"
	"slash_mochi/gen/go/slash_mochi/v1/can_i_use_up_food/can_i_use_up_foodv1connect"

	"connectrpc.com/connect"
)

type CanIUseUpFoodServiceTester struct {
	connectClient can_i_use_up_foodv1connect.CanIUseUpFoodServiceClient
}

func NewCanIUseUpFoodServiceTester(targetUrl string) *CanIUseUpFoodServiceTester {
	return &CanIUseUpFoodServiceTester{
		connectClient: can_i_use_up_foodv1connect.NewCanIUseUpFoodServiceClient(
			http.DefaultClient,
			targetUrl,
		),
	}
}

func (t *CanIUseUpFoodServiceTester) Test() []test_kit.TestResult {
	return test_kit.RunTests(
		t.testNewProject(),
	)
}

func (t *CanIUseUpFoodServiceTester) testNewProject() test_kit.TestResult {
	ret := test_kit.NewTestResult()

	res, err := t.connectClient.NewProject(
		context.Background(),
		connect.NewRequest(&can_i_use_up_foodv1.NewProjectRequest{
			ProjectName: "てすとぷろじぇくと！",
			StartDate: &can_i_use_up_foodv1.Date{
				Year:  2024,
				Month: 5,
				Date:  5,
			},
			EndDate: &can_i_use_up_foodv1.Date{
				Year:  2024,
				Month: 5,
				Date:  31,
			},
		}),
	)
	if err != nil {
		return ret
	}

	log.Println(res.Msg.GetId())

	ret.IsSucceeded = true
	return ret
}
