package flexible_reversi_store

type FlexibleReversiStore struct {
	Interfaces  *FlexibleReversiStoreInterfaces
	Controllers *FlexibleReversiStoreControllers
}

func NewFlexibleReversiStore() *FlexibleReversiStore {
	return &FlexibleReversiStore{
		Interfaces:  NewFlexibleReversiStoreInterfaces(),
		Controllers: NewFlexibleReversiControllers(),
	}
}
