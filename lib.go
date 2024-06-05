package msq

type Publisher struct{}

func NewPublisher(topics ...string) {}

type Consumer struct{}

func NewConsumer(topic string, name string) *Consumer {
	return &Consumer{}
}

func (c *Consumer) ResetOffset() {}

func (c *Consumer) Read() {}
