package seckill

import (
	"context"
	"course_seckill_clean_architecture/domain"
	interfaces "course_seckill_clean_architecture/interface"
	"errors"
	"strconv"
)

func (s *SeckillController) CreateOrder(ctx context.Context, uid int, cid int) error {
	database := s.db

	fn := func(ctx context.Context, db interfaces.Database, args ...interface{}) error {
		// check if the order already exists
		// args[0] is user_id, args[1] is course_id
		uid := args[0].(int)
		cid := args[1].(int)
		var err error
		err = db.Find(ctx, &domain.Order{}, "user_id = ? and course_id = ?", uid, cid)
		if err == nil {
			return errors.New("replicated order")
		}
		if errors.Is(err, db.NotFoundError()) {
			// continue excute tx
		} else {
			return errors.New("query order failed")
		}

		// decrease the stock of the course if enough stock
		if err = db.Update(ctx, &domain.Course{}, map[string]interface{}{"stock": "stock - 1"}, "id = ? AND stock > min_stock", cid); err != nil {
			return errors.New("update stock failed")
		}

		// create order
		if err = db.Create(ctx, &domain.Order{}, &domain.Order{
			UserID: uid,
			CourseID: cid,
		}); err != nil {
			return errors.New("create order failed")
		}

		return nil
	}


	rb := func(ctx context.Context, client interfaces.Database, err error, args ...interface{}) error {
		// args[0] is user_id, args[1] is course_id
		uid := args[0].(int)
		cid := args[1].(int)

		s.rollbackStock(ctx, cid)
		if err != errors.New("replicated order") {
			s.updateOrderStatus(ctx, uid, cid, -1)
		}
		return err
	}

	err := database.Transaction(ctx, fn, rb, uid, cid)
	if err != nil {
		return err
	}
	s.updateOrderStatus(ctx, uid, cid, 1)
	return nil
}

func (s *SeckillController) rollbackStock(ctx context.Context, cid int) {
	cache := s.cache
	cache.HIncrBy(ctx, "course:stock", strconv.Itoa(cid), 1)
}

func (s *SeckillController) updateOrderStatus(ctx context.Context, uid int, cid int, status int) {
	cache := s.cache
	cache.HSet(ctx, "order:status", strconv.Itoa(uid)+":"+strconv.Itoa(cid), status)
}
