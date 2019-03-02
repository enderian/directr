package types

import (
	"context"
	"github.com/go-pg/pg"
	"io"
	"time"
)

const loggerContext = "logger"
const confContext = "conf"
const databaseContext = "db"

type Context struct {
	ctx context.Context //Actual enclosed context
}

func NewContext(ctx context.Context, conf *Configuration, logger io.Writer) Context {
	return Context{context.WithValue(context.WithValue(ctx, loggerContext, logger), confContext, conf)}
}

func NewContextWithDB(ctx context.Context, db *pg.DB) Context {
	return Context{context.WithValue(ctx, databaseContext, db)}
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

func (c Context) DB() *pg.DB {
	return c.ctx.Value(databaseContext).(*pg.DB)
}

func (c Context) Logger() io.Writer {
	return c.ctx.Value(loggerContext).(io.Writer)
}
