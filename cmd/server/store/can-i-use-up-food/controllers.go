package can_i_use_up_food_store

import (
	"log"
	"slash_mochi/cmd/server/server_common"
	can_i_use_up_foodv1 "slash_mochi/gen/go/slash_mochi/v1/can_i_use_up_food"
	"time"
)

type CanIUseUpFoodStoreControllers struct {
	data *CanIUseUpFoodStoreStructures
}

func NewCanIUseUpFoodStoreControllers() *CanIUseUpFoodStoreControllers {
	return &CanIUseUpFoodStoreControllers{
		data: NewCanIUseUpFoodStoreStructures(),
	}
}

func (c *CanIUseUpFoodStoreControllers) AddFoodToList(req *server_common.SetRequest[can_i_use_up_foodv1.Food]) {

	// TODO

	req.ResChan <- true
}

func (c *CanIUseUpFoodStoreControllers) LoadProjectData(req *server_common.GetSetRequest[ProjectData, can_i_use_up_foodv1.LoadProjectDataRequest]) {
	projectId := req.DataToSet.GetProjectId()

	// if a project is exist,
	if projectData, ok := c.data.projects[projectId]; ok {
		log.Printf("project %s found", projectId)

		// prevent data r/w conflicts by copying data.
		// Connect handler can only read copied data.
		req.ResChan <- *projectData.Clone()
	} else {
		log.Printf("[E] project %s not found", projectId)

		// return empty data
		req.ResChan <- ProjectData{}
	}
}

func (c *CanIUseUpFoodStoreControllers) NewProject(req *server_common.GetSetRequest[string, can_i_use_up_foodv1.NewProjectRequest]) {
	startDate := req.DataToSet.GetStartDate()
	startDatetime := time.Date(
		int(startDate.GetYear()),
		time.Month(startDate.GetMonth()),
		int(startDate.GetDate()),
		0,
		0,
		0,
		0,
		time.Local,
	)

	endDate := req.DataToSet.GetEndDate()
	endDatetime := time.Date(
		int(endDate.GetYear()),
		time.Month(endDate.GetMonth()),
		int(endDate.GetDate()),
		0,
		0,
		0,
		0,
		time.Local,
	)

	projectId := GenerateNewId()

	// put request data to the store
	c.data.projects[projectId] = NewProjectData(
		req.DataToSet.GetProjectName(),
		startDatetime,
		endDatetime,
	)

	req.ResChan <- projectId
}
