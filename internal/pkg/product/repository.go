package product

import (
	"github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/models"
)

//go:generate mockgen -destination=./mock/mock_repository.go -package=mock github.com/go-park-mail-ru/2021_1_DuckLuck/internal/pkg/product Repository

type Repository interface {
	SelectProductById(productId uint64) (*models.Product, error)
	GetCountPages(category uint64, count int) (int, error)
	CreateSortString(sortKey, sortDirection string) (string, error)
	SelectRangeProducts(paginator *models.PaginatorProducts,
		sortString string) ([]*models.ViewProduct, error)
	SearchRangeProducts(searchQuery *models.SearchQuery,
		sortString string) ([]*models.ViewProduct, error)
}
