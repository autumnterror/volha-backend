package psql

import (
	"database/sql"

	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
)

type EntityScanner interface {
	Scan(rows *sql.Rows) error
	GetList() any
}

type BrandScanner struct {
	list []*productsRPC.Brand
}

func (s *BrandScanner) Scan(rows *sql.Rows) error {
	var b productsRPC.Brand
	if err := rows.Scan(&b.Id, &b.Title); err != nil {
		return err
	}
	s.list = append(s.list, &b)
	return nil
}

func (s *BrandScanner) GetList() any {
	return s.list
}

type CountryScanner struct {
	list []*productsRPC.Country
}

func (s *CountryScanner) Scan(rows *sql.Rows) error {
	var c productsRPC.Country
	if err := rows.Scan(&c.Id, &c.Title, &c.Friendly); err != nil {
		return err
	}
	s.list = append(s.list, &c)
	return nil
}

func (s *CountryScanner) GetList() any {
	return s.list
}

type ColorScanner struct {
	list []*productsRPC.Color
}

func (s *ColorScanner) Scan(rows *sql.Rows) error {
	var c productsRPC.Color
	if err := rows.Scan(&c.Id, &c.Title, &c.Hex); err != nil {
		return err
	}
	s.list = append(s.list, &c)
	return nil
}

func (s *ColorScanner) GetList() any {
	return s.list
}

type MaterialScanner struct {
	list []*productsRPC.Material
}

func (s *MaterialScanner) Scan(rows *sql.Rows) error {
	var m productsRPC.Material
	if err := rows.Scan(&m.Id, &m.Title); err != nil {
		return err
	}
	s.list = append(s.list, &m)
	return nil
}

func (s *MaterialScanner) GetList() any {
	return s.list
}

type CategoryScanner struct {
	list []*productsRPC.Category
}

func (s *CategoryScanner) Scan(rows *sql.Rows) error {
	var c productsRPC.Category
	if err := rows.Scan(&c.Id, &c.Title, &c.Uri, &c.Img); err != nil {
		return err
	}
	s.list = append(s.list, &c)
	return nil
}

func (s *CategoryScanner) GetList() any {
	return s.list
}
