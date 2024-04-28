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

func NewCanIUseUpFoodService(channels *can_i_use_up_food_store.CanIUseUpFoodStoreInterfaces) *CanIUseUpFoodService {
	return &CanIUseUpFoodService{
		storeInterface: channels,
	}
}

// AddFoodToList implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) AddFoodToList(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.AddFoodToListRequest]) (*connect.Response[can_i_use_up_foodv1.NewID], error) {
	panic("unimplemented")
}

// LoadProjectData implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) LoadProjectData(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.LoadProjectDataRequest]) (*connect.Response[can_i_use_up_foodv1.ProjectData], error) {
	panic("unimplemented")
}

// MakeCalendarItemGroup implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) MakeCalendarItemGroup(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.MakeCalendarItemGroupRequest]) (*connect.Response[can_i_use_up_foodv1.NewID], error) {
	panic("unimplemented")
}

// ModifyCalendarPeriod implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) ModifyCalendarPeriod(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.ModifyCalendarPeriodRequest]) (*connect.Response[emptypb.Empty], error) {
	panic("unimplemented")
}

// ModifyFoodAmountInCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) ModifyFoodAmountInCalendar(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.ModifyFoodAmountInCalendarRequest]) (*connect.Response[emptypb.Empty], error) {
	panic("unimplemented")
}

// ModifyFoodInList implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) ModifyFoodInList(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.ModifyFoodInListRequest]) (*connect.Response[emptypb.Empty], error) {
	panic("unimplemented")
}

// MoveFoodInCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) MoveFoodInCalendar(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.MoveFoodInCalendarRequest]) (*connect.Response[emptypb.Empty], error) {
	panic("unimplemented")
}

// NewProject implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) NewProject(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.NewProjectRequest],
) (*connect.Response[can_i_use_up_foodv1.ProjectUri], error) {
	panic("unimplemented")
}

// PutFoodToCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) PutFoodToCalendar(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.PutFoodToCalendarRequest]) (*connect.Response[can_i_use_up_foodv1.NewID], error) {
	panic("unimplemented")
}

// RemoveFoodFromCalendar implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) RemoveFoodFromCalendar(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.RemoveFoodFromCalendarRequest]) (*connect.Response[emptypb.Empty], error) {
	panic("unimplemented")
}

// RemoveFoodsFromList implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) RemoveFoodsFromList(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.RemoveFoodsFromListRequest],
) (*connect.Response[emptypb.Empty], error) {
	panic("unimplemented")
}

// UpdateAllData implements can_i_use_up_foodv1connect.CanIUseUpFoodServiceHandler.
func (c *CanIUseUpFoodService) UpdateAllData(
	ctx context.Context,
	req *connect.Request[can_i_use_up_foodv1.UpdateAllDataRequest],
	stream *connect.ServerStream[can_i_use_up_foodv1.ProjectData],
) error {
	panic("unimplemented")
}
