package redis

import (
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/api/proto/gen"
	"github.com/autumnterror/volha-backend/internal/gateway/config"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestDictionariesGood(t *testing.T) {
	c := MustNew(config.Test())

	td := &productsRPC.productsRPC{
		Brands:     []*productsRPC.Brand{},
		Categories: []*productsRPC.Category{},
		Countries:  []*productsRPC.Country{},
		Materials:  []*productsRPC.Material{},
		Colors:     []*productsRPC.Color{},
		MinPrice:   1,
		MaxPrice:   2,
		MinWidth:   3,
		MaxWidth:   4,
		MinHeight:  5,
		MaxHeight:  6,
		MinDepth:   7,
		MaxDepth:   8,
	}

	assert.NoError(t, c.SetDictionaries(td))

	d, err := c.GetDictionaries()
	assert.NoError(t, err)
	assert.NotNil(t, d)
	log.Println(d)

	assert.NoError(t, c.CleanDictionaries())

	d, err = c.GetDictionaries()
	assert.Error(t, err)
	assert.Nil(t, d)
}

func TestDictionariesNoCache(t *testing.T) {
	c := MustNew(config.Test())

	assert.NoError(t, c.CleanDictionaries())

	d, err := c.GetDictionaries()
	assert.Error(t, err)
	assert.Nil(t, d)
}

func TestDictionariesByCategoryGood(t *testing.T) {
	c := MustNew(config.Test())

	id := "category-1"

	td := &productsRPC.Dictionaries{
		Brands:     []*productsRPC.Brand{},
		Categories: []*productsRPC.Category{},
		Countries:  []*productsRPC.Country{},
		Materials:  []*productsRPC.Material{},
		Colors:     []*productsRPC.Color{},
		MinPrice:   1,
		MaxPrice:   2,
		MinWidth:   3,
		MaxWidth:   4,
		MinHeight:  5,
		MaxHeight:  6,
		MinDepth:   7,
		MaxDepth:   8,
	}

	assert.NoError(t, c.SetDictionariesByCategory(id, td))

	d, err := c.GetDictionariesByCategory(id)
	assert.NoError(t, err)
	assert.NotNil(t, d)
	log.Println(d)

	assert.Equal(t, td.MinPrice, d.MinPrice)
	assert.Equal(t, td.MaxPrice, d.MaxPrice)
	assert.Equal(t, td.MinWidth, d.MinWidth)
	assert.Equal(t, td.MaxWidth, d.MaxWidth)
	assert.Equal(t, td.MinHeight, d.MinHeight)
	assert.Equal(t, td.MaxHeight, d.MaxHeight)
	assert.Equal(t, td.MinDepth, d.MinDepth)
	assert.Equal(t, td.MaxDepth, d.MaxDepth)

	assert.NoError(t, c.CleanDictionaries())

	d, err = c.GetDictionariesByCategory(id)
	assert.Error(t, err)
	assert.Nil(t, d)
}

func TestDictionariesByCategoryNoCache(t *testing.T) {
	c := MustNew(config.Test())

	id := "category-2"

	assert.NoError(t, c.CleanDictionaries())

	d, err := c.GetDictionariesByCategory(id)
	assert.Error(t, err)
	assert.Nil(t, d)
}
