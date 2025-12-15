package entities

type Cat struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Age    uint    `json:"age"`
	Breed  string  `json:"breed"`
	Weight float32 `json:"weight"`
}
