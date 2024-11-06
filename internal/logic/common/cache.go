package common

import (
	"context"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/wingfeng/idxadmin/internal/conf"
)

var Cache *gcache.Cache

func init() {
	ctx := context.Background()
	options := conf.GetOptions(ctx)
	var (
		err error

		redisConfig = &gredis.Config{
			Address: options.RedisUrl,
			Db:      options.RedisDB,
		}
	)

	Cache = gcache.New()
	// Create redis client object.
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// Create redis cache adapter and set it to cache object.
	Cache.SetAdapter(gcache.NewAdapterRedis(redis))
}
