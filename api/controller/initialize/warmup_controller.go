package initialize

import (
	"context"
	"course_seckill_clean_architecture/domain"
	"strconv"
	"time"
)

func (i *InitController) WarmupCache(ctx context.Context) error {
	db := i.db
	cache := i.cache
	var courseList []domain.Course
	err := db.Find(ctx, &courseList, "all")
	if err != nil {
		return err
	}

	cache.Del(ctx, "course:stock")    
	cache.Del(ctx, "order:requests") 
	cache.Del(ctx, "order:status")
	
	for _, course := range courseList {
		cache.HSet(ctx, "course:stock", strconv.Itoa(course.ID), course.Stock - course.MinStock)
		cache.Expire(ctx, "course:stock", time.Minute*2)
	}
	
	cache.HSet(ctx, "order:status", "uid:cid", 0)
	cache.Expire(ctx, "order:status", time.Minute*1)

	cache.SAdd(ctx, "order:requests", "uid:cid")
	cache.Expire(ctx, "order:requests", time.Second*3)

	return nil
}

