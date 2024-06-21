package main

import (
	"context"
	"log"

	"github.com/jmrflora/bazarTudao/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	c, err := client.Cliente.Create().SetCpf("1111111111").SetNome("nome teste").
		SetEmail("email").
		SetTelefone("999999999").
		Save(context.Background())
	if err != nil {
		panic(err.Error())
	}
	println(c.Nome)
}
