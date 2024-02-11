package server_common

type GetRequest[DataType any] struct {
	ResChan chan DataType
}

func NewGetRequest[DataType any]() *GetRequest[DataType] {
	return &GetRequest[DataType]{
		ResChan: make(chan DataType, 100),
	}
}

type SetRequest[DataType any] struct {
	Data    DataType
	ResChan chan bool
}

func NewSetRequest[DataType any](data DataType) *SetRequest[DataType] {
	return &SetRequest[DataType]{
		Data:    data,
		ResChan: make(chan bool, 100),
	}
}

type GetSetRequest[GetDataType any, SetDataType any] struct {
	DataToSet SetDataType
	ResChan   chan GetDataType
}

func NewGetSetRequest[GetDataType any, SetDataType any](dataToSet SetDataType) *GetSetRequest[GetDataType, SetDataType] {
	return &GetSetRequest[GetDataType, SetDataType]{
		DataToSet: dataToSet,
		ResChan:   make(chan GetDataType, 100),
	}
}
