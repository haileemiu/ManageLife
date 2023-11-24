package client

import (
	"context"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/haileemiu/manage-life/ent"
)

var entClient *ent.Client

func GetEnt() (*ent.Client, error) {
	if entClient == nil {
		// TODO: use config for database connection string
		client, err := ent.Open("postgres", "postgres://postgres:postgres@db/manage_life_dev?sslmode=disable")
		if err != nil {
			return nil, err
		}

		writeSchemaFile(client)

		if err := client.Schema.Create(context.Background()); err != nil {
			return nil, err
		}

		entClient = client
	}

	return entClient, nil
}

func writeSchemaFile(entClient *ent.Client) {
	f, err := os.OpenFile("db.sql", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening schema file: %s", err)
		return
	}
	defer f.Close()

	if err = entClient.Schema.WriteTo(context.Background(), f); err != nil {
		fmt.Fprintf(os.Stderr, "error writing schema to file: %s", err)
		return
	}
}
