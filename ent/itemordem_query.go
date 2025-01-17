// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jmrflora/bazarTudao/ent/envio"
	"github.com/jmrflora/bazarTudao/ent/itemordem"
	"github.com/jmrflora/bazarTudao/ent/ordem"
	"github.com/jmrflora/bazarTudao/ent/predicate"
	"github.com/jmrflora/bazarTudao/ent/produto"
)

// ItemOrdemQuery is the builder for querying ItemOrdem entities.
type ItemOrdemQuery struct {
	config
	ctx         *QueryContext
	order       []itemordem.OrderOption
	inters      []Interceptor
	predicates  []predicate.ItemOrdem
	withOrdem   *OrdemQuery
	withProduto *ProdutoQuery
	withEnvio   *EnvioQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemOrdemQuery builder.
func (ioq *ItemOrdemQuery) Where(ps ...predicate.ItemOrdem) *ItemOrdemQuery {
	ioq.predicates = append(ioq.predicates, ps...)
	return ioq
}

// Limit the number of records to be returned by this query.
func (ioq *ItemOrdemQuery) Limit(limit int) *ItemOrdemQuery {
	ioq.ctx.Limit = &limit
	return ioq
}

// Offset to start from.
func (ioq *ItemOrdemQuery) Offset(offset int) *ItemOrdemQuery {
	ioq.ctx.Offset = &offset
	return ioq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ioq *ItemOrdemQuery) Unique(unique bool) *ItemOrdemQuery {
	ioq.ctx.Unique = &unique
	return ioq
}

// Order specifies how the records should be ordered.
func (ioq *ItemOrdemQuery) Order(o ...itemordem.OrderOption) *ItemOrdemQuery {
	ioq.order = append(ioq.order, o...)
	return ioq
}

// QueryOrdem chains the current query on the "ordem" edge.
func (ioq *ItemOrdemQuery) QueryOrdem() *OrdemQuery {
	query := (&OrdemClient{config: ioq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ioq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ioq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemordem.Table, itemordem.FieldID, selector),
			sqlgraph.To(ordem.Table, ordem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, itemordem.OrdemTable, itemordem.OrdemColumn),
		)
		fromU = sqlgraph.SetNeighbors(ioq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduto chains the current query on the "produto" edge.
func (ioq *ItemOrdemQuery) QueryProduto() *ProdutoQuery {
	query := (&ProdutoClient{config: ioq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ioq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ioq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemordem.Table, itemordem.FieldID, selector),
			sqlgraph.To(produto.Table, produto.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, itemordem.ProdutoTable, itemordem.ProdutoColumn),
		)
		fromU = sqlgraph.SetNeighbors(ioq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEnvio chains the current query on the "envio" edge.
func (ioq *ItemOrdemQuery) QueryEnvio() *EnvioQuery {
	query := (&EnvioClient{config: ioq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ioq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ioq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemordem.Table, itemordem.FieldID, selector),
			sqlgraph.To(envio.Table, envio.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, itemordem.EnvioTable, itemordem.EnvioColumn),
		)
		fromU = sqlgraph.SetNeighbors(ioq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ItemOrdem entity from the query.
// Returns a *NotFoundError when no ItemOrdem was found.
func (ioq *ItemOrdemQuery) First(ctx context.Context) (*ItemOrdem, error) {
	nodes, err := ioq.Limit(1).All(setContextOp(ctx, ioq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{itemordem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ioq *ItemOrdemQuery) FirstX(ctx context.Context) *ItemOrdem {
	node, err := ioq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ItemOrdem ID from the query.
// Returns a *NotFoundError when no ItemOrdem ID was found.
func (ioq *ItemOrdemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ioq.Limit(1).IDs(setContextOp(ctx, ioq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{itemordem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ioq *ItemOrdemQuery) FirstIDX(ctx context.Context) int {
	id, err := ioq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ItemOrdem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ItemOrdem entity is found.
// Returns a *NotFoundError when no ItemOrdem entities are found.
func (ioq *ItemOrdemQuery) Only(ctx context.Context) (*ItemOrdem, error) {
	nodes, err := ioq.Limit(2).All(setContextOp(ctx, ioq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{itemordem.Label}
	default:
		return nil, &NotSingularError{itemordem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ioq *ItemOrdemQuery) OnlyX(ctx context.Context) *ItemOrdem {
	node, err := ioq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ItemOrdem ID in the query.
// Returns a *NotSingularError when more than one ItemOrdem ID is found.
// Returns a *NotFoundError when no entities are found.
func (ioq *ItemOrdemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ioq.Limit(2).IDs(setContextOp(ctx, ioq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{itemordem.Label}
	default:
		err = &NotSingularError{itemordem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ioq *ItemOrdemQuery) OnlyIDX(ctx context.Context) int {
	id, err := ioq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ItemOrdems.
func (ioq *ItemOrdemQuery) All(ctx context.Context) ([]*ItemOrdem, error) {
	ctx = setContextOp(ctx, ioq.ctx, "All")
	if err := ioq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ItemOrdem, *ItemOrdemQuery]()
	return withInterceptors[[]*ItemOrdem](ctx, ioq, qr, ioq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ioq *ItemOrdemQuery) AllX(ctx context.Context) []*ItemOrdem {
	nodes, err := ioq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ItemOrdem IDs.
func (ioq *ItemOrdemQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ioq.ctx.Unique == nil && ioq.path != nil {
		ioq.Unique(true)
	}
	ctx = setContextOp(ctx, ioq.ctx, "IDs")
	if err = ioq.Select(itemordem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ioq *ItemOrdemQuery) IDsX(ctx context.Context) []int {
	ids, err := ioq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ioq *ItemOrdemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ioq.ctx, "Count")
	if err := ioq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ioq, querierCount[*ItemOrdemQuery](), ioq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ioq *ItemOrdemQuery) CountX(ctx context.Context) int {
	count, err := ioq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ioq *ItemOrdemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ioq.ctx, "Exist")
	switch _, err := ioq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ioq *ItemOrdemQuery) ExistX(ctx context.Context) bool {
	exist, err := ioq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemOrdemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ioq *ItemOrdemQuery) Clone() *ItemOrdemQuery {
	if ioq == nil {
		return nil
	}
	return &ItemOrdemQuery{
		config:      ioq.config,
		ctx:         ioq.ctx.Clone(),
		order:       append([]itemordem.OrderOption{}, ioq.order...),
		inters:      append([]Interceptor{}, ioq.inters...),
		predicates:  append([]predicate.ItemOrdem{}, ioq.predicates...),
		withOrdem:   ioq.withOrdem.Clone(),
		withProduto: ioq.withProduto.Clone(),
		withEnvio:   ioq.withEnvio.Clone(),
		// clone intermediate query.
		sql:  ioq.sql.Clone(),
		path: ioq.path,
	}
}

// WithOrdem tells the query-builder to eager-load the nodes that are connected to
// the "ordem" edge. The optional arguments are used to configure the query builder of the edge.
func (ioq *ItemOrdemQuery) WithOrdem(opts ...func(*OrdemQuery)) *ItemOrdemQuery {
	query := (&OrdemClient{config: ioq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ioq.withOrdem = query
	return ioq
}

// WithProduto tells the query-builder to eager-load the nodes that are connected to
// the "produto" edge. The optional arguments are used to configure the query builder of the edge.
func (ioq *ItemOrdemQuery) WithProduto(opts ...func(*ProdutoQuery)) *ItemOrdemQuery {
	query := (&ProdutoClient{config: ioq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ioq.withProduto = query
	return ioq
}

// WithEnvio tells the query-builder to eager-load the nodes that are connected to
// the "envio" edge. The optional arguments are used to configure the query builder of the edge.
func (ioq *ItemOrdemQuery) WithEnvio(opts ...func(*EnvioQuery)) *ItemOrdemQuery {
	query := (&EnvioClient{config: ioq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ioq.withEnvio = query
	return ioq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Quantidade int `json:"quantidade,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ItemOrdem.Query().
//		GroupBy(itemordem.FieldQuantidade).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ioq *ItemOrdemQuery) GroupBy(field string, fields ...string) *ItemOrdemGroupBy {
	ioq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ItemOrdemGroupBy{build: ioq}
	grbuild.flds = &ioq.ctx.Fields
	grbuild.label = itemordem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Quantidade int `json:"quantidade,omitempty"`
//	}
//
//	client.ItemOrdem.Query().
//		Select(itemordem.FieldQuantidade).
//		Scan(ctx, &v)
func (ioq *ItemOrdemQuery) Select(fields ...string) *ItemOrdemSelect {
	ioq.ctx.Fields = append(ioq.ctx.Fields, fields...)
	sbuild := &ItemOrdemSelect{ItemOrdemQuery: ioq}
	sbuild.label = itemordem.Label
	sbuild.flds, sbuild.scan = &ioq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ItemOrdemSelect configured with the given aggregations.
func (ioq *ItemOrdemQuery) Aggregate(fns ...AggregateFunc) *ItemOrdemSelect {
	return ioq.Select().Aggregate(fns...)
}

func (ioq *ItemOrdemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ioq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ioq); err != nil {
				return err
			}
		}
	}
	for _, f := range ioq.ctx.Fields {
		if !itemordem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ioq.path != nil {
		prev, err := ioq.path(ctx)
		if err != nil {
			return err
		}
		ioq.sql = prev
	}
	return nil
}

func (ioq *ItemOrdemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ItemOrdem, error) {
	var (
		nodes       = []*ItemOrdem{}
		withFKs     = ioq.withFKs
		_spec       = ioq.querySpec()
		loadedTypes = [3]bool{
			ioq.withOrdem != nil,
			ioq.withProduto != nil,
			ioq.withEnvio != nil,
		}
	)
	if ioq.withEnvio != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, itemordem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ItemOrdem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ItemOrdem{config: ioq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ioq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ioq.withOrdem; query != nil {
		if err := ioq.loadOrdem(ctx, query, nodes, nil,
			func(n *ItemOrdem, e *Ordem) { n.Edges.Ordem = e }); err != nil {
			return nil, err
		}
	}
	if query := ioq.withProduto; query != nil {
		if err := ioq.loadProduto(ctx, query, nodes, nil,
			func(n *ItemOrdem, e *Produto) { n.Edges.Produto = e }); err != nil {
			return nil, err
		}
	}
	if query := ioq.withEnvio; query != nil {
		if err := ioq.loadEnvio(ctx, query, nodes, nil,
			func(n *ItemOrdem, e *Envio) { n.Edges.Envio = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ioq *ItemOrdemQuery) loadOrdem(ctx context.Context, query *OrdemQuery, nodes []*ItemOrdem, init func(*ItemOrdem), assign func(*ItemOrdem, *Ordem)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ItemOrdem)
	for i := range nodes {
		fk := nodes[i].OrdemID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ordem.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ordem_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ioq *ItemOrdemQuery) loadProduto(ctx context.Context, query *ProdutoQuery, nodes []*ItemOrdem, init func(*ItemOrdem), assign func(*ItemOrdem, *Produto)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ItemOrdem)
	for i := range nodes {
		fk := nodes[i].ProdutoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(produto.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "produto_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ioq *ItemOrdemQuery) loadEnvio(ctx context.Context, query *EnvioQuery, nodes []*ItemOrdem, init func(*ItemOrdem), assign func(*ItemOrdem, *Envio)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ItemOrdem)
	for i := range nodes {
		if nodes[i].envio_itens == nil {
			continue
		}
		fk := *nodes[i].envio_itens
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(envio.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "envio_itens" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ioq *ItemOrdemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ioq.querySpec()
	_spec.Node.Columns = ioq.ctx.Fields
	if len(ioq.ctx.Fields) > 0 {
		_spec.Unique = ioq.ctx.Unique != nil && *ioq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ioq.driver, _spec)
}

func (ioq *ItemOrdemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(itemordem.Table, itemordem.Columns, sqlgraph.NewFieldSpec(itemordem.FieldID, field.TypeInt))
	_spec.From = ioq.sql
	if unique := ioq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ioq.path != nil {
		_spec.Unique = true
	}
	if fields := ioq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemordem.FieldID)
		for i := range fields {
			if fields[i] != itemordem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ioq.withOrdem != nil {
			_spec.Node.AddColumnOnce(itemordem.FieldOrdemID)
		}
		if ioq.withProduto != nil {
			_spec.Node.AddColumnOnce(itemordem.FieldProdutoID)
		}
	}
	if ps := ioq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ioq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ioq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ioq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ioq *ItemOrdemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ioq.driver.Dialect())
	t1 := builder.Table(itemordem.Table)
	columns := ioq.ctx.Fields
	if len(columns) == 0 {
		columns = itemordem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ioq.sql != nil {
		selector = ioq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ioq.ctx.Unique != nil && *ioq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ioq.predicates {
		p(selector)
	}
	for _, p := range ioq.order {
		p(selector)
	}
	if offset := ioq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ioq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ItemOrdemGroupBy is the group-by builder for ItemOrdem entities.
type ItemOrdemGroupBy struct {
	selector
	build *ItemOrdemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iogb *ItemOrdemGroupBy) Aggregate(fns ...AggregateFunc) *ItemOrdemGroupBy {
	iogb.fns = append(iogb.fns, fns...)
	return iogb
}

// Scan applies the selector query and scans the result into the given value.
func (iogb *ItemOrdemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iogb.build.ctx, "GroupBy")
	if err := iogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemOrdemQuery, *ItemOrdemGroupBy](ctx, iogb.build, iogb, iogb.build.inters, v)
}

func (iogb *ItemOrdemGroupBy) sqlScan(ctx context.Context, root *ItemOrdemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(iogb.fns))
	for _, fn := range iogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*iogb.flds)+len(iogb.fns))
		for _, f := range *iogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*iogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ItemOrdemSelect is the builder for selecting fields of ItemOrdem entities.
type ItemOrdemSelect struct {
	*ItemOrdemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ios *ItemOrdemSelect) Aggregate(fns ...AggregateFunc) *ItemOrdemSelect {
	ios.fns = append(ios.fns, fns...)
	return ios
}

// Scan applies the selector query and scans the result into the given value.
func (ios *ItemOrdemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ios.ctx, "Select")
	if err := ios.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemOrdemQuery, *ItemOrdemSelect](ctx, ios.ItemOrdemQuery, ios, ios.inters, v)
}

func (ios *ItemOrdemSelect) sqlScan(ctx context.Context, root *ItemOrdemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ios.fns))
	for _, fn := range ios.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ios.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ios.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
