package ports

// PORTS: interfaces to attach on external api's

type API interface {
	GetAddition(int32, int32) (int32, error)
	GetDivision(int32, int32) (int32, error)
	GetSubtraction(int32, int32) (int32, error)
	GetMultiplication(int32, int32) (int32, error)
}

type Db interface {
	Close()
	ToHistory(answer int32, operation string) error
}
