package base

type CarBase interface {
	SpeedUp(add int) int
	Park() bool
	Info() string
}

type CarLight struct {
	Brand string
	IsLed bool
	Color int
}

type CarDoor struct {
	Brand    string
	Color    int
	Meterial int
}
