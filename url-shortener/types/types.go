package types

type ResultData struct {
	Link string
}

type ErrorData struct {
	Status      int
	Message     string
	Description string
}
