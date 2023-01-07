package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "shop/api/service/user/v1"
	"shop/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	log    *log.Helper
	client v1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, client v1.UserClient) (*Data, error) {
	return &Data{
		log:    log.NewHelper(log.With(logger, "module", "data")),
		client: client,
	}, nil
}

// NewUserServiceClient 连接用户服务
func NewUserServiceClient(ac *conf.Auth, sr *conf.Service, rd registry.Discovery) v1.UserClient {

	clientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.User.Endpoint), // consul
		grpc.WithDiscovery(rd),
		grpc.WithMiddleware(recovery.Recovery()),
		grpc.WithTimeout(3*time.Second),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewUserClient(clientConn)
}

// NewRegistrar add consul
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme

	client, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	return consul.New()
}
