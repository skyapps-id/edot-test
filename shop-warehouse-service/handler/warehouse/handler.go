package warehouse

import (
	"github.com/labstack/echo/v4"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/usecase/warehouse"
)

type WarehouseHandler interface {
	Create(c echo.Context) (err error)
	Gets(c echo.Context) (err error)
	Get(c echo.Context) (err error)
	CreateWarehouseProduct(c echo.Context) (err error)
	GetMaxQuantityByProductUUIDs(c echo.Context) (err error)
	ProductStockAddition(c echo.Context) (err error)
	ProductStockReduction(c echo.Context) (err error)
	GetProductStock(c echo.Context) (err error)
	WarehouseUpdateActive(c echo.Context) (err error)
	ProductRestock(c echo.Context) (err error)
	TransferStock(c echo.Context) (err error)
}

type handler struct {
	warehouseUsecase warehouse.WarehouseUsecase
}

func NewHandler(usecase warehouse.WarehouseUsecase) WarehouseHandler {
	return &handler{
		warehouseUsecase: usecase,
	}
}
