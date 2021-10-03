package base

type Response struct {
	Code    int
	Status  string
	Message string
	Data    interface{}
}
