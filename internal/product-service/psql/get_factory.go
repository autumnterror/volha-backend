package psql

import (
	"database/sql"

	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
)

type EntityScannerRow interface {
	Scan(rows *sql.Row) error
	Get() any
}

type BrandScannerRow struct {
	b *productsRPC.Brand
}

func (s *BrandScannerRow) Scan(rows *sql.Row) error {
	var b productsRPC.Brand
	if err := rows.Scan(&b.Id, &b.Title); err != nil {
		return err
	}
	s.b = &b
	return nil
}

func (s *BrandScannerRow) Get() any {
	return s.b
}

type CountryScannerRow struct {
	c *productsRPC.Country
}

func (s *CountryScannerRow) Scan(rows *sql.Row) error {
	var c productsRPC.Country
	if err := rows.Scan(&c.Id, &c.Title, &c.Friendly); err != nil {
		return err
	}
	s.c = &c
	return nil
}

func (s *CountryScannerRow) Get() any {
	return s.c
}

type ColorScannerRow struct {
	c *productsRPC.Color
}

func (s *ColorScannerRow) Scan(rows *sql.Row) error {
	var c productsRPC.Color
	if err := rows.Scan(&c.Id, &c.Title, &c.Hex); err != nil {
		return err
	}
	s.c = &c
	return nil
}

func (s *ColorScannerRow) Get() any {
	return s.c
}

type MaterialScannerRow struct {
	m *productsRPC.Material
}

func (s *MaterialScannerRow) Scan(rows *sql.Row) error {
	var m productsRPC.Material
	if err := rows.Scan(&m.Id, &m.Title); err != nil {
		return err
	}
	s.m = &m
	return nil
}

func (s *MaterialScannerRow) Get() any {
	return s.m
}

type CategoryScannerRow struct {
	c *productsRPC.Category
}

func (s *CategoryScannerRow) Scan(rows *sql.Row) error {
	var c productsRPC.Category
	if err := rows.Scan(&c.Id, &c.Title, &c.Uri, &c.Img); err != nil {
		return err
	}
	s.c = &c
	return nil
}

func (s *CategoryScannerRow) Get() any {
	return s.c
}
