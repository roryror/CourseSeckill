package seckill

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func (s *SeckillController) RunRoutines(ctx context.Context) {
	go s.NewProducerRoutine(ctx)

	groupSize := s.mq.GroupSize()
	for i := 0; i < groupSize; i++ {
		go s.NewConsumerRoutine(ctx, i)
	}

	activator := s.channel
	activator.Ch() <- "activate reader"
}

func (s *SeckillController) NewProducerRoutine(ctx context.Context) {
	buffer := s.channel
	writer := s.mq

	var batch = make([]interface{}, 0, buffer.BatchSize())
	ticker := time.NewTicker(time.Duration(buffer.Duration()) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-buffer.Ch():
			if !ok {
				return
			}
			batch = append(batch, msg)
			
			if len(batch) >= buffer.BatchSize() {
				writer.Write(ctx, batch)
				batch = batch[:0]
			}
		// send messages every flushInterval
		case <-ticker.C:
			if len(batch) > 0 {
				writer.Write(ctx, batch)
				batch = batch[:0]
			}
		}
	}
	
}

func (s *SeckillController) NewConsumerRoutine(ctx context.Context, rid int) {
	reader := s.mq

	parseMessage := func(message string) (int, int, error) {
		parts := strings.Split(message, ":")
		uid, _ := strconv.Atoi(parts[0])
		cid, _ := strconv.Atoi(parts[1])
		return uid, cid, nil
	}
	
	fmt.Println("Consumer is running")
	for {
		msg, err := reader.Read(ctx, rid)
		fmt.Println("Consumer received message:", msg)
		if err != nil {
			if err == io.EOF {
				continue
			}
			fmt.Println("Failed to read message:", err)
			continue
		}
		
		msgStr, _ := msg.(string)
		if msgStr == "activate reader" {
			fmt.Println("Reader activated, ready for seckill!")
			continue
		}
		
		uid, cid, err := parseMessage(msgStr)
		if err != nil {
			fmt.Println("Failed to parse message:", err)
			continue
		}
		
		err = s.CreateOrder(ctx, uid, cid)
		if err != nil {
			fmt.Println("Failed to create order:", err)
		} 
	}
}

