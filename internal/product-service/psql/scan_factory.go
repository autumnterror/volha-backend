package psql

import (
	"database/sql"

	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/lib/pq"
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

type PCPScanner struct {
	list []*productsRPC.ProductColorPhotos
}

func (s *PCPScanner) Scan(rows *sql.Rows) error {
	var pcp productsRPC.ProductColorPhotos
	if err := rows.Scan(&pcp.ProductId, &pcp.ColorId, pq.Array(&pcp.Photos)); err != nil {
		return err
	}
	s.list = append(s.list, &pcp)
	return nil
}

func (s *PCPScanner) GetList() any {
	return s.list
}

type SlideScanner struct {
	list []*productsRPC.Slide
}

func (s *SlideScanner) Scan(rows *sql.Rows) error {
	var sl productsRPC.Slide
	if err := rows.Scan(&sl.Id, &sl.Link, &sl.Img, &sl.Img762); err != nil {
		return err
	}
	s.list = append(s.list, &sl)
	return nil
}

func (s *SlideScanner) GetList() any {
	return s.list
}

//-----------------------------------ROW-----------------------------------

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

type PCPScannerRow struct {
	pcp *productsRPC.ProductColorPhotos
}

func (s *PCPScannerRow) Scan(rows *sql.Row) error {
	var pcp productsRPC.ProductColorPhotos
	if err := rows.Scan(&pcp.ProductId, &pcp.ColorId, pq.Array(&pcp.Photos)); err != nil {
		return err
	}
	s.pcp = &pcp
	return nil
}

func (s *PCPScannerRow) Get() any {
	return s.pcp
}

type SlideScannerRow struct {
	sl *productsRPC.Slide
}

func (s *SlideScannerRow) Scan(rows *sql.Row) error {
	var sl productsRPC.Slide
	if err := rows.Scan(&sl.Id, &sl.Link, &sl.Img, &sl.Img762); err != nil {
		return err
	}
	s.sl = &sl
	return nil
}

func (s *SlideScannerRow) Get() any {
	return s.sl
}
