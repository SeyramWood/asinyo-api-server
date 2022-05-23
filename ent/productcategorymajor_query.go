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
	"github.com/SeyramWood/ent/predicate"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategorymajor"
	"github.com/SeyramWood/ent/productcategoryminor"
)

// ProductCategoryMajorQuery is the builder for querying ProductCategoryMajor entities.
type ProductCategoryMajorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProductCategoryMajor
	// eager-loading edges.
	withMinors   *ProductCategoryMinorQuery
	withProducts *ProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductCategoryMajorQuery builder.
func (pcmq *ProductCategoryMajorQuery) Where(ps ...predicate.ProductCategoryMajor) *ProductCategoryMajorQuery {
	pcmq.predicates = append(pcmq.predicates, ps...)
	return pcmq
}

// Limit adds a limit step to the query.
func (pcmq *ProductCategoryMajorQuery) Limit(limit int) *ProductCategoryMajorQuery {
	pcmq.limit = &limit
	return pcmq
}

// Offset adds an offset step to the query.
func (pcmq *ProductCategoryMajorQuery) Offset(offset int) *ProductCategoryMajorQuery {
	pcmq.offset = &offset
	return pcmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcmq *ProductCategoryMajorQuery) Unique(unique bool) *ProductCategoryMajorQuery {
	pcmq.unique = &unique
	return pcmq
}

// Order adds an order step to the query.
func (pcmq *ProductCategoryMajorQuery) Order(o ...OrderFunc) *ProductCategoryMajorQuery {
	pcmq.order = append(pcmq.order, o...)
	return pcmq
}

