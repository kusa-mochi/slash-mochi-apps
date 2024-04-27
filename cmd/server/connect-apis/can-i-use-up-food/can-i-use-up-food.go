package can_i_use_up_food_service

import (
	"context"
	can_i_use_up_food_store "slash_mochi/cmd/server/store/can-i-use-up-food"
	can_i_use_up_foodv1 "slash_mochi/gen/go/slash_mochi/v1/can_i_use_up_food"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CanIUseUpFoodService struct {
	storeInterface *can_i_use_up_food_store.CanIUseUpFoodStoreInterfaces
}

// AddFoodToList implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) AddFoodToList(context.Context, *connect.Request[can_i_use_up_foodv1.AddFoodToListRequest]) (*connect.Response[can_i_use_up_foodv1.FoodList], error) {
	panic("unimplemented")
}

// LoadProjectData implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) LoadProjectData(context.Context, *connect.Request[can_i_use_up_foodv1.LoadProjectDataRequest]) (*connect.Response[can_i_use_up_foodv1.ProjectData], error) {
	panic("unimplemented")
}

// ModifyCalendarPeriod implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) ModifyCalendarPeriod(context.Context, *connect.Request[can_i_use_up_foodv1.ModifyCalendarPeriodRequest]) (*connect.Response[emptypb.Empty], error) {
	panic("unimplemented")
}

// ModifyFoodAmountInCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) ModifyFoodAmountInCalendar(context.Context, *connect.Request[can_i_use_up_foodv1.ModifyFoodAmountInCalendarRequest]) (*connect.Response[can_i_use_up_foodv1.FoodList], error) {
	panic("unimplemented")
}

// ModifyFoodInList implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) ModifyFoodInList(context.Context, *connect.Request[can_i_use_up_foodv1.ModifyFoodInListRequest]) (*connect.Response[can_i_use_up_foodv1.FoodList], error) {
	panic("unimplemented")
}

// MoveFoodInCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) MoveFoodInCalendar(context.Context, *connect.Request[can_i_use_up_foodv1.MoveFoodInCalendarRequest]) (*connect.Response[can_i_use_up_foodv1.FoodList], error) {
	panic("unimplemented")
}

// NewProject implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) NewProject(context.Context, *connect.Request[can_i_use_up_foodv1.NewProjectRequest]) (*connect.Response[can_i_use_up_foodv1.ProjectUri], error) {
	panic("unimplemented")
}

// PutFoodToCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) PutFoodToCalendar(context.Context, *connect.Request[can_i_use_up_foodv1.PutFoodToCalendarRequest]) (*connect.Response[can_i_use_up_foodv1.FoodList], error) {
	panic("unimplemented")
}

// RemoveFoodFromCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) RemoveFoodFromCalendar(context.Context, *connect.Request[can_i_use_up_foodv1.RemoveFoodFromCalendarRequest]) (*connect.Response[can_i_use_up_foodv1.FoodList], error) {
	panic("unimplemented")
}

// RemoveFoodsFromList implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (*CanIUseUpFoodService) RemoveFoodsFromList(context.Context, *connect.Request[can_i_use_up_foodv1.RemoveFoodsFromListRequest]) (*connect.Response[can_i_use_up_foodv1.FoodList], error) {
	panic("unimplemented")
}

func NewCanIUseUpFoodService(channels *can_i_use_up_food_store.CanIUseUpFoodStoreInterfaces) *CanIUseUpFoodService {
	return &CanIUseUpFoodService{
		storeInterface: channels,
	}
}
