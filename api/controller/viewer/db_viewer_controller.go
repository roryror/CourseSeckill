package viewer

import (
	"context"
	"course_seckill_clean_architecture/domain"

)

func (v *ViewerController) GetAllOrders(ctx context.Context) ([]domain.Order, error) {
	var orders []domain.Order
	err := v.db.Find(ctx, &orders, "all")
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (v *ViewerController) GetAllCourses(ctx context.Context) ([]domain.Course, error) {
	var courses []domain.Course
	err := v.db.Find(ctx, &courses, "all")
	if err != nil {
		return nil, err
	}
	return courses, nil
}