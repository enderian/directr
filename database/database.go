package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/enderian/directrd/types"
	"log"
)

func SetupDatabase(ctx types.Context) types.Context {
	db, err := gorm.Open("postgres", ctx.Conf().Database)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %s", err.Error())
	}

	db.AutoMigrate(&types.Session{}, &types.Terminal{}, &types.User{})

	log.Printf("Connected to Postgres database successfully!")
	return types.NewContextWithDB(ctx, db)
}

func CloseDatabase(ctx types.Context) {
	db := ctx.DB()
	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close the Postgres connection: %s", err.Error())
	}
}
