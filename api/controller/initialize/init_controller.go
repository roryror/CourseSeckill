package initialize

import (
	"context"
	"course_seckill_clean_architecture/domain"
	interfaces "course_seckill_clean_architecture/interface"
	"fmt"
)

type InitController struct {
	db interfaces.Database
	cache interfaces.Cache
}

func NewInitController(ctx context.Context, internal *domain.Internals) *InitController {
	controller := &InitController{
		db: internal.Db,
		cache: internal.Cache,
	}
	controller.InitDatabase(ctx)
	return controller
}

func (i *InitController) InitDatabase(ctx context.Context) error {
	db := i.db
	courseList := []domain.Course{
		{
			ID: 1,
			Name: "Course 1",
			Stock: 100,
			MaxStock: 100,
			MinStock: 20,
		},
		{
			ID: 2,
			Name: "Course 2",
			Stock: 150,
			MaxStock: 150,
			MinStock: 50,
		},
		{
			ID: 3,
			Name: "Course 3",
			Stock: 100,
			MaxStock: 150,
			MinStock: 0,
		},
	}
	err := db.Delete(ctx, &domain.Course{}, "all")
	if err != nil {
		return err
	}
	err = db.Create(ctx, &domain.Course{}, courseList)
	if err != nil {
		fmt.Println("Failed to create course:", err)
		return err
	}
	err = db.Delete(ctx, &domain.Order{}, "all")
	if err != nil {
		return err
	}
	err = db.Create(ctx, &domain.Order{})
	if err != nil {
		return err
	}
	return nil
}

