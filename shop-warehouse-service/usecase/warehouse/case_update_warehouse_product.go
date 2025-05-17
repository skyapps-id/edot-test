package warehouse

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/entity"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/apperror"
	"github.com/skyapps-id/edot-test/shop-warehouse-service/pkg/tracer"
	"gorm.io/gorm"
)

func (uc *usecase) ProductRestock(ctx context.Context, req ProductRestockReqeust) (resp ProductRestockResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.ProductRestock")
	defer span.End()

	warehouse, err := uc.warehouseRepository.FindByUUID(ctx, req.WarehouseUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("warehouse not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse err %w", err))
		return
	}

	if !warehouse.Active {
		err = apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("warehouse inactive"))
		return
	}

	_, err = uc.warehouseProductRepository.GetByWarehouseUUIDAndProductUUID(ctx, req.WarehouseUUID, req.ProductUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("warehouse product not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse product err %w", err))
		return
	}

	err = uc.warehouseProductRepository.ProductStockAddition(ctx, []entity.ProductStock{
		{ProductUUID: req.ProductUUID, WarehouseUUID: req.WarehouseUUID, Quantity: req.Quantity},
	})
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to product restock err %w", err))
		return
	}

	resp.UUID = req.ProductUUID

	return
}

func (uc *usecase) ProductStockAddition(ctx context.Context, req ProductStockAdditionRequest) (resp ProductStockAdditionResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.ProductStockAddition")
	defer span.End()

	productUUIDs := make([]uuid.UUID, 0, len(req.Products))
	products := make([]entity.ProductStock, 0, len(req.Products))

	for _, row := range req.Products {
		productUUIDs = append(productUUIDs, row.ProductUUID)
		products = append(products, entity.ProductStock{
			ProductUUID:   row.ProductUUID,
			WarehouseUUID: row.WarehouseUUID,
			Quantity:      row.Quantity,
		})
	}

	err = uc.warehouseProductRepository.ProductStockAddition(ctx, products)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to product stock addition err %w", err))
		return
	}

	resp.ProductUUIDs = productUUIDs

	return
}

func (uc *usecase) ProductStockReduction(ctx context.Context, req ProductStockReductionRequest) (resp ProductStockReductionResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.ProductStockReduction")
	defer span.End()

	productUUIDs := make([]uuid.UUID, 0, len(req.Products))
	products := make([]entity.ProductStock, 0, len(req.Products))

	for _, row := range req.Products {
		productUUIDs = append(productUUIDs, row.ProductUUID)
		products = append(products, entity.ProductStock{
			ProductUUID:   row.ProductUUID,
			WarehouseUUID: row.WarehouseUUID,
			Quantity:      row.Quantity,
		})
	}

	err = uc.warehouseProductRepository.ProductStockReduction(ctx, products)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to product stock reduction err %w", err))
		return
	}

	resp.ProductUUIDs = productUUIDs

	return
}

func (uc *usecase) TransferStock(ctx context.Context, req TransferStockReqeust) (resp TransferStockResponse, err error) {
	ctx, span := tracer.Define().Start(ctx, "WarehouseUsecase.TransferStock")
	defer span.End()

	warehouseSrc, err := uc.warehouseRepository.FindByUUID(ctx, req.WarehouseUUIDSrc)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("warehouse source not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse source err %w", err))
		return
	}

	if !warehouseSrc.Active {
		err = apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("warehouse source inactive"))
		return
	}

	warehouseDst, err := uc.warehouseRepository.FindByUUID(ctx, req.WarehouseUUIDDst)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("warehouse destination not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse destination err %w", err))
		return
	}

	if !warehouseDst.Active {
		err = apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("warehouse destination inactive"))
		return
	}

	warehouseProduct, err := uc.warehouseProductRepository.GetByWarehouseUUIDAndProductUUID(ctx, req.WarehouseUUIDSrc, req.ProductUUID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = apperror.New(http.StatusNotFound, fmt.Errorf("warehouse product not found"))
		return
	}
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to get warehouse product err %w", err))
		return
	}

	if req.Quantity > warehouseProduct.Quantity {
		err = apperror.New(http.StatusUnprocessableEntity, fmt.Errorf("warehouse source quantity stock limited"))
		return
	}

	err = uc.warehouseProductRepository.TransferStock(ctx, req.ProductUUID, req.WarehouseUUIDSrc, req.WarehouseUUIDDst, req.Quantity)
	if err != nil {
		err = apperror.New(http.StatusInternalServerError, fmt.Errorf("fail to product transfer stock err %w", err))
		return
	}

	resp.UUID = req.ProductUUID

	return
}
