package database

import (
	"context"
	"github.com/enderian.directrd/types"
	"github.com/go-pg/pg"
	"log"
)

func SetupDatabase(configuration *types.Configuration) context.Context {
	opts, err := pg.ParseURL(configuration.Database)
	if err != nil {
		log.Fatalf("Failed to parse Postgres URL: %s", err.Error())
	}
	db := pg.Connect(opts)
	if tx, err := db.Begin(); err == nil {
		_ = tx.Commit()
	} else {
		log.Fatalf("Failed to connect to Postgres: %s", err.Error())
	}

	log.Printf("Connected to Postgres database %s@%s successfully!", opts.Database, opts.Addr)
	return context.WithValue(context.Background(), types.CtxDatabase, db)
}

func CloseDatabase(ctx context.Context) {
	db := ctx.Value(types.CtxDatabase).(*pg.DB)
	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close the Postgres connection: %s", err.Error())
	}
}
