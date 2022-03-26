package main

import "github.com/confluentinc/confluent-kafka-go/kafka"


var configs = map[string]*kafka.ConfigMap {
	"local": { 
		"bootstrap.servers": "localhost:9092",
		"group.id": "custom-group",
		"auto.offset.reset": "earliest",
	},
	"dev": { 
		"bootstrap.servers": "broker:9092",
		"group.id": "dev-group",
		"auto.offset.reset": "earliest",
	},
}


func GetKafkaConfig(tag string) *kafka.ConfigMap {
	return configs[tag]
}
