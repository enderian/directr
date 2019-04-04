package database

import (
	"github.com/enderian/directrd/types"
	"github.com/go-redis/redis"
	"log"
)

func SetupRedis(ctx types.Context) types.Context {
	opts, err := redis.ParseURL(ctx.Conf().Redis)
	if err != nil {
		log.Panicf("error when parsing redis: %v", err)
	}
	client := redis.NewClient(opts)
	return types.NewContextWithRedis(ctx, client)
}
