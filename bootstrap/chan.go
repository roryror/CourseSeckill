package bootstrap

import (
	"course_seckill_clean_architecture/internal/mq/channel"
	"course_seckill_clean_architecture/interface"
)


func NewChannel(env *Env) interfaces.Channel {
	size := env.ChanSize
	batchSize := env.ChanBatchSize
	duration := env.ChanDuration

	channel := channel.NewInstance(size, batchSize, duration)
	return channel
}

func CloseChannel(c interfaces.Channel) {
	c.Close()
}
