// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/pickupstation"
	"github.com/SeyramWood/ent/predicate"
)

// PickupStationQuery is the builder for querying PickupStation entities.
type PickupStationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PickupStation
	// eager-loading edges.
	withOrders *OrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PickupStationQuery builder.
func (psq *PickupStationQuery) Where(ps ...predicate.PickupStation) *PickupStationQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit adds a limit step to the query.
func (psq *PickupStationQuery) Limit(limit int) *PickupStationQuery {
	psq.limit = &limit
	return psq
}

// Offset adds an offset step to the query.
func (psq *PickupStationQuery) Offset(offset int) *PickupStationQuery {
	psq.offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *PickupStationQuery) Unique(unique bool) *PickupStationQuery {
	psq.unique = &unique
	return psq
}

// Order adds an order step to the query.
func (psq *PickupStationQuery) Order(o ...OrderFunc) *PickupStationQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// QueryOrders chains the current query on the "orders" edge.
func (psq *PickupStationQuery) QueryOrders() *OrderQuery {
	query := &OrderQuery{config: psq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pickupstation.Table, pickupstation.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pickupstation.OrdersTable, pickupstation.OrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PickupStation entity from the query.
// Returns a *NotFoundError when no PickupStation was found.
func (psq *PickupStationQuery) First(ctx context.Context) (*PickupStation, error) {
	nodes, err := psq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pickupstation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *PickupStationQuery) FirstX(ctx context.Context) *PickupStation {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PickupStation ID from the query.
// Returns a *NotFoundError when no PickupStation ID was found.
func (psq *PickupStationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pickupstation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *PickupStationQuery) FirstIDX(ctx context.Context) int {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PickupStation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PickupStation entity is found.
// Returns a *NotFoundError when no PickupStation entities are found.
func (psq *PickupStationQuery) Only(ctx context.Context) (*PickupStation, error) {
	nodes, err := psq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pickupstation.Label}
	default:
		return nil, &NotSingularError{pickupstation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *PickupStationQuery) OnlyX(ctx context.Context) *PickupStation {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PickupStation ID in the query.
// Returns a *NotSingularError when more than one PickupStation ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *PickupStationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = &NotSingularError{pickupstation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *PickupStationQuery) OnlyIDX(ctx context.Context) int {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PickupStations.
func (psq *PickupStationQuery) All(ctx context.Context) ([]*PickupStation, error) {
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return psq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (psq *PickupStationQuery) AllX(ctx context.Context) []*PickupStation {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PickupStation IDs.
func (psq *PickupStationQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := psq.Select(pickupstation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *PickupStationQuery) IDsX(ctx context.Context) []int {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *PickupStationQuery) Count(ctx context.Context) (int, error) {
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return psq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (psq *PickupStationQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *PickupStationQuery) Exist(ctx context.Context) (bool, error) {
	if err := psq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return psq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *PickupStationQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PickupStationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *PickupStationQuery) Clone() *PickupStationQuery {
	if psq == nil {
		return nil
	}
	return &PickupStationQuery{
		config:     psq.config,
		limit:      psq.limit,
		offset:     psq.offset,
		order:      append([]OrderFunc{}, psq.order...),
		predicates: append([]predicate.PickupStation{}, psq.predicates...),
		withOrders: psq.withOrders.Clone(),
		// clone intermediate query.
		sql:    psq.sql.Clone(),
		path:   psq.path,
		unique: psq.unique,
	}
}

// WithOrders tells the query-builder to eager-load the nodes that are connected to
// the "orders" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *PickupStationQuery) WithOrders(opts ...func(*OrderQuery)) *PickupStationQuery {
	query := &OrderQuery{config: psq.config}
	for _, opt := range opts {
		opt(query)
	}
	psq.withOrders = query
	return psq
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
//	client.PickupStation.Query().
//		GroupBy(pickupstation.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (psq *PickupStationQuery) GroupBy(field string, fields ...string) *PickupStationGroupBy {
	group := &PickupStationGroupBy{config: psq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return psq.sqlQuery(ctx), nil
	}
	return group
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
//	client.PickupStation.Query().
//		Select(pickupstation.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (psq *PickupStationQuery) Select(fields ...string) *PickupStationSelect {
	psq.fields = append(psq.fields, fields...)
	return &PickupStationSelect{PickupStationQuery: psq}
}

func (psq *PickupStationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range psq.fields {
		if !pickupstation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *PickupStationQuery) sqlAll(ctx context.Context) ([]*PickupStation, error) {
	var (
		nodes       = []*PickupStation{}
		_spec       = psq.querySpec()
		loadedTypes = [1]bool{
			psq.withOrders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PickupStation{config: psq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := psq.withOrders; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*PickupStation)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Orders = []*Order{}
		}
		query.withFKs = true
		query.Where(predicate.Order(func(s *sql.Selector) {
			s.Where(sql.InValues(pickupstation.OrdersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.pickup_station_orders
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "pickup_station_orders" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "pickup_station_orders" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Orders = append(node.Edges.Orders, n)
		}
	}

	return nodes, nil
}

func (psq *PickupStationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	_spec.Node.Columns = psq.fields
	if len(psq.fields) > 0 {
		_spec.Unique = psq.unique != nil && *psq.unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *PickupStationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := psq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (psq *PickupStationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pickupstation.Table,
			Columns: pickupstation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pickupstation.FieldID,
			},
		},
		From:   psq.sql,
		Unique: true,
	}
	if unique := psq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := psq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pickupstation.FieldID)
		for i := range fields {
			if fields[i] != pickupstation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *PickupStationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(pickupstation.Table)
	columns := psq.fields
	if len(columns) == 0 {
		columns = pickupstation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.unique != nil && *psq.unique {
		selector.Distinct()
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PickupStationGroupBy is the group-by builder for PickupStation entities.
type PickupStationGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *PickupStationGroupBy) Aggregate(fns ...AggregateFunc) *PickupStationGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the group-by query and scans the result into the given value.
func (psgb *PickupStationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := psgb.path(ctx)
	if err != nil {
		return err
	}
	psgb.sql = query
	return psgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (psgb *PickupStationGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := psgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("ent: PickupStationGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (psgb *PickupStationGroupBy) StringsX(ctx context.Context) []string {
	v, err := psgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = psgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (psgb *PickupStationGroupBy) StringX(ctx context.Context) string {
	v, err := psgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("ent: PickupStationGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (psgb *PickupStationGroupBy) IntsX(ctx context.Context) []int {
	v, err := psgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = psgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (psgb *PickupStationGroupBy) IntX(ctx context.Context) int {
	v, err := psgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("ent: PickupStationGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (psgb *PickupStationGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := psgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = psgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (psgb *PickupStationGroupBy) Float64X(ctx context.Context) float64 {
	v, err := psgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("ent: PickupStationGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (psgb *PickupStationGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := psgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *PickupStationGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = psgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (psgb *PickupStationGroupBy) BoolX(ctx context.Context) bool {
	v, err := psgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (psgb *PickupStationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range psgb.fields {
		if !pickupstation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := psgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (psgb *PickupStationGroupBy) sqlQuery() *sql.Selector {
	selector := psgb.sql.Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(psgb.fields)+len(psgb.fns))
		for _, f := range psgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(psgb.fields...)...)
}

// PickupStationSelect is the builder for selecting fields of PickupStation entities.
type PickupStationSelect struct {
	*PickupStationQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pss *PickupStationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	pss.sql = pss.PickupStationQuery.sqlQuery(ctx)
	return pss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pss *PickupStationSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("ent: PickupStationSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pss *PickupStationSelect) StringsX(ctx context.Context) []string {
	v, err := pss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pss *PickupStationSelect) StringX(ctx context.Context) string {
	v, err := pss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("ent: PickupStationSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pss *PickupStationSelect) IntsX(ctx context.Context) []int {
	v, err := pss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pss *PickupStationSelect) IntX(ctx context.Context) int {
	v, err := pss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("ent: PickupStationSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pss *PickupStationSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pss *PickupStationSelect) Float64X(ctx context.Context) float64 {
	v, err := pss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("ent: PickupStationSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pss *PickupStationSelect) BoolsX(ctx context.Context) []bool {
	v, err := pss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pss *PickupStationSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pickupstation.Label}
	default:
		err = fmt.Errorf("ent: PickupStationSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pss *PickupStationSelect) BoolX(ctx context.Context) bool {
	v, err := pss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pss *PickupStationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pss.sql.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
