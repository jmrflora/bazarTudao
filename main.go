package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/jmrflora/bazarTudao/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:bz.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// c, err := client.Cliente.Create().SetCpf("1111111111").SetNome("nome teste").
	// 	SetEmail("email").
	// 	SetTelefone("999999999").
	// 	Save(context.Background())
	// if err != nil {
	// 	panic(err.Error())
	// }

	// c, err := client.Cliente.Get(context.Background(), 1)
	// if err != nil {
	// 	panic(err.Error())
	// }

	client.Cliente.Create()
}

func LoadPedidos() error {
	file, err := os.Open("pedidos.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
	}
}
