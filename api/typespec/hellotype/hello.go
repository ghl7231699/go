package hellotype

type GreeterRequest struct {
	Name string `form:"name"`
}

type GreeterResponse struct {
	Data string `json:"data"`
}
