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
type Email struct {
	RecipientEmail string `json:"recipient_email" example:"example@example.com"`
	RecipientName  string `json:"recipient_name" example:"Ivan Ivanov"`
	Subject        string `json:"subject" example:"Ivan Ivanov"`
	Text           string `json:"text" example:"This is the text content"`
	Html           string `json:"html" example:"<p>This is the HTML content</p>"`
}
