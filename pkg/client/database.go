package client

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/haileemiu/manage-life/ent"
)

var entClient *ent.Client

func GetEnt() (*ent.Client, error) {
	if entClient == nil {
		// TODO: use config for database connection string
		client, err := ent.Open("postgres", "postgres://postgres:postgres@db/ass_app_dev?sslmode=disable")
		if err != nil {
			return nil, err
		}

		if err := client.Schema.Create(context.Background()); err != nil {
			return nil, err
		}

		entClient = client
	}

	return entClient, nil
}
