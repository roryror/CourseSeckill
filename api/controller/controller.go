package controller

import (
	"context"
	"course_seckill_clean_architecture/domain"
	"course_seckill_clean_architecture/api/controller/seckill"
	"course_seckill_clean_architecture/api/controller/initialize"
	"course_seckill_clean_architecture/api/controller/viewer"
)

type Controller struct {
	InitController    *initialize.InitController
	SeckillController *seckill.SeckillController
	ViewerController  *viewer.ViewerController
}

func NewController(ctx context.Context, internal *domain.Internals) *Controller {
	controller := &Controller{
		InitController: initialize.NewInitController(ctx, internal),
		SeckillController: seckill.NewSeckillController(ctx, internal),
		ViewerController: viewer.NewViewerController(ctx, internal),
	}

	return controller
}
