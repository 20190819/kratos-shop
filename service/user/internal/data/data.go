package data

import (
	"gorm.io/driver/mysql"
	"os"
	"time"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel/v8"
	redisV8 "github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	slog "log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,NewDB,NewRedis,NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redisV8.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {

	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},
	)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}

	return db
}

func NewRedis(c *conf.Data) *redisV8.Client {
	rdb := redisV8.NewClient(&redisV8.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.NewTracingHook())
	if err := rdb.Close(); err != nil {
		log.Error(err)
	}
	return rdb
}
