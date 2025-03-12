package channel

import (
	interfaces "course_seckill_clean_architecture/interface"	
)

type channel struct {
	ch chan interface{}
	batchSize int
	duration int
}

func NewInstance(size int, batchSize int, duration int) interfaces.Channel {
	return &channel{
		ch: make(chan interface{}, size),
		batchSize: batchSize,
		duration: duration,
	}
}

func (c *channel) BatchSize() int {
	return c.batchSize
}

func (c *channel) Duration() int {
	return c.duration
}

func (c *channel) Ch() chan interface{} {
	return c.ch
}

func (c *channel) Close() error {
	close(c.ch)
	return nil
}
