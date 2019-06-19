package database

import (
	"github.com/enderian/directrd/types"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

func SetupDatabase(ctx types.Context) types.Context {
	opts, err := pg.ParseURL(ctx.Conf().Database)
	if err != nil {
		log.Fatalf("Failed to parse Postgres URL: %s", err.Error())
	}
	db := pg.Connect(opts)
	if tx, err := db.Begin(); err == nil {
		_ = tx.CreateTable((*types.User)(nil), &orm.CreateTableOptions{IfNotExists: true})
		_ = tx.CreateTable((*types.Session)(nil), &orm.CreateTableOptions{IfNotExists: true})
		_ = tx.CreateTable((*types.Terminal)(nil), &orm.CreateTableOptions{IfNotExists: true})
		_ = tx.Commit()
	} else {
		log.Fatalf("Failed to connect to Postgres: %s", err.Error())
	}

	log.Printf("Connected to Postgres database %s@%s successfully!", opts.Database, opts.Addr)
	return types.NewContextWithDB(ctx, db)
}

func CloseDatabase(ctx types.Context) {
	db := ctx.DB()
	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close the Postgres connection: %s", err.Error())
	}
}
