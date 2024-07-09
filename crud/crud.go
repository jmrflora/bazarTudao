package crud

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/jmrflora/bazarTudao/ent"
	"github.com/jmrflora/bazarTudao/ent/cliente"
	"github.com/jmrflora/bazarTudao/ent/itemordem"
	"github.com/jmrflora/bazarTudao/ent/ordem"
	"github.com/jmrflora/bazarTudao/ent/produto"
)

type Crud struct {
	c *ent.Client
}

func NewCrud(driverName string, dataSourceName string) (*Crud, error) {
	c, err := ent.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Crud{c: c}, nil
}

func (crud *Crud) CloseCrud() {
	crud.c.Close()
}

func (crud *Crud) CreateSchema() error {
	err := crud.c.Schema.Create(context.Background())
	return err
}

func (crud *Crud) AddCliente(nome, email, telefone, endereco, cpf string) (*ent.Cliente, error) {
	create := crud.c.Cliente.Create()

	create.SetNome(nome).SetCpf(cpf).SetEmail(email).SetTelefone(telefone).SetEnderecoEntrega(endereco)
	cliente, err := create.Save(context.Background())
	if err != nil {
		return nil, err
	}
	return cliente, nil
}

func (crud *Crud) GetCliente(email string) (*ent.Cliente, error) {
	cliente, err := crud.c.Cliente.Query().Where(cliente.Email(email)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return cliente, nil
}

func (crud *Crud) GetAllClientes() ([]*ent.Cliente, error) {
	clientes, err := crud.c.Cliente.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return clientes, err
}

func (crud *Crud) AddOrdem(cliente *ent.Cliente, id int) (*ent.Ordem, error) {
	create := crud.c.Ordem.Create()

	create.SetCliente(cliente).SetID(id).SetPrecoDaOrdem(0)
	ordem, err := create.Save(context.Background())
	if err != nil {
		return nil, err
	}
	return ordem, nil
}

func (crud *Crud) AddProduto(sku string, nome string) (*ent.Produto, error) {
	create := crud.c.Produto.Create()
	create.SetSku(sku).SetNome(nome).SetQuantNoEstoque(0)

	produto, err := create.Save(context.Background())
	if err != nil {
		return nil, err
	}

	return produto, nil
}

func (crud *Crud) GetProduto(sku string) (*ent.Produto, error) {
	produto, err := crud.c.Produto.Query().Where(produto.Sku(sku)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return produto, nil
}

func (crud *Crud) GetAllProdutos() ([]*ent.Produto, error) {
	produtos, err := crud.c.Produto.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func (crud *Crud) AddItem(p *ent.Produto, ordem *ent.Ordem, quant int, precoUnitario float64, precoTotal float64) (*ent.ItemOrdem, error) {
	create := crud.c.ItemOrdem.Create()
	create.SetProduto(p).SetQuantidade(quant).SetPrecoUnitario(precoUnitario).SetPrecoTotal(precoTotal).SetOrdem(ordem)

	item, err := create.Save(context.Background())
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (crud *Crud) GetItens(ordem *ent.Ordem) ([]*ent.ItemOrdem, error) {
	itens, err := ordem.QueryItems().WithProduto().All(context.Background())
	if err != nil {
		return nil, err
	}
	return itens, nil
}

func (crud *Crud) GetItensPorIdOrdem(id int) ([]*ent.ItemOrdem, error) {
	itens, err := crud.c.ItemOrdem.Query().WithProduto().Where(itemordem.OrdemID(id)).All(context.Background())
	if err != nil {
		return nil, err
	}
	return itens, nil
}

func (crud *Crud) GetItensNaoEnviados() ([]*ent.ItemOrdem, error) {
	itens, err := crud.c.ItemOrdem.Query().
		Where(itemordem.Not(itemordem.HasEnvio()), itemordem.And(
			itemordem.HasOrdemWith(ordem.StatusEQ(ordem.StatusParcial)))).
		WithProduto().All(context.Background())
	if err != nil {
		return nil, err
	}
	return itens, nil
}

func (crud *Crud) GetItensNaoEnviadosPorOrdem(o *ent.Ordem) ([]*ent.ItemOrdem, error) {
	itens, err := o.QueryItems().WithProduto().
		Where(itemordem.Not(itemordem.HasEnvio())).All(context.Background())
	if err != nil {
		return nil, err
	}
	return itens, nil
}

func (crud *Crud) GetOrdem(id int) (*ent.Ordem, error) {
	ordem, err := crud.c.Ordem.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return ordem, nil
}

func (crud *Crud) GetOrdensIntocadas() ([]*ent.Ordem, error) {
	ordens, err := crud.c.Ordem.Query().Where(ordem.StatusEQ(ordem.StatusIntocada)).Order(ordem.ByDataOrdem(sql.OrderDesc())).WithItems().All(context.Background())
	if err != nil {
		return nil, err
	}

	return ordens, nil
}

func (crud *Crud) GetOrdensParciais() ([]*ent.Ordem, error) {
	ordens, err := crud.c.Ordem.
		Query().Where(ordem.StatusEQ(ordem.StatusParcial)).
		Order(ordem.ByPrecoDaOrdem(sql.OrderDesc())).WithItems().All(context.Background())
	if err != nil {
		return nil, err
	}
	return ordens, nil
}

func (crud *Crud) GetOrdensCompletas() ([]*ent.Ordem, error) {
	ordens, err := crud.c.Ordem.Query().Where(ordem.StatusEQ(ordem.StatusCompleta)).Order(ordem.ByDataOrdem(sql.OrderDesc())).WithItems().All(context.Background())
	if err != nil {
		return nil, err
	}

	return ordens, nil
}

func (crud *Crud) GetAllOrdens() ([]*ent.Ordem, error) {
	ordens, err := crud.c.Ordem.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return ordens, nil
}

func (crud *Crud) AddEnvio(i ...*ent.ItemOrdem) (*ent.Envio, error) {
	create := crud.c.Envio.Create()
	create.AddItens(i...)
	envio, err := create.Save(context.Background())
	if err != nil {
		return nil, err
	}
	return envio, nil
}

func (crud *Crud) GetAllEnvios() ([]*ent.Envio, error) {
	envios, err := crud.c.Envio.Query().WithItens().All(context.Background())
	if err != nil {
		return nil, err
	}
	return envios, nil
}

func (crud *Crud) AddCompraEstoque(p *ent.Produto, quant int) error {
	_, err := crud.c.Stock.Create().
		SetQuantidade(quant).SetProdutos(p).Save(context.Background())
	if err != nil {
		return err
	}
	_, err = p.Update().SetQuantNoEstoque(quant).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (crud *Crud) AddCompraProdutoPorId(id, quant int) error {
	stock, err := crud.c.Stock.Create().SetProdutosID(id).SetQuantidade(quant).Save(context.Background())
	if err != nil {
		return err
	}

	pr, err := stock.QueryProdutos().Only(context.Background())
	if err != nil {
		return err
	}

	_, err = pr.Update().SetQuantNoEstoque(pr.QuantNoEstoque + quant).Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}
