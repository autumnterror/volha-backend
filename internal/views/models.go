package views

type Product struct {
	Id          string
	Title       string   `json:"title"`
	Article     string   `json:"article"`
	Brand       string   `json:"brand"`
	Country     string   `json:"country"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Depth       int      `json:"depth"`
	Materials   []string `json:"materials"`
	Colors      []string `json:"colors"`
	Photos      []string `json:"photos"`
	Seems       []string `json:"seems"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
}

type ProductFilter struct {
	Brand     []string `json:"brand"`
	Country   []string `json:"country"`
	MinWidth  int      `json:"min_width"`
	MaxWidth  int      `json:"max_width"`
	MinHeight int      `json:"min_height"`
	MaxHeight int      `json:"max_height"`
	MinDepth  int      `json:"min_depth"`
	MaxDepth  int      `json:"max_depth"`
	Materials []string `json:"materials"`
	Colors    []string `json:"colors"`
	MinPrice  int      `json:"min_price"`
	MaxPrice  int      `json:"max_price"`
	SortBy    string   `json:"sort_by"`
	SortOrder string   `json:"sort_order"`
	Offset    int      `json:"offset"`
	Limit     int      `json:"limit"`
}
