package rabbit

// TopicExchange declares an exchange of type 'topic'.
func (conn *Conn) TopicExchange(name string) error {
	return conn.Channel.ExchangeDeclare(
		name,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}