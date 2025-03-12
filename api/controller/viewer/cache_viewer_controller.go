package viewer

import (
	"context"
	"course_seckill_clean_architecture/domain"
	interfaces "course_seckill_clean_architecture/interface"
)

type ViewerController struct {
	db interfaces.Database
	cache interfaces.Cache
}

func NewViewerController(ctx context.Context, internal *domain.Internals) *ViewerController {
	controller := &ViewerController{
		db: internal.Db,
		cache: internal.Cache,
	}
	return controller
}

func (v *ViewerController) GetAllOrderStatus(ctx context.Context) (map[string]string, error) {
	status, err := v.cache.HGet(ctx, "order:status", "all")
	if err != nil {
		return nil, err
	}
	return status.(map[string]string), nil
}

func (v *ViewerController) GetAllStock(ctx context.Context) (map[string]string, error) {
	stock, err := v.cache.HGet(ctx, "course:stock", "all")
	if err != nil {
		return nil, err
	}
	return stock.(map[string]string), nil
}