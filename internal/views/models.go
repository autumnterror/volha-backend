package views

type Product struct {
	Id string

	Title   string
	Article string
	Brand   string
	Country string

	Width  int
	Height int
	Depth  int

	Materials []string
	Colors    []string
	Photos    []string
	Seems     []string

	Price       int
	Description string
}

type ProductFilter struct {
	Brand     []string
	Country   []string
	MinWidth  int
	MaxWidth  int
	MinHeight int
	MaxHeight int
	MinDepth  int
	MaxDepth  int
	Materials []string
	Colors    []string
	MinPrice  int
	MaxPrice  int
	SortBy    string
	SortOrder string
	Offset    int
	Limit     int
}
