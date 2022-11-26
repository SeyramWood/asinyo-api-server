// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/businesscustomer"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/predicate"
)

// BusinessCustomerQuery is the builder for querying BusinessCustomer entities.
type BusinessCustomerQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.BusinessCustomer
	withCustomer *CustomerQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BusinessCustomerQuery builder.
func (bcq *BusinessCustomerQuery) Where(ps ...predicate.BusinessCustomer) *BusinessCustomerQuery {
	bcq.predicates = append(bcq.predicates, ps...)
	return bcq
}

// Limit adds a limit step to the query.
func (bcq *BusinessCustomerQuery) Limit(limit int) *BusinessCustomerQuery {
	bcq.limit = &limit
	return bcq
}

// Offset adds an offset step to the query.
func (bcq *BusinessCustomerQuery) Offset(offset int) *BusinessCustomerQuery {
	bcq.offset = &offset
	return bcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcq *BusinessCustomerQuery) Unique(unique bool) *BusinessCustomerQuery {
	bcq.unique = &unique
	return bcq
}

// Order adds an order step to the query.
func (bcq *BusinessCustomerQuery) Order(o ...OrderFunc) *BusinessCustomerQuery {
	bcq.order = append(bcq.order, o...)
	return bcq
}

// QueryCustomer chains the current query on the "customer" edge.
func (bcq *BusinessCustomerQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: bcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businesscustomer.Table, businesscustomer.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, businesscustomer.CustomerTable, businesscustomer.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BusinessCustomer entity from the query.
// Returns a *NotFoundError when no BusinessCustomer was found.
func (bcq *BusinessCustomerQuery) First(ctx context.Context) (*BusinessCustomer, error) {
	nodes, err := bcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{businesscustomer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) FirstX(ctx context.Context) *BusinessCustomer {
	node, err := bcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BusinessCustomer ID from the query.
// Returns a *NotFoundError when no BusinessCustomer ID was found.
func (bcq *BusinessCustomerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{businesscustomer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) FirstIDX(ctx context.Context) int {
	id, err := bcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BusinessCustomer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BusinessCustomer entity is found.
// Returns a *NotFoundError when no BusinessCustomer entities are found.
func (bcq *BusinessCustomerQuery) Only(ctx context.Context) (*BusinessCustomer, error) {
	nodes, err := bcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{businesscustomer.Label}
	default:
		return nil, &NotSingularError{businesscustomer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) OnlyX(ctx context.Context) *BusinessCustomer {
	node, err := bcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BusinessCustomer ID in the query.
// Returns a *NotSingularError when more than one BusinessCustomer ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcq *BusinessCustomerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{businesscustomer.Label}
	default:
		err = &NotSingularError{businesscustomer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) OnlyIDX(ctx context.Context) int {
	id, err := bcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BusinessCustomers.
func (bcq *BusinessCustomerQuery) All(ctx context.Context) ([]*BusinessCustomer, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) AllX(ctx context.Context) []*BusinessCustomer {
	nodes, err := bcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BusinessCustomer IDs.
func (bcq *BusinessCustomerQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bcq.Select(businesscustomer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) IDsX(ctx context.Context) []int {
	ids, err := bcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcq *BusinessCustomerQuery) Count(ctx context.Context) (int, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) CountX(ctx context.Context) int {
	count, err := bcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcq *BusinessCustomerQuery) Exist(ctx context.Context) (bool, error) {
	if err := bcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bcq *BusinessCustomerQuery) ExistX(ctx context.Context) bool {
	exist, err := bcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BusinessCustomerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcq *BusinessCustomerQuery) Clone() *BusinessCustomerQuery {
	if bcq == nil {
		return nil
	}
	return &BusinessCustomerQuery{
		config:       bcq.config,
		limit:        bcq.limit,
		offset:       bcq.offset,
		order:        append([]OrderFunc{}, bcq.order...),
		predicates:   append([]predicate.BusinessCustomer{}, bcq.predicates...),
		withCustomer: bcq.withCustomer.Clone(),
		// clone intermediate query.
		sql:    bcq.sql.Clone(),
		path:   bcq.path,
		unique: bcq.unique,
	}
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BusinessCustomerQuery) WithCustomer(opts ...func(*CustomerQuery)) *BusinessCustomerQuery {
	query := &CustomerQuery{config: bcq.config}
	for _, opt := range opts {
		opt(query)
	}
	bcq.withCustomer = query
	return bcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BusinessCustomer.Query().
//		GroupBy(businesscustomer.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bcq *BusinessCustomerQuery) GroupBy(field string, fields ...string) *BusinessCustomerGroupBy {
	grbuild := &BusinessCustomerGroupBy{config: bcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bcq.sqlQuery(ctx), nil
	}
	grbuild.label = businesscustomer.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.BusinessCustomer.Query().
//		Select(businesscustomer.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (bcq *BusinessCustomerQuery) Select(fields ...string) *BusinessCustomerSelect {
	bcq.fields = append(bcq.fields, fields...)
	selbuild := &BusinessCustomerSelect{BusinessCustomerQuery: bcq}
	selbuild.label = businesscustomer.Label
	selbuild.flds, selbuild.scan = &bcq.fields, selbuild.Scan
	return selbuild
}

func (bcq *BusinessCustomerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bcq.fields {
		if !businesscustomer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bcq.path != nil {
		prev, err := bcq.path(ctx)
		if err != nil {
			return err
		}
		bcq.sql = prev
	}
	return nil
}

func (bcq *BusinessCustomerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BusinessCustomer, error) {
	var (
		nodes       = []*BusinessCustomer{}
		withFKs     = bcq.withFKs
		_spec       = bcq.querySpec()
		loadedTypes = [1]bool{
			bcq.withCustomer != nil,
		}
	)
	if bcq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, businesscustomer.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BusinessCustomer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BusinessCustomer{config: bcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bcq.withCustomer; query != nil {
		if err := bcq.loadCustomer(ctx, query, nodes, nil,
			func(n *BusinessCustomer, e *Customer) { n.Edges.Customer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bcq *BusinessCustomerQuery) loadCustomer(ctx context.Context, query *CustomerQuery, nodes []*BusinessCustomer, init func(*BusinessCustomer), assign func(*BusinessCustomer, *Customer)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BusinessCustomer)
	for i := range nodes {
		if nodes[i].customer_business == nil {
			continue
		}
		fk := *nodes[i].customer_business
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(customer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "customer_business" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bcq *BusinessCustomerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcq.querySpec()
	_spec.Node.Columns = bcq.fields
	if len(bcq.fields) > 0 {
		_spec.Unique = bcq.unique != nil && *bcq.unique
	}
	return sqlgraph.CountNodes(ctx, bcq.driver, _spec)
}

func (bcq *BusinessCustomerQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := bcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (bcq *BusinessCustomerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   businesscustomer.Table,
			Columns: businesscustomer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: businesscustomer.FieldID,
			},
		},
		From:   bcq.sql,
		Unique: true,
	}
	if unique := bcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businesscustomer.FieldID)
		for i := range fields {
			if fields[i] != businesscustomer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcq *BusinessCustomerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcq.driver.Dialect())
	t1 := builder.Table(businesscustomer.Table)
	columns := bcq.fields
	if len(columns) == 0 {
		columns = businesscustomer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcq.sql != nil {
		selector = bcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcq.unique != nil && *bcq.unique {
		selector.Distinct()
	}
	for _, p := range bcq.predicates {
		p(selector)
	}
	for _, p := range bcq.order {
		p(selector)
	}
	if offset := bcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BusinessCustomerGroupBy is the group-by builder for BusinessCustomer entities.
type BusinessCustomerGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcgb *BusinessCustomerGroupBy) Aggregate(fns ...AggregateFunc) *BusinessCustomerGroupBy {
	bcgb.fns = append(bcgb.fns, fns...)
	return bcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bcgb *BusinessCustomerGroupBy) Scan(ctx context.Context, v any) error {
	query, err := bcgb.path(ctx)
	if err != nil {
		return err
	}
	bcgb.sql = query
	return bcgb.sqlScan(ctx, v)
}

func (bcgb *BusinessCustomerGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range bcgb.fields {
		if !businesscustomer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bcgb *BusinessCustomerGroupBy) sqlQuery() *sql.Selector {
	selector := bcgb.sql.Select()
	aggregation := make([]string, 0, len(bcgb.fns))
	for _, fn := range bcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bcgb.fields)+len(bcgb.fns))
		for _, f := range bcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bcgb.fields...)...)
}

// BusinessCustomerSelect is the builder for selecting fields of BusinessCustomer entities.
type BusinessCustomerSelect struct {
	*BusinessCustomerQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bcs *BusinessCustomerSelect) Scan(ctx context.Context, v any) error {
	if err := bcs.prepareQuery(ctx); err != nil {
		return err
	}
	bcs.sql = bcs.BusinessCustomerQuery.sqlQuery(ctx)
	return bcs.sqlScan(ctx, v)
}

func (bcs *BusinessCustomerSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := bcs.sql.Query()
	if err := bcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
