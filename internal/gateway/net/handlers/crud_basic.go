package handlers

import (
	"context"

	productsRPC "github.com/autumnterror/volha-backend/api/proto/gen"

	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/breezynotes/pkg/utils/format"
	"github.com/autumnterror/volha-backend/pkg/views"
	"github.com/rs/xid"

	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Apis) getAll(c echo.Context, op string, _type views.Type) (any, int, views.SWGError) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	var res any
	var err error

	switch _type {
	case views.Brand:
		l, nErr := a.apiProduct.API.GetAllBrands(ctx, nil)
		if len(l.GetItems()) == 0 {
			res = []*productsRPC.Brand{}
		} else {
			res = l.GetItems()
		}
		err = nErr
	case views.Country:
		l, nErr := a.apiProduct.API.GetAllCountries(ctx, nil)
		if len(l.GetItems()) == 0 {
			res = []*productsRPC.Country{}
		} else {
			res = l.GetItems()
		}
		err = nErr
	case views.Material:
		l, nErr := a.apiProduct.API.GetAllMaterials(ctx, nil)
		if len(l.GetItems()) == 0 {
			res = []*productsRPC.Material{}
		} else {
			res = l.GetItems()
		}
		err = nErr
	case views.Color:
		l, nErr := a.apiProduct.API.GetAllColors(ctx, nil)
		if len(l.GetItems()) == 0 {
			res = []*productsRPC.Color{}
		} else {
			res = l.GetItems()
		}
		err = nErr
	case views.Category:
		l, nErr := a.apiProduct.API.GetAllCategories(ctx, nil)
		if len(l.GetItems()) == 0 {
			res = []*productsRPC.Category{}
		} else {
			res = l.GetItems()
		}
		err = nErr
	case views.ProductColorPhotos:
		l, nErr := a.apiProduct.API.GetAllProductColorPhotos(ctx, &emptypb.Empty{})
		if len(l.GetItems()) == 0 {
			res = []*productsRPC.ProductColorPhotos{}
		} else {
			res = l.GetItems()
		}
		err = nErr
	case views.Slide:
		l, nErr := a.apiProduct.API.GetAllSlides(ctx, nil)
		if len(l.GetItems()) == 0 {
			res = []*productsRPC.Slide{}
		} else {
			res = l.GetItems()
		}
		err = nErr
	default:
		return nil, http.StatusInternalServerError, views.SWGError{Error: "bad type on gateway"}
	}

	if err != nil {
		code, swg := a.mapGRPCError(op, err)
		return nil, code, swg
	}
	return res, http.StatusOK, views.SWGError{}
}

func (a *Apis) get(c echo.Context, op string, _type views.Type) (any, int, views.SWGError) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	var res any
	var err error

	id := c.QueryParam("id")
	if id == "" {
		return nil, http.StatusBadRequest, views.SWGError{Error: "missing id"}
	}

	switch _type {
	case views.Brand:
		res, err = a.apiProduct.API.GetBrand(ctx, &productsRPC.Id{Id: id})
	case views.Country:
		res, err = a.apiProduct.API.GetCountry(ctx, &productsRPC.Id{Id: id})
	case views.Material:
		res, err = a.apiProduct.API.GetMaterial(ctx, &productsRPC.Id{Id: id})
	case views.Color:
		res, err = a.apiProduct.API.GetColor(ctx, &productsRPC.Id{Id: id})
	case views.Category:
		res, err = a.apiProduct.API.GetCategory(ctx, &productsRPC.Id{Id: id})
	case views.Slide:
		res, err = a.apiProduct.API.GetSlide(ctx, &productsRPC.Id{Id: id})
	case views.Product:
		res, err = a.apiProduct.API.GetProduct(ctx, &productsRPC.Id{Id: id})
	default:
		return nil, http.StatusInternalServerError, views.SWGError{Error: "bad type on gateway"}
	}

	if err != nil {
		code, swg := a.mapGRPCError(op, err)
		return nil, code, swg
	}
	return res, http.StatusOK, views.SWGError{}
}

