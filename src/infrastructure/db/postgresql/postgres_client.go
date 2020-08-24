package postgresql_client

import (
	"context"
	"fmt"
	configs "github.com/nhsh1997/go-boilerplate/config"
	"github.com/nhsh1997/go-boilerplate/ent"
	"log"
)

type PostgreClient struct {
	Client *ent.Client
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "test"
)


func NewPostgresClient (configuration *configs.Configuration) *PostgreClient{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	client, err := ent.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed opening connection to postgresql: %v", err)
	}
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &PostgreClient{
		Client: client,
	}
}
