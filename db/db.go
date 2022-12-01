package db

import (
	"context"
	"log"
	"voidmanager/db/ent"
	"voidmanager/db/ent/migrate"
	"voidmanager/utils"

	_ "github.com/lib/pq"
)

func InitDB(cfg *utils.Config) *ent.Client {
	client, err := ent.Open("postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to open database client: %v", err.Error())
	}

	if err := client.Schema.Create(context.Background(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("failed to create schema resources: %v", err.Error())
	}

	return client
}
