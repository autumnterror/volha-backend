package psql

import (
	"productService/internal/views"
	"testing"
)

//func TestCRUD(t *testing.T) {
//	db, err := NewConnect(config.Test())
//	assert.NoError(t, err)
//
//	tx, err := db.ProductRepo.Begin()
//	assert.NoError(t, err)
//
//	productOp := ProductRepo{ProductRepo: tx}
//
//	t.Run("create", func(t *testing.T) {
//		err = productOp.Create(views.Product{
//			Id:        "test",
//			Sizes:     "test",
//			Materials: "test",
//			Colors:    "test",
//			Photos: []string{
//				"test",
//				"test",
//			},
//			Seems: []string{
//				"test",
//				"test",
//			},
//			Price: 1488,
//		})
//		assert.NoError(t, err)
//	})
//	t.Run("create2", func(t *testing.T) {
//		err = productOp.Create(views.Product{
//			Id:        "test other",
//			Sizes:     "test other",
//			Materials: "test other",
//			Colors:    "test other",
//			Photos: []string{
//				"test other",
//				"test other",
//			},
//			Seems: []string{
//				"test other",
//				"test other",
//			},
//			Price: 1488,
//		})
//		assert.NoError(t, err)
//	})
//
//	t.Run("get all", func(t *testing.T) {
//		p, err := productOp.GetAll()
//		log.Println(p)
//
//		assert.NoError(t, err)
//		assert.NotEqual(t, len(p), 0)
//	})
//
//	t.Run("update", func(t *testing.T) {
//		assert.NoError(t, productOp.Update(views.Product{
//			//Id:        "test",
//			Sizes:     "test2",
//			Materials: "test2",
//			Colors:    "test2",
//			Photos: []string{
//				"test2",
//				"test2",
//			},
//			Seems: []string{
//				"test2",
//				"test2",
//			},
//			Price: 1488,
//		}, "test"))
//	})
//
//	t.Run("get all2", func(t *testing.T) {
//		p, err := productOp.GetAll()
//		log.Println(p)
//
//		assert.NoError(t, err)
//		assert.NotEqual(t, len(p), 0)
//	})
//
//	t.Run("delete", func(t *testing.T) {
//		assert.NoError(t, productOp.Delete("test"))
//	})
//
//	t.Cleanup(func() {
//		assert.NoError(t, tx.Rollback())
//		assert.NoError(t, db.Disconnect())
//	})
//}

type Test struct {
	f     views.ProductFilter
	title string
}

func TestFilter(t *testing.T) {
	tts := []Test{
		{
			f: views.ProductFilter{
				Brand:     []string{"penis", "pisos"},
				Country:   []string{"pro", "denis"},
				SortBy:    "price",
				SortOrder: "desc",
			},
			title: "Фильтр по бренду и стране с сортировкой по цене (по убыванию)",
		},
		{
			f: views.ProductFilter{
				MinWidth:  50,
				MaxWidth:  100,
				MinHeight: 70,
				MaxHeight: 120,
				Materials: []string{"wood", "metal"},
			},
			title: "Фильтр по размерам (диапазон) и материалу",
		},

		{
			f: views.ProductFilter{
				Colors:   []string{"black", "white"},
				MinPrice: 1000,
				MaxPrice: 5000,
				Limit:    10,
				Offset:   20,
			},
			title: "Фильтр по цвету и цене (диапазон) с пагинацией",
		},

		{
			f: views.ProductFilter{
				MinDepth:  30,
				MaxDepth:  60,
				SortBy:    "height",
				SortOrder: "asc",
			},
			title: "Фильтр по швам и глубине с сортировкой по высоте",
		},
		{
			f: views.ProductFilter{
				Brand:     []string{"ADIDAS", "NIKE"},
				Country:   []string{"Ribakovia", "Moldova"},
				MinWidth:  40,
				MaxWidth:  80,
				MinHeight: 60,
				MaxHeight: 90,
				MinDepth:  20,
				MaxDepth:  50,
				Materials: []string{"cotton", "polyester"},
				Colors:    []string{"blue", "green"},
				MinPrice:  500,
				MaxPrice:  2000,
				SortBy:    "price",
				SortOrder: "asc",
				Limit:     15,
				Offset:    10,
			},
			title: "Комплексный фильтр (все параметры)",
		},
		{
			f: views.ProductFilter{
				Limit: 5,
			},
			title: "Минимальный фильтр (только лимит)",
		},
		{
			f: views.ProductFilter{
				Materials: []string{"leather"},
				Colors:    []string{"brown"},
			},
			title: "Фильтр по одному материалу и цвету",
		},
	}

	for _, tt := range tts {
		t.Run(tt.title, func(t *testing.T) {
			//FilterProducts(tt.f)
		})
	}
}
