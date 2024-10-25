package util

func RMQUrl() string {
	return GetEnv("RABBITMQ_URL", "amqp://gamenode:gamenode@localhost:5672")
}
