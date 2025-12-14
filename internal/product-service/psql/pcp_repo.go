package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/autumnterror/breezynotes/pkg/utils/format"
)

func (d Driver) GetProductColorPhotos(
	ctx context.Context,
	productID string,
	colorID string,
) (any, error) {
	op := "psql.Get.ProductColorPhotos"
	var query string
	var scanner EntityScannerRow
	scanner = &PCPScannerRow{}
	query = `SELECT product_id, color_id, photos FROM product_color_photos WHERE product_id = $1 AND color_id = $2`
	res := d.Driver.QueryRowContext(ctx, query, productID, colorID)
	if err := scanner.Scan(res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, format.Error(op, ErrNotFound)
		}
		return nil, format.Error(op, err)
	}
	return scanner.Get(), nil
}

func (d Driver) DeleteProductColorPhotos(
	ctx context.Context,
	productID string,
	colorID string,
) error {
	op := "psql.deleteProductColorPhotos"
	query := `DELETE FROM product_color_photos WHERE product_id = $1 AND color_id = $2`
	res, err := d.Driver.ExecContext(ctx, query, productID, colorID)
	if err != nil {
		return format.Error(op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return format.Error(op, err)
	}
	if rowsAffected == 0 {
		return format.Error(op, ErrNotFound)
	}
	return nil
}