func (a *Apis) create(c echo.Context, op string, _type views.Type) (any, int, views.SWGError) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()
	_ = a.rds.CleanDictionaries()

	switch _type {
	case views.Brand:
		var b productsRPC.Brand
		if err := c.Bind(&b); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		b.Id = xid.New().String()
		if _, err := a.apiProduct.API.CreateBrand(ctx, &b); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGId{Id: b.Id}, http.StatusOK, views.SWGError{}
	case views.Country:
		var ctr productsRPC.Country
		if err := c.Bind(&ctr); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		ctr.Id = xid.New().String()
		if _, err := a.apiProduct.API.CreateCountry(ctx, &ctr); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGId{Id: ctr.Id}, http.StatusOK, views.SWGError{}
	case views.Material:
		var m productsRPC.Material
		if err := c.Bind(&m); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		m.Id = xid.New().String()
		if _, err := a.apiProduct.API.CreateMaterial(ctx, &m); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGId{Id: m.Id}, http.StatusOK, views.SWGError{}
	case views.Color:
		var clr productsRPC.Color
		if err := c.Bind(&clr); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		clr.Id = xid.New().String()
		if _, err := a.apiProduct.API.CreateColor(ctx, &clr); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGId{Id: clr.Id}, http.StatusOK, views.SWGError{}
	case views.Category:
		var cat productsRPC.Category
		if err := c.Bind(&cat); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		cat.Id = xid.New().String()
		if _, err := a.apiProduct.API.CreateCategory(ctx, &cat); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGId{Id: cat.Id}, http.StatusOK, views.SWGError{}
	case views.Product:
		var p productsRPC.ProductId
		if err := c.Bind(&p); err != nil {
			log.Println(format.Error(op, err))
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if len(p.Article) != 8 {
			return nil, http.StatusBadRequest, views.SWGError{Error: "article must be 8 digits"}
		}

		p.Id = xid.New().String()
		if _, err := a.apiProduct.API.CreateProduct(ctx, &p); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGId{Id: p.Id}, http.StatusOK, views.SWGError{}
	case views.ProductColorPhotos:
		var pcp productsRPC.ProductColorPhotos
		if err := c.Bind(&pcp); err != nil {
			log.Println(format.Error(op, err))
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.CreateProductColorPhotos(ctx, &pcp); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "product color photos created successfully"}, http.StatusOK, views.SWGError{}
	case views.Slide:
		var s productsRPC.Slide
		if err := c.Bind(&s); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		s.Id = xid.New().String()
		if _, err := a.apiProduct.API.CreateSlide(ctx, &s); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGId{Id: s.Id}, http.StatusOK, views.SWGError{}
	default:
		return nil, http.StatusInternalServerError, views.SWGError{Error: "bad type on gateway"}
	}
}

func (a *Apis) update(c echo.Context, op string, _type views.Type) (any, int, views.SWGError) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	_ = a.rds.CleanDictionaries()

	switch _type {
	case views.Brand:
		var b productsRPC.Brand
		if err := c.Bind(&b); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateBrand(ctx, &b); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "brand updated"}, http.StatusOK, views.SWGError{}
	case views.Country:
		var ctr productsRPC.Country
		if err := c.Bind(&ctr); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateCountry(ctx, &ctr); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "country updated successfully"}, http.StatusOK, views.SWGError{}
	case views.Material:
		var m productsRPC.Material
		if err := c.Bind(&m); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateMaterial(ctx, &m); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "material updated successfully"}, http.StatusOK, views.SWGError{}
	case views.Color:
		var clr productsRPC.Color
		if err := c.Bind(&clr); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateColor(ctx, &clr); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "color updated successfully"}, http.StatusOK, views.SWGError{}
	case views.Category:
		var cat productsRPC.Category
		if err := c.Bind(&cat); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateCategory(ctx, &cat); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "category updated successfully"}, http.StatusOK, views.SWGError{}
	case views.Product:
		var p productsRPC.ProductId
		if err := c.Bind(&p); err != nil {
			log.Println(format.Error(op, err))
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateProduct(ctx, &p); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "product updated successfully"}, http.StatusOK, views.SWGError{}
	case views.ProductColorPhotos:
		var pcp productsRPC.ProductColorPhotos
		if err := c.Bind(&pcp); err != nil {
			log.Println(format.Error(op, err))
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateProductColorPhotos(ctx, &pcp); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "product color photos updated successfully"}, http.StatusOK, views.SWGError{}
	case views.Slide:
		var s productsRPC.Slide
		if err := c.Bind(&s); err != nil {
			return nil, http.StatusBadRequest, views.SWGError{Error: "bad JSON"}
		}
		if _, err := a.apiProduct.API.UpdateSlide(ctx, &s); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "slide updated successfully"}, http.StatusOK, views.SWGError{}
	default:
		return nil, http.StatusInternalServerError, views.SWGError{Error: "bad type on gateway"}
	}
}

