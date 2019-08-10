package types

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis"
	"io"
	"time"
)

const loggerContext = "logger"
const confContext = "conf"
const databaseContext = "db"
const redisContext = "redis"

type Context struct {
	ctx context.Context //Actual enclosed context
}

func NewContext(ctx context.Context, conf *Configuration, logger io.Writer) Context {
	return Context{context.WithValue(context.WithValue(ctx, loggerContext, logger), confContext, conf)}
}

func NewContextWithDB(ctx context.Context, db *gorm.DB) Context {
	return Context{context.WithValue(ctx, databaseContext, db)}
}

func NewContextWithRedis(ctx context.Context, red *redis.Client) Context {
	return Context{context.WithValue(ctx, redisContext, red)}
}

func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c Context) Err() error {
	return c.ctx.Err()
}

func (c Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c Context) Conf() *Configuration {
	return c.ctx.Value(confContext).(*Configuration)
}

func (c Context) DB() *gorm.DB {
	return c.ctx.Value(databaseContext).(*gorm.DB)
}

func (c Context) Redis() *redis.Client {
	return c.ctx.Value(redisContext).(*redis.Client)
}

func (c Context) Logger() io.Writer {
	return c.ctx.Value(loggerContext).(io.Writer)
}
