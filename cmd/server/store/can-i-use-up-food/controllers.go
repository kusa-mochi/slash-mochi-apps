package can_i_use_up_food_store

import (
	"slash_mochi/cmd/server/server_common"
	can_i_use_up_foodv1 "slash_mochi/gen/go/slash_mochi/v1/can_i_use_up_food"
)

type CanIUseUpFoodStoreControllers struct {
	data *CanIUseUpFoodStoreStructures
}

func NewCanIUseUpFoodStoreControllers() *CanIUseUpFoodStoreControllers {
	return &CanIUseUpFoodStoreControllers{
		data: NewCanIUseUpFoodStoreStructures(),
	}
}

func (c *CanIUseUpFoodStoreControllers) NewProject(req *server_common.GetSetRequest[string, can_i_use_up_foodv1.NewProjectRequest]) {
	projectId := GenerateNewId()

	// TODO: put request data to the store

	req.ResChan <- projectId
}

func (c *CanIUseUpFoodStoreControllers) LoadProjectData(req *server_common.GetSetRequest[ProjectData, can_i_use_up_foodv1.LoadProjectDataRequest]) {
	// TODO: check if a project is exist
	// TODO: return the project data to ResChan
}
