package views

import productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"

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
type Dictionaries_by_cat struct {
	Brands     []*productsRPC.Brand    `json:"brands"`
	Categories []*productsRPC.Category `json:"categories"`
	Countries  []*productsRPC.Country  `json:"countries"`
	Materials  []*productsRPC.Material `json:"materials"`
	Colors     []*productsRPC.Color    `json:"colors"`

	MinPrice  int32 `json:"min_price"`
	MaxPrice  int32 `json:"max_price"`
	MinWidth  int32 `json:"min_width"`
	MaxWidth  int32 `json:"max_width"`
	MinHeight int32 `json:"min_height"`
	MaxHeight int32 `json:"max_height"`
	MinDepth  int32 `json:"min_depth"`
	MaxDepth  int32 `json:"max_depth"`

	FilteredCategory string `json:"filtered_category"`
}
