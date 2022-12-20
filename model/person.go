package model

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}
