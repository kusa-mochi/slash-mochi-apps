package can_i_use_up_food_store

type CanIUseUpFoodStore struct {
	Interfaces  *CanIUseUpFoodStoreInterfaces
	Controllers *CanIUseUpFoodStoreControllers
}

func NewCanIUseUpFoodStore() *CanIUseUpFoodStore {
	return &CanIUseUpFoodStore{
		Interfaces:  NewCanIUseUpFoodStoreInterfaces(),
		Controllers: NewCanIUseUpFoodStoreControllers(),
	}
}
