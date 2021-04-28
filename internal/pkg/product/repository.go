package product

import (
	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/models"
)

//go:generate mockgen -destination=./mock/mock_repository.go -package=mock github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/product Repository

type Repository interface {
	SelectProductById(productId uint64) (*models.Product, error)
	GetCountPages(paginator *models.PaginatorProducts, filterString string) (int, error)
	CreateSortString(paginator *models.PaginatorProducts) (string, error)
	SelectRangeProducts(paginator *models.PaginatorProducts,
		sortString, filterString string) ([]*models.ViewProduct, error)
	CreateFilterString(filter *models.ProductFilter) string
}
