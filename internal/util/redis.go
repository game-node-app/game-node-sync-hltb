package util

func RedisURL() string {
	return GetEnv("REDIS_URL", "localhost:9112")
}
