package handler

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jmrflora/bazarTudao/crud"
	"github.com/jmrflora/bazarTudao/ent"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	crud *crud.Crud
}

func NewHandler(crud *crud.Crud) *Handler {
	return &Handler{crud: crud}
}

func (h *Handler) HandleGetOrdens(c echo.Context) error {
	ordens, err := h.crud.GetAllOrdens()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetOrdensCompletas(c echo.Context) error {
	ordens, err := h.crud.GetOrdensCompletas()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetOrdensIntocadas(c echo.Context) error {
	ordens, err := h.crud.GetOrdensIntocadas()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetOrdensParciais(c echo.Context) error {
	ordens, err := h.crud.GetOrdensParciais()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ordens)
}

func (h *Handler) HandleGetEnvios(c echo.Context) error {
	envios, err := h.crud.GetAllEnvios()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, envios)
}

func (h *Handler) HandleGetItensPorIdOrdem(c echo.Context) error {
	idst := c.Param("id")
	id, err := strconv.Atoi(idst)
	if err != nil {
		return err
	}
	itens, err := h.crud.GetItensPorIdOrdem(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, itens)
}

func (h *Handler) HandleGetProdutos(c echo.Context) error {
	produtos, err := h.crud.GetAllProdutos()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, produtos)
}

func (h *Handler) HandlerGetAllClientes(c echo.Context) error {
	clientes, err := h.crud.GetAllClientes()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, clientes)
}

func (h *Handler) LoadPedidos(c echo.Context) error {
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
		ordem, err := h.crud.GetOrdem(id)
		// não existe ordem
		if err != nil {
			cliente, err = h.crud.GetCliente(record[1])
			// não existe cliente
			if err != nil {
				cliente, err = h.crud.AddCliente(record[2], record[1], record[4], record[9], record[3])
				if err != nil {
					return err
				}
			}
			ordem, err = h.crud.AddOrdem(cliente, id)
			if err != nil {
				return err
			}
		}
		produto, err := h.crud.GetProduto(record[5])
		// não existe produto
		if err != nil {
			produto, err = h.crud.AddProduto(record[5], record[6])
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

		item, err := h.crud.AddItem(produto, ordem, quant, precoU, precoT)
		if err != nil {
			return err
		}

		ordem.Update().SetPrecoDaOrdem(ordem.PrecoDaOrdem + item.PrecoTotal).SaveX(context.Background())

		println("/n")
		fmt.Printf("item: %v\n", item)
		println("ola de novo")

	}
	println("ola")
	return c.JSON(http.StatusOK, "tudo ok")
}

func (h *Handler) TratarPedidosIntocados(c echo.Context) error {
	ordens, err := h.crud.GetOrdensIntocadas()
	println("\n")
	fmt.Printf("ordens: %v\n", ordens)
	if err != nil {
		return err
	}
	for _, ordem := range ordens {
		itens, err := h.crud.GetItens(ordem)
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
				_, err := h.crud.AddEnvio(item)
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
	return c.JSON(http.StatusOK, "tudo ok")
}

func (h *Handler) TratarPedidosParciais(c echo.Context) error {
	ordens, err := h.crud.GetOrdensParciais()
	if err != nil {
		return err
	}
	for _, ordem := range ordens {
		itens, err := h.crud.GetItensNaoEnviadosPorOrdem(ordem)
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
				_, err := h.crud.AddEnvio(item)
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
	return c.JSON(http.StatusOK, "tudo ok")
}

func (h *Handler) ComprarOqueFalta(c echo.Context) error {
	itens, err := h.crud.GetItensNaoEnviados()
	if err != nil {
		return err
	}

	for _, item := range itens {

		fmt.Printf("item: %v\n", item)

		if item.Edges.Produto.QuantNoEstoque < item.Quantidade {
			err = h.crud.AddCompraEstoque(item.Edges.Produto, (item.Quantidade-item.Edges.Produto.QuantNoEstoque)+1)
			if err != nil {
				return err
			}
		}
	}
	return c.JSON(http.StatusOK, "tudo ok")
}
