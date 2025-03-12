package bootstrap

import (
		"log"

	"course_seckill_clean_architecture/internal/repository/redis"
	"course_seckill_clean_architecture/interface"
)

func NewRedis(env *Env) interfaces.Cache {
	host := env.CacheHost
	port := env.CachePort
	pass := env.CachePassword
	db := env.CacheDB
	
	cache, err := redis.NewInstance(host, port, pass, db)
	if err != nil {
		log.Fatal("Failed to create redis client:", err)
	}

	return cache
}

func CloseRedis(cache interfaces.Cache) {
	err := cache.Close()
	if err != nil {
		log.Fatal("Failed to close redis:", err)
	}
}