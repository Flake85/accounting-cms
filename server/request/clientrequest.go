package request

type ClientRequest struct {
	Name	string `json:"name"`
	Address string `json:"address"`
	Email	string `json:"email"`
}