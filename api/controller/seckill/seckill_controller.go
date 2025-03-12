package seckill

import (
	"context"
	"errors"
	"fmt"
	"course_seckill_clean_architecture/interface"
	"course_seckill_clean_architecture/domain"
)

type SeckillController struct {
	db interfaces.Database
	cache interfaces.Cache
	mq interfaces.MsgQueue
	channel interfaces.Channel
}

func NewSeckillController(ctx context.Context, internal *domain.Internals) *SeckillController {
	controller := &SeckillController{
		db: internal.Db,
		cache: internal.Cache,
		mq: internal.Mq,
		channel: internal.Channel,
	}
	controller.RunRoutines(ctx)
	return controller
}

func (s *SeckillController) RunSeckill(ctx context.Context, cid int, uid int) error {
	// Here cache works as a filter, to check if the user has already ordered the course
	// and channel works as a buffer, between Cache and Message Queue
	filter := s.cache
	buffer := s.channel

	// Lua script to make sure ACID
	luaScript := `
		local uid = ARGV[1] 
		local cid = ARGV[2]
		local requestKey = ARGV[1] .. ":" .. ARGV[2]
		
		if redis.call("SISMEMBER", "order:requests", requestKey) == 1 then
			return 0
		end

		redis.call("SADD", "order:requests", requestKey)
		redis.call("EXPIRE", "order:requests", 5)
		
		local stock = redis.call("HGET", "course:stock", cid)
		if not stock then
			return -1
		end
		
		stock = tonumber(stock)
		if stock <= 0 then
			return -1
		end
		
		local new_stock = redis.call("HINCRBY", "course:stock", cid, -1)
		if new_stock >= 0 then
			return 1
		else
			redis.call("HINCRBY", "course:stock", cid, 1)
			return -1
		end
	`
	// fmt.Println("Running seckill lua")

	// Run the lua script
	result, err := filter.RunScript(ctx, luaScript, []string{}, uid, cid)
	if err != nil {
		return errors.New("script execution failed")
	}
	intResult := result.(int64)
	if intResult == 1 {
		fmt.Println("uid:", uid, "cid:", cid, "result:", intResult)
	}

	// If the script returns 1, it means the user has successfully ordered the course (pending order)
	// So we send the message to the message queue
	// Otherwise, we rollback the stock and return specific error
	
	switch intResult {
	case 1:
		buffer.Ch() <- fmt.Sprintf("%d:%d", uid, cid)
		// err := buffer.In(ctx, fmt.Sprintf("%d:%d", uid, cid))
		// if err != nil {
		// 	s.rollbackStock(ctx, cid)
		// 	return err
		// }
		// fmt.Println("In buffer")
		return nil
	case 0:
		return errors.New("repeat order")
	case -1:
		return errors.New("out of stock")
	default:
		return fmt.Errorf("unexpected result: %v", result)
	}
}




		