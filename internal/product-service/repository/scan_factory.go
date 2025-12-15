package repository

import (
	"database/sql"

	"github.com/autumnterror/volha-backend/internal/product-service/domain"
)

type entityScanner interface {
	Scan(rows *sql.Rows) error
	GetList() any
}

type brandScanner struct {
	list []*domain.Brand
}

func (s *brandScanner) Scan(rows *sql.Rows) error {
	var b domain.Brand
	if err := rows.Scan(&b.Id, &b.Title); err != nil {
		return err
	}
	s.list = append(s.list, &b)
	return nil
}

func (s *brandScanner) GetList() any {
	return s.list
}

type countryScanner struct {
	list []*domain.Country
}

func (s *countryScanner) Scan(rows *sql.Rows) error {
	var c domain.Country
	if err := rows.Scan(&c.Id, &c.Title, &c.Friendly); err != nil {
		return err
	}
	s.list = append(s.list, &c)
	return nil
}

func (s *countryScanner) GetList() any {
	return s.list
}

type colorScanner struct {
	list []*domain.Color
}

func (s *colorScanner) Scan(rows *sql.Rows) error {
	var c domain.Color
	if err := rows.Scan(&c.Id, &c.Title, &c.Hex); err != nil {
		return err
	}
	s.list = append(s.list, &c)
	return nil
}

func (s *colorScanner) GetList() any {
	return s.list
}

type materialScanner struct {
	list []*domain.Material
}

func (s *materialScanner) Scan(rows *sql.Rows) error {
	var m domain.Material
	if err := rows.Scan(&m.Id, &m.Title); err != nil {
		return err
	}
	s.list = append(s.list, &m)
	return nil
}

func (s *materialScanner) GetList() any {
	return s.list
}

type categoryScanner struct {
	list []*domain.Category
}

func (s *categoryScanner) Scan(rows *sql.Rows) error {
	var c domain.Category
	if err := rows.Scan(&c.Id, &c.Title, &c.Uri, &c.Img); err != nil {
		return err
	}
	s.list = append(s.list, &c)
	return nil
}

func (s *categoryScanner) GetList() any {
	return s.list
}

type slideScanner struct {
	list []*domain.Slide
}

func (s *slideScanner) Scan(rows *sql.Rows) error {
	var sl domain.Slide
	if err := rows.Scan(&sl.Id, &sl.Link, &sl.Img, &sl.Img762); err != nil {
		return err
	}
	s.list = append(s.list, &sl)
	return nil
}

func (s *slideScanner) GetList() any {
	return s.list
}

//-----------------------------------ROW-----------------------------------

type entityScannerRow interface {
	Scan(rows *sql.Row) error
	Get() any
}

type brandScannerRow struct {
	b *domain.Brand
}

func (s *brandScannerRow) Scan(rows *sql.Row) error {
	var b domain.Brand
	if err := rows.Scan(&b.Id, &b.Title); err != nil {
		return err
	}
	s.b = &b
	return nil
}

func (s *brandScannerRow) Get() any {
	return s.b
}

type countryScannerRow struct {
	c *domain.Country
}

func (s *countryScannerRow) Scan(rows *sql.Row) error {
	var c domain.Country
	if err := rows.Scan(&c.Id, &c.Title, &c.Friendly); err != nil {
		return err
	}
	s.c = &c
	return nil
}

func (s *countryScannerRow) Get() any {
	return s.c
}

type colorScannerRow struct {
	c *domain.Color
}

func (s *colorScannerRow) Scan(rows *sql.Row) error {
	var c domain.Color
	if err := rows.Scan(&c.Id, &c.Title, &c.Hex); err != nil {
		return err
	}
	s.c = &c
	return nil
}

func (s *colorScannerRow) Get() any {
	return s.c
}

type materialScannerRow struct {
	m *domain.Material
}

func (s *materialScannerRow) Scan(rows *sql.Row) error {
	var m domain.Material
	if err := rows.Scan(&m.Id, &m.Title); err != nil {
		return err
	}
	s.m = &m
	return nil
}

func (s *materialScannerRow) Get() any {
	return s.m
}

type categoryScannerRow struct {
	c *domain.Category
}

func (s *categoryScannerRow) Scan(rows *sql.Row) error {
	var c domain.Category
	if err := rows.Scan(&c.Id, &c.Title, &c.Uri, &c.Img); err != nil {
		return err
	}
	s.c = &c
	return nil
}

func (s *categoryScannerRow) Get() any {
	return s.c
}

type slideScannerRow struct {
	sl *domain.Slide
}

func (s *slideScannerRow) Scan(rows *sql.Row) error {
	var sl domain.Slide
	if err := rows.Scan(&sl.Id, &sl.Link, &sl.Img, &sl.Img762); err != nil {
		return err
	}
	s.sl = &sl
	return nil
}

func (s *slideScannerRow) Get() any {
	return s.sl
}
