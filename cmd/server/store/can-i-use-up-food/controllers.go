package can_i_use_up_food_store

type CanIUseUpFoodStoreControllers struct {
	data *CanIUseUpFoodStoreStructures
}

func NewCanIUseUpFoodStoreControllers() *CanIUseUpFoodStoreControllers {
	return &CanIUseUpFoodStoreControllers{
		data: NewCanIUseUpFoodStoreStructures(),
	}
}
