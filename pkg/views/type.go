package views

type Type int

const (
	Brand Type = iota
	Category
	Color
	Material
	Country
	ProductColorPhotos
	Slide
	Product
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
	case ProductColorPhotos:
		return "ProductColorPhotos"
	case Slide:
		return "Slide"
	case Product:
		return "Product"
	default:
		return "Unknown"
	}
}
