package bootstrap

import (
	"log"

	kafkamq "course_seckill_clean_architecture/internal/mq/kafka"
	interfaces "course_seckill_clean_architecture/interface"
)	

func NewKafka(env *Env) interfaces.MsgQueue {
	host := env.MQHost
	port := env.MQPort
	topic := env.MQTopic
	groupID := env.MQGroupID
	groupSize := env.MQGroupSize
	brokers := env.MQBrokers
	partition := env.MQPartition
	replicationFactor := env.MQReplicationFactor
	minBytes := env.MQMinBytes
	maxBytes := env.MQMaxBytes
	startOffset := env.MQStartOffset
	maxWait := env.MQMaxWait
	readBackoffMin := env.MQReadBackoffMin
	readBackoffMax := env.MQReadBackoffMax
	commitInterval := env.MQCommitInterval

	mq := kafkamq.NewInstance(host, port, topic, partition, replicationFactor, brokers, groupID, groupSize, minBytes, maxBytes, startOffset, maxWait, readBackoffMin, readBackoffMax, commitInterval)
	return mq
}

func CloseKafka(mq interfaces.MsgQueue) {
	err := mq.Close()
	if err != nil {
		log.Fatal("Failed to close kafka:", err)
	}
}
