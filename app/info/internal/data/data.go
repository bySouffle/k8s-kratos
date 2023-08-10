package data

import (
	v1 "bysos/api/user/v1"
	"bysos/app/info/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewInfoRepo, NewRedis)

// Data .
type Data struct {
	// TODO wrapped database client
	redis    *redis.Client
	httpUser v1.UserHTTPClient
}

// NewData .
func NewData(c *conf.Data, cr *conf.Registry, discovery registry.Discovery, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{redis: NewRedis(c), httpUser: NewInfoHttpClient(cr, discovery)}, cleanup, nil
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	return rdb
}

func NewInfoHttpClient(cfg *conf.Registry, discovery registry.Discovery) v1.UserHTTPClient {
	hConn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithEndpoint(cfg.Endpoint.Discovery),
		http.WithDiscovery(discovery),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer hConn.Close()

	return v1.NewUserHTTPClient(hConn)
}
