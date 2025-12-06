package views

type SWGError struct {
	Error string `json:"error" example:"error"`
}
type SWGMessage struct {
	Message string `json:"message" example:"some info"`
}
type SWGId struct {
	Id string `json:"id" example:"asidofadsklhf"`
}
type SWGFileUploadResponse struct {
	Name string `json:"name" example:"example.jpg"`
}
