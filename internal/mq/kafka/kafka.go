package kafka

import (
	"context"
	interfaces "course_seckill_clean_architecture/interface"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type kafkaMsgQueue struct {
	conn *kafka.Conn
	writer *kafka.Writer
	readers []*kafka.Reader
	groupSize int
}

func NewInstance(host string, port string, topic string, partition int, replicationFactor int, brokers []string, groupID string, groupSize int, minBytes int, maxBytes int, startOffset int, maxWait int, readBackoffMin int, readBackoffMax int, commitInterval int) interfaces.MsgQueue {
	conn := NewConn(host, port, topic, partition, replicationFactor)
	writer := NewWriter(conn, topic, partition, replicationFactor, brokers)
	readers := NewReaders(conn, topic, groupID, groupSize, brokers, minBytes, maxBytes, startOffset, maxWait, readBackoffMin, readBackoffMax, commitInterval)
	return &kafkaMsgQueue{conn, writer, readers, groupSize}
}

func NewConn(host string, port string, topic string, partition int, replicationFactor int) *kafka.Conn {
	conn, err := kafka.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal("Failed to connect to kafka:", err)
	}

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partition,
			ReplicationFactor: replicationFactor,
		},
	}
	
	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Fatal("Failed to create topics:", err)
	}
	return conn
}

func NewWriter(conn *kafka.Conn, topic string, partition int, replicationFactor int, brokers []string) *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
		Balancer: &kafka.Hash{},
	})
	return writer
}

func NewReaders(conn *kafka.Conn, topic string, groupID string, groupSize int, brokers []string, minBytes int, maxBytes int, startOffset int, maxWait int, readBackoffMin int, readBackoffMax int, commitInterval int) []*kafka.Reader {
	readers := make([]*kafka.Reader, groupSize)
	for i := 0; i < groupSize; i++ {
		readers[i] = kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
			MinBytes: minBytes,
			MaxBytes: maxBytes,
			StartOffset: int64(startOffset),
			MaxWait:       time.Millisecond * time.Duration(maxWait),
			ReadBackoffMin: time.Millisecond * time.Duration(readBackoffMin),
			ReadBackoffMax: time.Millisecond * time.Duration(readBackoffMax),
			CommitInterval: time.Millisecond * time.Duration(commitInterval),
		})
	}
	return readers
}

func (k *kafkaMsgQueue) GroupSize() int {
	return k.groupSize
}

func (k *kafkaMsgQueue) Write(ctx context.Context, msg []interface{}) error {
	kafkaMsg := make([]kafka.Message, len(msg))
	for i, m := range msg {
		kafkaMsg[i] = kafka.Message{Value: []byte(m.(string))}
	}
	return k.writer.WriteMessages(ctx, kafkaMsg...)
}

func (k *kafkaMsgQueue) Read(ctx context.Context, rid int) (interface{}, error) {
	msg, err := k.readers[rid].ReadMessage(ctx)
	if err != nil {
		fmt.Println("Failed to read message:", err)
		return "", err
	}
	// fmt.Println("Message read successfully:", string(msg.Value))
	return string(msg.Value), nil
}

func (k *kafkaMsgQueue) Close() error {
	return k.conn.Close()
}
