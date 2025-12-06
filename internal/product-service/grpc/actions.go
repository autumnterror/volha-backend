package grpc

import (
	"context"
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/internal/product-service/psql"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"
	"github.com/autumnterror/volha-backend/pkg/views"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerAPI) GetDictionaries(ctx context.Context, req *productsRPC.Id) (*productsRPC.Dictionaries, error) {
	const op = "grpc.GetDictionaries"
	log.Blue(op)

	d, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetDictionaries(ctx, req.GetId())
	})
	if err != nil {
		return nil, err
	}
	return d.(*productsRPC.Dictionaries), nil
}

// ============ Products ============

func (s *ServerAPI) SearchProducts(ctx context.Context, req *productsRPC.ProductSearch) (*productsRPC.ProductList, error) {
	const op = "grpc.SearchProducts"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.SearchProducts(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	products := res.([]*productsRPC.Product)
	return &productsRPC.ProductList{
		Items: products,
	}, nil
}

func (s *ServerAPI) FilterProducts(ctx context.Context, req *productsRPC.ProductFilter) (*productsRPC.ProductList, error) {
	const op = "grpc.FilterProducts"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.FilterProducts(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	products := res.([]*productsRPC.Product)
	return &productsRPC.ProductList{
		Items: products,
	}, nil
}

func (s *ServerAPI) GetAllProducts(ctx context.Context, req *productsRPC.Pagination) (*productsRPC.ProductList, error) {
	const op = "grpc.GetAllProducts"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAllProducts(ctx, int(req.GetStart()), int(req.GetFinish()))
	})
	if err != nil {
		return nil, err
	}

	products := res.([]*productsRPC.Product)
	return &productsRPC.ProductList{
		Items: products,
	}, nil
}

func (s *ServerAPI) GetProduct(ctx context.Context, req *productsRPC.Id) (*productsRPC.Product, error) {
	const op = "grpc.GetProduct"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetProduct(ctx, req.GetId())
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.Product), nil
}

func (s *ServerAPI) CreateProduct(ctx context.Context, req *productsRPC.ProductId) (*emptypb.Empty, error) {
	const op = "grpc.CreateProduct"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.CreateProduct(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateProduct(ctx context.Context, req *productsRPC.ProductId) (*emptypb.Empty, error) {
	const op = "grpc.UpdateProduct"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.UpdateProduct(ctx, req)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteProduct(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteProduct"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.DeleteProduct(ctx, req.GetId())
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ============ Brands ============

func (s *ServerAPI) CreateBrand(ctx context.Context, req *productsRPC.Brand) (*emptypb.Empty, error) {
	const op = "grpc.CreateBrand"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Create(ctx, req, views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateBrand(ctx context.Context, req *productsRPC.Brand) (*emptypb.Empty, error) {
	const op = "grpc.UpdateBrand"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Update(ctx, req, views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteBrand(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteBrand"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Delete(ctx, req.GetId(), views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllBrands(ctx context.Context, _ *emptypb.Empty) (*productsRPC.BrandList, error) {
	const op = "grpc.GetAllBrands"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAll(ctx, views.Brand)
	})
	if err != nil {
		return nil, err
	}

	brands := res.([]*productsRPC.Brand)
	return &productsRPC.BrandList{
		Items: brands,
	}, nil
}

func (s *ServerAPI) GetBrand(ctx context.Context, req *productsRPC.Id) (*productsRPC.Brand, error) {
	const op = "grpc.GetBrand"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.Get(ctx, req.GetId(), views.Brand)
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.Brand), nil
}

// ============ Categories ============

func (s *ServerAPI) CreateCategory(ctx context.Context, req *productsRPC.Category) (*emptypb.Empty, error) {
	const op = "grpc.CreateCategory"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Create(ctx, req, views.Category)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateCategory(ctx context.Context, req *productsRPC.Category) (*emptypb.Empty, error) {
	const op = "grpc.UpdateCategory"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Update(ctx, req, views.Category)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteCategory(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteCategory"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Delete(ctx, req.GetId(), views.Category)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllCategories(ctx context.Context, _ *emptypb.Empty) (*productsRPC.CategoryList, error) {
	const op = "grpc.GetAllCategories"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAll(ctx, views.Category)
	})
	if err != nil {
		return nil, err
	}

	categories := res.([]*productsRPC.Category)
	return &productsRPC.CategoryList{
		Items: categories,
	}, nil
}

func (s *ServerAPI) GetCategory(ctx context.Context, req *productsRPC.Id) (*productsRPC.Category, error) {
	const op = "grpc.GetCategory"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.Get(ctx, req.GetId(), views.Category)
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.Category), nil
}

// ============ Countries ============

func (s *ServerAPI) CreateCountry(ctx context.Context, req *productsRPC.Country) (*emptypb.Empty, error) {
	const op = "grpc.CreateCountry"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Create(ctx, req, views.Country)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateCountry(ctx context.Context, req *productsRPC.Country) (*emptypb.Empty, error) {
	const op = "grpc.UpdateCountry"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Update(ctx, req, views.Country)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteCountry(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteCountry"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Delete(ctx, req.GetId(), views.Country)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllCountries(ctx context.Context, _ *emptypb.Empty) (*productsRPC.CountryList, error) {
	const op = "grpc.GetAllCountries"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAll(ctx, views.Country)
	})
	if err != nil {
		return nil, err
	}

	countries := res.([]*productsRPC.Country)
	return &productsRPC.CountryList{
		Items: countries,
	}, nil
}

func (s *ServerAPI) GetCountry(ctx context.Context, req *productsRPC.Id) (*productsRPC.Country, error) {
	const op = "grpc.GetCountry"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.Get(ctx, req.GetId(), views.Country)
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.Country), nil
}

// ============ Materials ============

func (s *ServerAPI) CreateMaterial(ctx context.Context, req *productsRPC.Material) (*emptypb.Empty, error) {
	const op = "grpc.CreateMaterial"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Create(ctx, req, views.Material)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateMaterial(ctx context.Context, req *productsRPC.Material) (*emptypb.Empty, error) {
	const op = "grpc.UpdateMaterial"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Update(ctx, req, views.Material)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteMaterial(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteMaterial"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Delete(ctx, req.GetId(), views.Material)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllMaterials(ctx context.Context, _ *emptypb.Empty) (*productsRPC.MaterialList, error) {
	const op = "grpc.GetAllMaterials"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAll(ctx, views.Material)
	})
	if err != nil {
		return nil, err
	}

	materials := res.([]*productsRPC.Material)
	return &productsRPC.MaterialList{
		Items: materials,
	}, nil
}

func (s *ServerAPI) GetMaterial(ctx context.Context, req *productsRPC.Id) (*productsRPC.Material, error) {
	const op = "grpc.GetMaterial"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.Get(ctx, req.GetId(), views.Material)
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.Material), nil
}

// ============ Colors ============

func (s *ServerAPI) CreateColor(ctx context.Context, req *productsRPC.Color) (*emptypb.Empty, error) {
	const op = "grpc.CreateColor"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Create(ctx, req, views.Color)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateColor(ctx context.Context, req *productsRPC.Color) (*emptypb.Empty, error) {
	const op = "grpc.UpdateColor"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Update(ctx, req, views.Color)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteColor(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteColor"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Delete(ctx, req.GetId(), views.Color)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllColors(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ColorList, error) {
	const op = "grpc.GetAllColors"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAll(ctx, views.Color)
	})
	if err != nil {
		return nil, err
	}

	colors := res.([]*productsRPC.Color)
	return &productsRPC.ColorList{
		Items: colors,
	}, nil
}

func (s *ServerAPI) GetColor(ctx context.Context, req *productsRPC.Id) (*productsRPC.Color, error) {
	const op = "grpc.GetColor"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.Get(ctx, req.GetId(), views.Color)
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.Color), nil
}

// ============ Product Color Photos ============

func (s *ServerAPI) CreateProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotos) (*emptypb.Empty, error) {
	const op = "grpc.CreateProductColorPhotos"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Create(ctx, req, views.ProductColorPhotos)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotos) (*emptypb.Empty, error) {
	const op = "grpc.UpdateProductColorPhotos"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Update(ctx, req, views.ProductColorPhotos)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteProductColorPhotos(ctx context.Context, req *productsRPC.ProductColorPhotosId) (*emptypb.Empty, error) {
	const op = "grpc.DeleteProductColorPhotos"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.DeleteProductColorPhotos(ctx, req.GetProductId(), req.GetColorId())
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllProductColorPhotos(ctx context.Context, _ *emptypb.Empty) (*productsRPC.ProductColorPhotosList, error) {
	const op = "grpc.GetAllProductColorPhotos"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAll(ctx, views.ProductColorPhotos)
	})
	if err != nil {
		return nil, err
	}

	pcps := res.([]*productsRPC.ProductColorPhotos)
	return &productsRPC.ProductColorPhotosList{
		Items: pcps,
	}, nil
}

