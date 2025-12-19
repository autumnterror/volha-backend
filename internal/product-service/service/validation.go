package service

import (
	"errors"

	"github.com/autumnterror/volha-backend/internal/product-service/domain"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/rs/xid"
)

func validateID(id string) error {
	if id == "" {
		return errors.New("id is empty")
	}
	if _, err := xid.FromString(id); err != nil {
		return errors.New("id not in xid")
	}
	return nil
}

func validateBasicType(_type views.Type) error {
	switch _type {
	case views.Brand, views.Category, views.Color, views.Material, views.Country, views.Slide, views.Article:
		return nil
	default:
		return domain.ErrUnknownType
	}
}

func validateBasicPayload(obj any, _type views.Type) error {
	if obj == nil {
		return errors.New("object is nil")
	}
	switch _type {
	case views.Brand:
		b, ok := obj.(*domain.Brand)
		if !ok {
			return domain.ErrInvalidType
		}
		if err := validateID(b.Id); err != nil {
			return err
		}
		if b.Title == "" {
			return errors.New("brand title is empty")
		}
	case views.Category:
		c, ok := obj.(*domain.Category)
		if !ok {
			return domain.ErrInvalidType
		}
		if err := validateID(c.Id); err != nil {
			return err
		}
		if c.Title == "" {
			return errors.New("category title is empty")
		}
		if c.Uri == "" {
			return errors.New("category uri is empty")
		}
	case views.Country:
		c, ok := obj.(*domain.Country)
		if !ok {
			return domain.ErrInvalidType
		}
		if err := validateID(c.Id); err != nil {
			return err
		}
		if c.Title == "" {
			return errors.New("country title is empty")
		}
	case views.Material:
		m, ok := obj.(*domain.Material)
		if !ok {
			return domain.ErrInvalidType
		}
		if err := validateID(m.Id); err != nil {
			return err
		}
		if m.Title == "" {
			return errors.New("material title is empty")
		}
	case views.Color:
		clr, ok := obj.(*domain.Color)
		if !ok {
			return domain.ErrInvalidType
		}
		if err := validateID(clr.Id); err != nil {
			return err
		}
		if clr.Title == "" {
			return errors.New("color title is empty")
		}
		if clr.Hex == "" {
			return errors.New("color hex is empty")
		}
	case views.Slide:
		sl, ok := obj.(*domain.Slide)
		if !ok {
			return domain.ErrInvalidType
		}
		if err := validateID(sl.Id); err != nil {
			return err
		}
		if sl.Link == "" {
			return errors.New("slide link is empty")
		}
		if sl.Img == "" {
			return errors.New("slide img is empty")
		}
		if sl.Img762 == "" {
			return errors.New("slide img762 is empty")
		}
	case views.Article:
		sl, ok := obj.(*domain.Article)
		if !ok {
			return domain.ErrInvalidType
		}
		if err := validateID(sl.Id); err != nil {
			return err
		}
		if sl.Title == "" {
			return errors.New("article title is empty")
		}
		if sl.Img == "" {
			return errors.New("article img is empty")
		}
	default:
		return domain.ErrUnknownType
	}

	return nil
}

func validatePCPIDs(productID, colorID string) error {
	if err := validateID(productID); err != nil {
		return err
	}
	if err := validateID(colorID); err != nil {
		return err
	}
	return nil
}

func validatePCP(pcp *domain.ProductColorPhotos) error {
	if pcp == nil {
		return errors.New("object is nil")
	}
	if err := validatePCPIDs(pcp.ProductId, pcp.ColorId); err != nil {
		return err
	}
	if len(pcp.Photos) == 0 {
		return errors.New("photos are empty")
	}
	return nil
}

func validateProductPayload(p *domain.ProductId) error {
	if p == nil {
		return errors.New("object is nil")
	}
	if err := validateID(p.Id); err != nil {
		return err
	}
	if p.Title == "" {
		return errors.New("product title is empty")
	}
	if p.Article == "" {
		return errors.New("product article is empty")
	}
	if err := validateID(p.Brand); err != nil {
		return errors.New("brand id is invalid")
	}
	if err := validateID(p.Category); err != nil {
		return errors.New("category id is invalid")
	}
	if err := validateID(p.Country); err != nil {
		return errors.New("country id is invalid")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if p.Width <= 0 || p.Height <= 0 || p.Depth <= 0 {
		return errors.New("dimensions must be greater than zero")
	}
	if len(p.Photos) == 0 {
		return errors.New("photos are empty")
	}
	for _, s := range p.Seems {
		if err := validateID(s); err != nil {
			return errors.New("seem id is invalid")
		}
	}

	return nil
}

func validateProductRange(start, end int) error {
	if start < 0 || end <= start {
		return errors.New("range boundaries are invalid")
	}
	return nil
}

func validateSearch(filter *domain.ProductSearch) error {
	if filter == nil {
		return errors.New("filter is nil")
	}
	if filter.Id == "" && filter.Title == "" && filter.Article == "" {
		return errors.New("search criteria is empty")
	}
	return nil
}

func validateFilter(filter *domain.ProductFilter) error {
	if filter == nil {
		return errors.New("filter is nil")
	}
	if filter.Limit < 0 || filter.Offset < 0 {
		return errors.New("limit or offset cannot be negative")
	}
	return nil
}