// QueryMinors chains the current query on the "minors" edge.
func (pcmq *ProductCategoryMajorQuery) QueryMinors() *ProductCategoryMinorQuery {
	query := &ProductCategoryMinorQuery{config: pcmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productcategorymajor.Table, productcategorymajor.FieldID, selector),
			sqlgraph.To(productcategoryminor.Table, productcategoryminor.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, productcategorymajor.MinorsTable, productcategorymajor.MinorsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pcmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProducts chains the current query on the "products" edge.
func (pcmq *ProductCategoryMajorQuery) QueryProducts() *ProductQuery {
	query := &ProductQuery{config: pcmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productcategorymajor.Table, productcategorymajor.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, productcategorymajor.ProductsTable, productcategorymajor.ProductsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pcmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductCategoryMajor entity from the query.
// Returns a *NotFoundError when no ProductCategoryMajor was found.
func (pcmq *ProductCategoryMajorQuery) First(ctx context.Context) (*ProductCategoryMajor, error) {
	nodes, err := pcmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productcategorymajor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) FirstX(ctx context.Context) *ProductCategoryMajor {
	node, err := pcmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductCategoryMajor ID from the query.
// Returns a *NotFoundError when no ProductCategoryMajor ID was found.
func (pcmq *ProductCategoryMajorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productcategorymajor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) FirstIDX(ctx context.Context) int {
	id, err := pcmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductCategoryMajor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductCategoryMajor entity is found.
// Returns a *NotFoundError when no ProductCategoryMajor entities are found.
func (pcmq *ProductCategoryMajorQuery) Only(ctx context.Context) (*ProductCategoryMajor, error) {
	nodes, err := pcmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productcategorymajor.Label}
	default:
		return nil, &NotSingularError{productcategorymajor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) OnlyX(ctx context.Context) *ProductCategoryMajor {
	node, err := pcmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductCategoryMajor ID in the query.
// Returns a *NotSingularError when more than one ProductCategoryMajor ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcmq *ProductCategoryMajorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = &NotSingularError{productcategorymajor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) OnlyIDX(ctx context.Context) int {
	id, err := pcmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductCategoryMajors.
func (pcmq *ProductCategoryMajorQuery) All(ctx context.Context) ([]*ProductCategoryMajor, error) {
	if err := pcmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pcmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) AllX(ctx context.Context) []*ProductCategoryMajor {
	nodes, err := pcmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductCategoryMajor IDs.
func (pcmq *ProductCategoryMajorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pcmq.Select(productcategorymajor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) IDsX(ctx context.Context) []int {
	ids, err := pcmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcmq *ProductCategoryMajorQuery) Count(ctx context.Context) (int, error) {
	if err := pcmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pcmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) CountX(ctx context.Context) int {
	count, err := pcmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcmq *ProductCategoryMajorQuery) Exist(ctx context.Context) (bool, error) {
	if err := pcmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pcmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pcmq *ProductCategoryMajorQuery) ExistX(ctx context.Context) bool {
	exist, err := pcmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductCategoryMajorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcmq *ProductCategoryMajorQuery) Clone() *ProductCategoryMajorQuery {
	if pcmq == nil {
		return nil
	}
	return &ProductCategoryMajorQuery{
		config:       pcmq.config,
		limit:        pcmq.limit,
		offset:       pcmq.offset,
		order:        append([]OrderFunc{}, pcmq.order...),
		predicates:   append([]predicate.ProductCategoryMajor{}, pcmq.predicates...),
		withMinors:   pcmq.withMinors.Clone(),
		withProducts: pcmq.withProducts.Clone(),
		// clone intermediate query.
		sql:    pcmq.sql.Clone(),
		path:   pcmq.path,
		unique: pcmq.unique,
	}
}

// WithMinors tells the query-builder to eager-load the nodes that are connected to
// the "minors" edge. The optional arguments are used to configure the query builder of the edge.
func (pcmq *ProductCategoryMajorQuery) WithMinors(opts ...func(*ProductCategoryMinorQuery)) *ProductCategoryMajorQuery {
	query := &ProductCategoryMinorQuery{config: pcmq.config}
	for _, opt := range opts {
		opt(query)
	}
	pcmq.withMinors = query
	return pcmq
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (pcmq *ProductCategoryMajorQuery) WithProducts(opts ...func(*ProductQuery)) *ProductCategoryMajorQuery {
	query := &ProductQuery{config: pcmq.config}
	for _, opt := range opts {
		opt(query)
	}
	pcmq.withProducts = query
	return pcmq
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
//	client.ProductCategoryMajor.Query().
//		GroupBy(productcategorymajor.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pcmq *ProductCategoryMajorQuery) GroupBy(field string, fields ...string) *ProductCategoryMajorGroupBy {
	group := &ProductCategoryMajorGroupBy{config: pcmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pcmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pcmq.sqlQuery(ctx), nil
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
//	client.ProductCategoryMajor.Query().
//		Select(productcategorymajor.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (pcmq *ProductCategoryMajorQuery) Select(fields ...string) *ProductCategoryMajorSelect {
	pcmq.fields = append(pcmq.fields, fields...)
	return &ProductCategoryMajorSelect{ProductCategoryMajorQuery: pcmq}
}

func (pcmq *ProductCategoryMajorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pcmq.fields {
		if !productcategorymajor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcmq.path != nil {
		prev, err := pcmq.path(ctx)
		if err != nil {
			return err
		}
		pcmq.sql = prev
	}
	return nil
}

func (pcmq *ProductCategoryMajorQuery) sqlAll(ctx context.Context) ([]*ProductCategoryMajor, error) {
	var (
		nodes       = []*ProductCategoryMajor{}
		_spec       = pcmq.querySpec()
		loadedTypes = [2]bool{
			pcmq.withMinors != nil,
			pcmq.withProducts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProductCategoryMajor{config: pcmq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pcmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pcmq.withMinors; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ProductCategoryMajor, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Minors = []*ProductCategoryMinor{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*ProductCategoryMajor)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   productcategorymajor.MinorsTable,
				Columns: productcategorymajor.MinorsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(productcategorymajor.MinorsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, pcmq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "minors": %w`, err)
		}
		query.Where(productcategoryminor.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "minors" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Minors = append(nodes[i].Edges.Minors, n)
			}
		}
	}

	if query := pcmq.withProducts; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ProductCategoryMajor, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Products = []*Product{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*ProductCategoryMajor)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   productcategorymajor.ProductsTable,
				Columns: productcategorymajor.ProductsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(productcategorymajor.ProductsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, pcmq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "products": %w`, err)
		}
		query.Where(product.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "products" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Products = append(nodes[i].Edges.Products, n)
			}
		}
	}

	return nodes, nil
}

func (pcmq *ProductCategoryMajorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcmq.querySpec()
	_spec.Node.Columns = pcmq.fields
	if len(pcmq.fields) > 0 {
		_spec.Unique = pcmq.unique != nil && *pcmq.unique
	}
	return sqlgraph.CountNodes(ctx, pcmq.driver, _spec)
}

func (pcmq *ProductCategoryMajorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pcmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pcmq *ProductCategoryMajorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productcategorymajor.Table,
			Columns: productcategorymajor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productcategorymajor.FieldID,
			},
		},
		From:   pcmq.sql,
		Unique: true,
	}
	if unique := pcmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pcmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productcategorymajor.FieldID)
		for i := range fields {
			if fields[i] != productcategorymajor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcmq *ProductCategoryMajorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcmq.driver.Dialect())
	t1 := builder.Table(productcategorymajor.Table)
	columns := pcmq.fields
	if len(columns) == 0 {
		columns = productcategorymajor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcmq.sql != nil {
		selector = pcmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcmq.unique != nil && *pcmq.unique {
		selector.Distinct()
	}
	for _, p := range pcmq.predicates {
		p(selector)
	}
	for _, p := range pcmq.order {
		p(selector)
	}
	if offset := pcmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductCategoryMajorGroupBy is the group-by builder for ProductCategoryMajor entities.
type ProductCategoryMajorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcmgb *ProductCategoryMajorGroupBy) Aggregate(fns ...AggregateFunc) *ProductCategoryMajorGroupBy {
	pcmgb.fns = append(pcmgb.fns, fns...)
	return pcmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pcmgb *ProductCategoryMajorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pcmgb.path(ctx)
	if err != nil {
		return err
	}
	pcmgb.sql = query
	return pcmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pcmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pcmgb.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pcmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) StringsX(ctx context.Context) []string {
	v, err := pcmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) StringX(ctx context.Context) string {
	v, err := pcmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pcmgb.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pcmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) IntsX(ctx context.Context) []int {
	v, err := pcmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) IntX(ctx context.Context) int {
	v, err := pcmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcmgb.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pcmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pcmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pcmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pcmgb.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pcmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pcmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcmgb *ProductCategoryMajorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcmgb *ProductCategoryMajorGroupBy) BoolX(ctx context.Context) bool {
	v, err := pcmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcmgb *ProductCategoryMajorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pcmgb.fields {
		if !productcategorymajor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pcmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pcmgb *ProductCategoryMajorGroupBy) sqlQuery() *sql.Selector {
	selector := pcmgb.sql.Select()
	aggregation := make([]string, 0, len(pcmgb.fns))
	for _, fn := range pcmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pcmgb.fields)+len(pcmgb.fns))
		for _, f := range pcmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pcmgb.fields...)...)
}

// ProductCategoryMajorSelect is the builder for selecting fields of ProductCategoryMajor entities.
type ProductCategoryMajorSelect struct {
	*ProductCategoryMajorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pcms *ProductCategoryMajorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pcms.prepareQuery(ctx); err != nil {
		return err
	}
	pcms.sql = pcms.ProductCategoryMajorQuery.sqlQuery(ctx)
	return pcms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pcms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pcms.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pcms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) StringsX(ctx context.Context) []string {
	v, err := pcms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) StringX(ctx context.Context) string {
	v, err := pcms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pcms.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pcms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) IntsX(ctx context.Context) []int {
	v, err := pcms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) IntX(ctx context.Context) int {
	v, err := pcms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcms.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pcms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pcms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) Float64X(ctx context.Context) float64 {
	v, err := pcms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pcms.fields) > 1 {
		return nil, errors.New("ent: ProductCategoryMajorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pcms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) BoolsX(ctx context.Context) []bool {
	v, err := pcms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pcms *ProductCategoryMajorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productcategorymajor.Label}
	default:
		err = fmt.Errorf("ent: ProductCategoryMajorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcms *ProductCategoryMajorSelect) BoolX(ctx context.Context) bool {
	v, err := pcms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcms *ProductCategoryMajorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pcms.sql.Query()
	if err := pcms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}