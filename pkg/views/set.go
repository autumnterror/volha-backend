package views

type Type int

const (
	Brand Type = iota
	Category
	Color
	Material
	Country
)

func (t Type) String() string {
	switch t {
	case Brand:
		return "Brand"
	case Category:
		return "Category"
	case Color:
		return "Color"
	case Material:
		return "Material"
	case Country:
		return "Country"
	default:
		return "Unknown"
	}
}
