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
	return &CanIUseUpFoodStoreStructures{
		projects: map[string]*ProjectData{},
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
