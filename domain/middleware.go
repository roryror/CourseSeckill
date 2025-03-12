package domain

import (
	interfaces "course_seckill_clean_architecture/interface"
)

//lint:ignore U1000 
type Internals struct {
	Db interfaces.Database
	Cache interfaces.Cache
	Mq interfaces.MsgQueue
	Channel interfaces.Channel
}
