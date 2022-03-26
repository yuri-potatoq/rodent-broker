package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer interface {
	GetTopicState(string) TopicState
	ProduceAtTopic (string, map[string]string) error
	AddTopic(topic Topic)
}

type BaseProducer struct {
	topics   map[string]Topic
	producer *kafka.Producer
}

func (bp *BaseProducer) GetTopicState(topic string) TopicState {
	return bp.topics[topic].GetTopicState()
}

func (bp *BaseProducer) ProduceAtTopic(topic string, data map[string]string) error {
	defer func() {
		bp.producer.Flush(15 * 1000)
    	bp.producer.Close()
	}()

	go func() {
        for e := range bp.producer.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
                } else {
                    log.Printf("Produced event to topic %s: key = %-10s value = %s\n",
                        *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
                }
            }
        }
    }()
	
	for k, v := range data {
		if err := bp.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(k),
			Value:          []byte(v),
		}, nil); err != nil {
			return err
		}
	}


	return nil
}

func (bp *BaseProducer) AddTopic(topic Topic) {
	bp.topics[topic.Name] = topic
}

func DefaultProducer(topic string) (KafkaProducer, error) {
	producer, err := kafka.NewProducer(GetKafkaConfig(KafkaEnv))
	if err != nil {
		log.Println("Error on create producer")
		return nil, err
	}

	return &BaseProducer{
		topics:   map[string]Topic{ topic: NewTopic(topic, "default") },
		producer: producer,
	}, nil
}