func (s *ServerAPI) GetPhotosByProductAndColor(ctx context.Context, req *productsRPC.ProductColorPhotosId) (*productsRPC.PhotoList, error) {
	const op = "grpc.GetPhotosByProductAndColor"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetProductColorPhotos(ctx, req.GetProductId(), req.GetColorId())
	})
	if err != nil {
		return nil, err
	}

	photos := res.(*productsRPC.ProductColorPhotos)
	return &productsRPC.PhotoList{
		Items: photos.GetPhotos(),
	}, nil
}

// ============ Slides ============

func (s *ServerAPI) CreateSlide(ctx context.Context, req *productsRPC.Slide) (*emptypb.Empty, error) {
	const op = "grpc.CreateSlide"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Create(ctx, req, views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) UpdateSlide(ctx context.Context, req *productsRPC.Slide) (*emptypb.Empty, error) {
	const op = "grpc.UpdateSlide"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Update(ctx, req, views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) DeleteSlide(ctx context.Context, req *productsRPC.Id) (*emptypb.Empty, error) {
	const op = "grpc.DeleteSlide"
	log.Blue(op)

	_, err := s.runTx(ctx, op, func(repo psql.Repo) (any, error) {
		return nil, repo.Delete(ctx, req.GetId(), views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ServerAPI) GetAllSlides(ctx context.Context, _ *emptypb.Empty) (*productsRPC.SlideList, error) {
	const op = "grpc.GetAllSlides"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.GetAll(ctx, views.Slide)
	})
	if err != nil {
		return nil, err
	}

	slides := res.([]*productsRPC.Slide)
	return &productsRPC.SlideList{
		Items: slides,
	}, nil
}

func (s *ServerAPI) GetSlide(ctx context.Context, req *productsRPC.Id) (*productsRPC.Slide, error) {
	const op = "grpc.GetSlide"
	log.Blue(op)

	res, err := s.run(ctx, op, func(repo psql.Repo) (any, error) {
		return repo.Get(ctx, req.GetId(), views.Slide)
	})
	if err != nil {
		return nil, err
	}

	return res.(*productsRPC.Slide), nil
}
