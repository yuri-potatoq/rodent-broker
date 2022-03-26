package main

type Topic struct {
	Name       string
	topicState TopicState
	owner      string
}

type TopicState struct{}

func (t Topic) GetTopicState() TopicState {
	return t.topicState
}

func NewTopic(name string, owner string) Topic {
	return Topic{
		owner: owner,
		Name:  name,
	}
}
