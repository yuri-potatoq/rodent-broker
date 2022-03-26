package main


type kafkaConsumer interface {
	GetTopicState(Topic) TopicState
	ConsumeTopic(Topic) interface{}
}