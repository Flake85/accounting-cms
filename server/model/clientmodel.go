package model

type Client struct {
	BaseModel
	Name	string	`json:"name"`
	Address	string	`json:"address"`
	Email	string	`json:"email"`
}