func (a *Apis) delete(c echo.Context, op string, _type views.Type) (any, int, views.SWGError) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()

	id := c.QueryParam("id")
	if id == "" {
		return nil, http.StatusBadRequest, views.SWGError{Error: "missing id"}
	}
	_ = a.rds.CleanDictionaries()

	switch _type {
	case views.Brand:
		if _, err := a.apiProduct.API.DeleteBrand(ctx, &productsRPC.Id{Id: id}); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "brand deleted"}, http.StatusOK, views.SWGError{}
	case views.Country:
		if _, err := a.apiProduct.API.DeleteCountry(ctx, &productsRPC.Id{Id: id}); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "country deleted successfully"}, http.StatusOK, views.SWGError{}
	case views.Material:
		if _, err := a.apiProduct.API.DeleteMaterial(ctx, &productsRPC.Id{Id: id}); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "material deleted successfully"}, http.StatusOK, views.SWGError{}
	case views.Color:
		if _, err := a.apiProduct.API.DeleteColor(ctx, &productsRPC.Id{Id: id}); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "color deleted successfully"}, http.StatusOK, views.SWGError{}
	case views.Category:
		if _, err := a.apiProduct.API.DeleteCategory(ctx, &productsRPC.Id{Id: id}); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "category deleted successfully"}, http.StatusOK, views.SWGError{}
	case views.Product:
		if _, err := a.apiProduct.API.DeleteProduct(ctx, &productsRPC.Id{Id: id}); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "product deleted successfully"}, http.StatusOK, views.SWGError{}
	case views.Slide:
		if _, err := a.apiProduct.API.DeleteSlide(ctx, &productsRPC.Id{Id: id}); err != nil {
			code, swg := a.mapGRPCError(op, err)
			return nil, code, swg
		}
		return views.SWGMessage{Message: "slide deleted"}, http.StatusOK, views.SWGError{}
	default:
		return nil, http.StatusInternalServerError, views.SWGError{Error: "bad type on gateway"}
	}
}

//INTERNAL IN EVERY
//GetAll ErrUnknownType 	Unimplemented
//Get ErrUnknownType ErrNotFound 	Unimplemented
//GetPhotosByProductColor ErrNotFound 	NotFound
//Create ErrInvalidType ErrUnknownType ErrAlreadyExists ErrForeignKey 	InvalidArgument Unimplemented AlreadyExists FailedPrecondition
//Update ErrInvalidType ErrUnknownType ErrNotFound 		InvalidArgument Unimplemented NotFound
//Delete ErrUnknownType ErrNotFound 	Unimplemented NotFound
//DeleteProductColorPhotos ErrNotFound 		NotFound
//GetDictionaries
//FilterProducts
//SearchProducts
//Products all the def but without ErrInvalidType ErrUnknownType

func (a *Apis) mapGRPCError(op string, err error) (int, views.SWGError) {
	st, ok := status.FromError(err)
	if !ok {
		log.Error(op, "", err)
		return http.StatusBadGateway, views.SWGError{Error: "check logs on service"}
	}

	switch st.Code() {
	case codes.InvalidArgument:
		return http.StatusBadRequest, views.SWGError{Error: err.Error()}
	case codes.FailedPrecondition:
		return http.StatusBadRequest, views.SWGError{Error: "input error"}
	case codes.AlreadyExists:
		return http.StatusConflict, views.SWGError{Error: "already exists"}
	case codes.NotFound:
		return http.StatusNotFound, views.SWGError{Error: "not found"}
	case codes.Unimplemented, codes.Internal:
		return http.StatusBadGateway, views.SWGError{Error: "check logs on service"}
	default:
		return http.StatusBadGateway, views.SWGError{Error: "check logs on service"}
	}
}
