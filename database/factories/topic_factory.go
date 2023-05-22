package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/practice/app/models/topic"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: "3",
			UserID:     "2",
		}
		objs = append(objs, topicModel)
	}

	return objs
}
