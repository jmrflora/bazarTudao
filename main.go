package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmrflora/bazarTudao/crud"
	"github.com/jmrflora/bazarTudao/ent"
	"github.com/jmrflora/bazarTudao/handler"
	"github.com/labstack/echo/v4"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	crud, err := crud.NewCrud("sqlite3", "file:bz.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer crud.CloseCrud()
	// Run the auto migration tool.
	if err := crud.CreateSchema(); err != nil {
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

	// err = LoadPedidos(crud)
	// if err != nil {
	// 	panic(err)
	// }

	// err = TratarPedidosIntocados(*crud)
	// if err != nil {
	// 	println(err.Error())
	// }
	//
	// err = ComprarOqueFalta(*crud)
	// if err != nil {
	// 	panic(err)
	// }

	// err = TratarPedidosParciais(*crud)
	// if err != nil {
	// 	panic(err)
	// }

	h := handler.NewHandler(crud)
	e := echo.New()

	e.GET("/ordens", h.HandleGetOrdens)

	e.GET("/ordens/completas", h.HandleGetOrdensCompletas)

	e.GET("/ordens/intocadas", h.HandleGetOrdensIntocadas)

	e.GET("/ordens/parciais", h.HandleGetOrdensParciais)

	e.GET("/envios", h.HandleGetEnvios)

	e.GET("/itens/:id", h.HandleGetItensPorIdOrdem)

	e.GET("/produtos", h.HandleGetProdutos)

	e.GET("/clientes", h.HandlerGetAllClientes)

	e.POST("/load", h.LoadPedidos)

	e.POST("/tratar/intocadas", h.TratarPedidosIntocados)

	e.Logger.Fatal(e.Start(":1323"))
}

func LoadPedidos(crud *crud.Crud) error {
	file, err := os.Open("pedidos.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return err
	}
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var cliente *ent.Cliente

	for _, record := range records {
		println("aqui")
		ordemid := record[0]
		id, _ := strconv.Atoi(ordemid)
		ordem, err := crud.GetOrdem(id)
		// não existe ordem
		if err != nil {
			cliente, err = crud.GetCliente(record[1])
			// não existe cliente
			if err != nil {
				cliente, err = crud.AddCliente(record[2], record[1], record[4], record[9], record[3])
				if err != nil {
					return err
				}
			}
			ordem, err = crud.AddOrdem(cliente, id)
			if err != nil {
				return err
			}
		}
		produto, err := crud.GetProduto(record[5])
		// não existe produto
		if err != nil {
			produto, err = crud.AddProduto(record[5], record[6])
			if err != nil {
				println(err.Error())
				return err
			}
		}
		println("record[8]")
		println(record[8])
		precoU, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			return err
		}
		quant, err := strconv.Atoi(record[7])
		if err != nil {
			return err
		}
		precoT := precoU * float64(quant)

		item, err := crud.AddItem(produto, ordem, quant, precoU, precoT)
		if err != nil {
			return err
		}

		ordem.Update().SetPrecoDaOrdem(ordem.PrecoDaOrdem + item.PrecoTotal).SaveX(context.Background())

		println("/n")
		fmt.Printf("item: %v\n", item)
		println("ola de novo")

	}
	println("ola")
	return nil
}

func TratarPedidosIntocados(crud crud.Crud) error {
	ordens, err := crud.GetOrdensIntocadas()
	println("\n")
	fmt.Printf("ordens: %v\n", ordens)
	if err != nil {
		return err
	}
	for _, ordem := range ordens {
		itens, err := crud.GetItens(ordem)
		if err != nil {
			return err
		}
		fmt.Printf("itens: %v\n", itens)
		i := 0
		for _, item := range itens {
			println("/n")
			fmt.Printf("item: %v\n", item)
			quantEstoque := item.Edges.Produto.QuantNoEstoque
			quantItem := item.Quantidade
			if quantEstoque >= quantItem {
				println("quantEstoque >= quantItem")
				_, err := crud.AddEnvio(item)
				if err != nil {
					return err
				}
				item.Edges.Produto.Update().SetQuantNoEstoque(quantEstoque - quantItem).SaveX(context.Background())
			} else {
				i++
			}
		}
		if i > 0 {
			ordem.Update().SetStatus("parcial").SaveX(context.Background())
		} else {
			ordem.Update().SetStatus("completa").SaveX(context.Background())
		}
	}
	return nil
}

func TratarPedidosParciais(crud crud.Crud) error {
	ordens, err := crud.GetOrdensParciais()
	if err != nil {
		return err
	}
	for _, ordem := range ordens {
		itens, err := crud.GetItensNaoEnviadosPorOrdem(ordem)
		if err != nil {
			return err
		}
		i := 0
		for _, item := range itens {
			println("/n")
			fmt.Printf("item: %v\n", item)
			quantEstoque := item.Edges.Produto.QuantNoEstoque
			quantItem := item.Quantidade
			if quantEstoque >= quantItem {
				println("quantEstoque >= quantItem")
				_, err := crud.AddEnvio(item)
				if err != nil {
					return err
				}
				item.Edges.Produto.Update().SetQuantNoEstoque(quantEstoque - quantItem).SaveX(context.Background())
			} else {
				i++
			}
		}
		if i == 0 {
			ordem.Update().SetStatus("completa").SaveX(context.Background())
		}
	}
	return nil
}

func ComprarOqueFalta(crud crud.Crud) error {
	itens, err := crud.GetItensNaoEnviados()
	if err != nil {
		return err
	}

	for _, item := range itens {

		fmt.Printf("item: %v\n", item)

		if item.Edges.Produto.QuantNoEstoque < item.Quantidade {
			err = crud.AddCompraEstoque(item.Edges.Produto, (item.Quantidade-item.Edges.Produto.QuantNoEstoque)+1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
