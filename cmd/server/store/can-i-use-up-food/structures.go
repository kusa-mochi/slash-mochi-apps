package can_i_use_up_food_store

import (
	"crypto/sha256"
	"fmt"
	"io"
	can_i_use_up_foodv1 "slash_mochi/gen/go/slash_mochi/v1/can_i_use_up_food"
	"time"
)

type CanIUseUpFoodStoreStructures struct {
	projects map[string]*ProjectData
}

func NewCanIUseUpFoodStoreStructures() *CanIUseUpFoodStoreStructures {
	// TODO: this is dummy data for debugging
	var projects map[string]*ProjectData = map[string]*ProjectData{}
	location, _ := time.LoadLocation("UTC")
	var dummyData *ProjectData = NewProjectData(
		"テストサーバープロジェクト",
		time.Date(2024, 8, 12, 0, 0, 0, 0, location),
		time.Date(2024, 9, 10, 0, 0, 0, 0, location),
	)
	var dummyProjectId string = "abc"
	projects[dummyProjectId] = dummyData

	return &CanIUseUpFoodStoreStructures{
		// TODO: this is dummy data for dubugging
		projects: projects,
		// projects: map[string]*ProjectData{},
	}
}

// ID generator
func GenerateNewId() string {
	s256 := sha256.New()
	now := time.Now().Format("2000-01-01 11:11:11.000000000")
	io.WriteString(s256, now)

	h := s256.Sum(nil)
	hstr := fmt.Sprintf("%x", h)

	return hstr
}

// ProjectData
// use this for store response.

type ProjectData struct {
	projectName        string
	startDate          time.Time
	endDate            time.Time
	foods              []*Food
	calendarItemGroups []*CalendarItemGroup
}

func NewProjectData(projectName string, startDate time.Time, endDate time.Time) *ProjectData {
	return &ProjectData{
		projectName:        projectName,
		startDate:          startDate,
		endDate:            endDate,
		foods:              make([]*Food, 0),
		calendarItemGroups: make([]*CalendarItemGroup, 0),
	}
}

func (d *ProjectData) Clone() *ProjectData {
	// make copy of foods
	foods := make([]*Food, 0)
	for _, food := range d.foods {
		copiedFood := food.Clone()
		foods = append(foods, copiedFood)
	}

	// make copy of calendarItemGroups
	calendarItemGroups := make([]*CalendarItemGroup, 0)
	for _, group := range d.calendarItemGroups {
		copiedGroup := group.Clone()
		calendarItemGroups = append(calendarItemGroups, copiedGroup)
	}

	return &ProjectData{
		projectName:        d.projectName,
		startDate:          d.startDate.AddDate(0, 0, 0), // copy Time struct
		endDate:            d.endDate.AddDate(0, 0, 0),   // copy Time struct
		foods:              foods,
		calendarItemGroups: calendarItemGroups,
	}
}

// ProjectUri

type ProjectUri struct {
	projectId string
}

func NewProjectUri() *ProjectUri {
	return &ProjectUri{
		projectId: GenerateNewId(),
	}
}

// Food

type Food struct {
	id            string
	name          string
	totalAmount   int
	amountUnit    can_i_use_up_foodv1.AmountUnit
	limitDate     time.Time
	calendarItems []*CalendarItem
}

func NewFood(
	id string,
	name string,
	totalAmount int,
	amountUnit can_i_use_up_foodv1.AmountUnit,
	limitDate time.Time,
) *Food {
	return &Food{
		id:            id,
		name:          name,
		totalAmount:   totalAmount,
		amountUnit:    amountUnit,
		limitDate:     limitDate,
		calendarItems: make([]*CalendarItem, 0),
	}
}

func (f *Food) Clone() *Food {
	calendarItems := make([]*CalendarItem, 0)
	for _, item := range f.calendarItems {
		copiedItem := *item
		calendarItems = append(calendarItems, &copiedItem)
	}
	return &Food{
		id:            f.id,
		name:          f.name,
		totalAmount:   f.totalAmount,
		limitDate:     f.limitDate,
		calendarItems: calendarItems,
	}
}

// Calendar Item

type CalendarItem struct {
	id       string
	date     time.Time
	mealTime can_i_use_up_foodv1.MealTime
	amount   int
}

func NewCalendarItem(
	id string,
	date time.Time,
	mealTime can_i_use_up_foodv1.MealTime,
	amount int,
) *CalendarItem {
	return &CalendarItem{
		id:       id,
		date:     date,
		mealTime: mealTime,
		amount:   amount,
	}
}

// Calendar Item Group

type CalendarItemGroup struct {
	name            string
	calendarItemIds []string
}

func NewCalendarItemGroup(name string) *CalendarItemGroup {
	return &CalendarItemGroup{
		name:            name,
		calendarItemIds: make([]string, 0),
	}
}

func (g *CalendarItemGroup) Clone() *CalendarItemGroup {
	copiedIds := make([]string, 0)
	copy(copiedIds, g.calendarItemIds)

	return &CalendarItemGroup{
		name:            g.name,
		calendarItemIds: copiedIds,
	}
}
