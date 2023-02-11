// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/agentrequest"
	"github.com/SeyramWood/ent/logistic"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/predicate"
)

// MerchantStoreQuery is the builder for querying MerchantStore entities.
type MerchantStoreQuery struct {
	config
	limit            *int
	offset           *int
	unique           *bool
	order            []OrderFunc
	fields           []string
	predicates       []predicate.MerchantStore
	withMerchant     *MerchantQuery
	withAgent        *AgentQuery
	withLogistics    *LogisticQuery
	withRequests     *AgentRequestQuery
	withOrders       *OrderQuery
	withOrderDetails *OrderDetailQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MerchantStoreQuery builder.
func (msq *MerchantStoreQuery) Where(ps ...predicate.MerchantStore) *MerchantStoreQuery {
	msq.predicates = append(msq.predicates, ps...)
	return msq
}

// Limit adds a limit step to the query.
func (msq *MerchantStoreQuery) Limit(limit int) *MerchantStoreQuery {
	msq.limit = &limit
	return msq
}

// Offset adds an offset step to the query.
func (msq *MerchantStoreQuery) Offset(offset int) *MerchantStoreQuery {
	msq.offset = &offset
	return msq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (msq *MerchantStoreQuery) Unique(unique bool) *MerchantStoreQuery {
	msq.unique = &unique
	return msq
}

// Order adds an order step to the query.
func (msq *MerchantStoreQuery) Order(o ...OrderFunc) *MerchantStoreQuery {
	msq.order = append(msq.order, o...)
	return msq
}

// QueryMerchant chains the current query on the "merchant" edge.
func (msq *MerchantStoreQuery) QueryMerchant() *MerchantQuery {
	query := &MerchantQuery{config: msq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := msq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchantstore.Table, merchantstore.FieldID, selector),
			sqlgraph.To(merchant.Table, merchant.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, merchantstore.MerchantTable, merchantstore.MerchantColumn),
		)
		fromU = sqlgraph.SetNeighbors(msq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAgent chains the current query on the "agent" edge.
func (msq *MerchantStoreQuery) QueryAgent() *AgentQuery {
	query := &AgentQuery{config: msq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := msq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchantstore.Table, merchantstore.FieldID, selector),
			sqlgraph.To(agent.Table, agent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, merchantstore.AgentTable, merchantstore.AgentColumn),
		)
		fromU = sqlgraph.SetNeighbors(msq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLogistics chains the current query on the "logistics" edge.
func (msq *MerchantStoreQuery) QueryLogistics() *LogisticQuery {
	query := &LogisticQuery{config: msq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := msq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchantstore.Table, merchantstore.FieldID, selector),
			sqlgraph.To(logistic.Table, logistic.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, merchantstore.LogisticsTable, merchantstore.LogisticsColumn),
		)
		fromU = sqlgraph.SetNeighbors(msq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRequests chains the current query on the "requests" edge.
func (msq *MerchantStoreQuery) QueryRequests() *AgentRequestQuery {
	query := &AgentRequestQuery{config: msq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := msq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchantstore.Table, merchantstore.FieldID, selector),
			sqlgraph.To(agentrequest.Table, agentrequest.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, merchantstore.RequestsTable, merchantstore.RequestsColumn),
		)
		fromU = sqlgraph.SetNeighbors(msq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrders chains the current query on the "orders" edge.
func (msq *MerchantStoreQuery) QueryOrders() *OrderQuery {
	query := &OrderQuery{config: msq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := msq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchantstore.Table, merchantstore.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, merchantstore.OrdersTable, merchantstore.OrdersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(msq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderDetails chains the current query on the "order_details" edge.
func (msq *MerchantStoreQuery) QueryOrderDetails() *OrderDetailQuery {
	query := &OrderDetailQuery{config: msq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := msq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(merchantstore.Table, merchantstore.FieldID, selector),
			sqlgraph.To(orderdetail.Table, orderdetail.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, merchantstore.OrderDetailsTable, merchantstore.OrderDetailsColumn),
		)
		fromU = sqlgraph.SetNeighbors(msq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MerchantStore entity from the query.
// Returns a *NotFoundError when no MerchantStore was found.
func (msq *MerchantStoreQuery) First(ctx context.Context) (*MerchantStore, error) {
	nodes, err := msq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{merchantstore.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (msq *MerchantStoreQuery) FirstX(ctx context.Context) *MerchantStore {
	node, err := msq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MerchantStore ID from the query.
// Returns a *NotFoundError when no MerchantStore ID was found.
func (msq *MerchantStoreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = msq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{merchantstore.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (msq *MerchantStoreQuery) FirstIDX(ctx context.Context) int {
	id, err := msq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MerchantStore entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MerchantStore entity is found.
// Returns a *NotFoundError when no MerchantStore entities are found.
func (msq *MerchantStoreQuery) Only(ctx context.Context) (*MerchantStore, error) {
	nodes, err := msq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{merchantstore.Label}
	default:
		return nil, &NotSingularError{merchantstore.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (msq *MerchantStoreQuery) OnlyX(ctx context.Context) *MerchantStore {
	node, err := msq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MerchantStore ID in the query.
// Returns a *NotSingularError when more than one MerchantStore ID is found.
// Returns a *NotFoundError when no entities are found.
func (msq *MerchantStoreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = msq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{merchantstore.Label}
	default:
		err = &NotSingularError{merchantstore.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (msq *MerchantStoreQuery) OnlyIDX(ctx context.Context) int {
	id, err := msq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MerchantStores.
func (msq *MerchantStoreQuery) All(ctx context.Context) ([]*MerchantStore, error) {
	if err := msq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return msq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (msq *MerchantStoreQuery) AllX(ctx context.Context) []*MerchantStore {
	nodes, err := msq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MerchantStore IDs.
func (msq *MerchantStoreQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := msq.Select(merchantstore.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (msq *MerchantStoreQuery) IDsX(ctx context.Context) []int {
	ids, err := msq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (msq *MerchantStoreQuery) Count(ctx context.Context) (int, error) {
	if err := msq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return msq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (msq *MerchantStoreQuery) CountX(ctx context.Context) int {
	count, err := msq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (msq *MerchantStoreQuery) Exist(ctx context.Context) (bool, error) {
	if err := msq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return msq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (msq *MerchantStoreQuery) ExistX(ctx context.Context) bool {
	exist, err := msq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MerchantStoreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (msq *MerchantStoreQuery) Clone() *MerchantStoreQuery {
	if msq == nil {
		return nil
	}
	return &MerchantStoreQuery{
		config:           msq.config,
		limit:            msq.limit,
		offset:           msq.offset,
		order:            append([]OrderFunc{}, msq.order...),
		predicates:       append([]predicate.MerchantStore{}, msq.predicates...),
		withMerchant:     msq.withMerchant.Clone(),
		withAgent:        msq.withAgent.Clone(),
		withLogistics:    msq.withLogistics.Clone(),
		withRequests:     msq.withRequests.Clone(),
		withOrders:       msq.withOrders.Clone(),
		withOrderDetails: msq.withOrderDetails.Clone(),
		// clone intermediate query.
		sql:    msq.sql.Clone(),
		path:   msq.path,
		unique: msq.unique,
	}
}

// WithMerchant tells the query-builder to eager-load the nodes that are connected to
// the "merchant" edge. The optional arguments are used to configure the query builder of the edge.
func (msq *MerchantStoreQuery) WithMerchant(opts ...func(*MerchantQuery)) *MerchantStoreQuery {
	query := &MerchantQuery{config: msq.config}
	for _, opt := range opts {
		opt(query)
	}
	msq.withMerchant = query
	return msq
}

// WithAgent tells the query-builder to eager-load the nodes that are connected to
// the "agent" edge. The optional arguments are used to configure the query builder of the edge.
func (msq *MerchantStoreQuery) WithAgent(opts ...func(*AgentQuery)) *MerchantStoreQuery {
	query := &AgentQuery{config: msq.config}
	for _, opt := range opts {
		opt(query)
	}
	msq.withAgent = query
	return msq
}

// WithLogistics tells the query-builder to eager-load the nodes that are connected to
// the "logistics" edge. The optional arguments are used to configure the query builder of the edge.
func (msq *MerchantStoreQuery) WithLogistics(opts ...func(*LogisticQuery)) *MerchantStoreQuery {
	query := &LogisticQuery{config: msq.config}
	for _, opt := range opts {
		opt(query)
	}
	msq.withLogistics = query
	return msq
}

// WithRequests tells the query-builder to eager-load the nodes that are connected to
// the "requests" edge. The optional arguments are used to configure the query builder of the edge.
func (msq *MerchantStoreQuery) WithRequests(opts ...func(*AgentRequestQuery)) *MerchantStoreQuery {
	query := &AgentRequestQuery{config: msq.config}
	for _, opt := range opts {
		opt(query)
	}
	msq.withRequests = query
	return msq
}

// WithOrders tells the query-builder to eager-load the nodes that are connected to
// the "orders" edge. The optional arguments are used to configure the query builder of the edge.
func (msq *MerchantStoreQuery) WithOrders(opts ...func(*OrderQuery)) *MerchantStoreQuery {
	query := &OrderQuery{config: msq.config}
	for _, opt := range opts {
		opt(query)
	}
	msq.withOrders = query
	return msq
}

// WithOrderDetails tells the query-builder to eager-load the nodes that are connected to
// the "order_details" edge. The optional arguments are used to configure the query builder of the edge.
func (msq *MerchantStoreQuery) WithOrderDetails(opts ...func(*OrderDetailQuery)) *MerchantStoreQuery {
	query := &OrderDetailQuery{config: msq.config}
	for _, opt := range opts {
		opt(query)
	}
	msq.withOrderDetails = query
	return msq
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
//	client.MerchantStore.Query().
//		GroupBy(merchantstore.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (msq *MerchantStoreQuery) GroupBy(field string, fields ...string) *MerchantStoreGroupBy {
	grbuild := &MerchantStoreGroupBy{config: msq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := msq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return msq.sqlQuery(ctx), nil
	}
	grbuild.label = merchantstore.Label
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
//	client.MerchantStore.Query().
//		Select(merchantstore.FieldCreatedAt).
//		Scan(ctx, &v)
func (msq *MerchantStoreQuery) Select(fields ...string) *MerchantStoreSelect {
	msq.fields = append(msq.fields, fields...)
	selbuild := &MerchantStoreSelect{MerchantStoreQuery: msq}
	selbuild.label = merchantstore.Label
	selbuild.flds, selbuild.scan = &msq.fields, selbuild.Scan
	return selbuild
}

func (msq *MerchantStoreQuery) prepareQuery(ctx context.Context) error {
	for _, f := range msq.fields {
		if !merchantstore.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if msq.path != nil {
		prev, err := msq.path(ctx)
		if err != nil {
			return err
		}
		msq.sql = prev
	}
	return nil
}

func (msq *MerchantStoreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MerchantStore, error) {
	var (
		nodes       = []*MerchantStore{}
		withFKs     = msq.withFKs
		_spec       = msq.querySpec()
		loadedTypes = [6]bool{
			msq.withMerchant != nil,
			msq.withAgent != nil,
			msq.withLogistics != nil,
			msq.withRequests != nil,
			msq.withOrders != nil,
			msq.withOrderDetails != nil,
		}
	)
	if msq.withMerchant != nil || msq.withAgent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, merchantstore.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MerchantStore).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MerchantStore{config: msq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, msq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := msq.withMerchant; query != nil {
		if err := msq.loadMerchant(ctx, query, nodes, nil,
			func(n *MerchantStore, e *Merchant) { n.Edges.Merchant = e }); err != nil {
			return nil, err
		}
	}
	if query := msq.withAgent; query != nil {
		if err := msq.loadAgent(ctx, query, nodes, nil,
			func(n *MerchantStore, e *Agent) { n.Edges.Agent = e }); err != nil {
			return nil, err
		}
	}
	if query := msq.withLogistics; query != nil {
		if err := msq.loadLogistics(ctx, query, nodes,
			func(n *MerchantStore) { n.Edges.Logistics = []*Logistic{} },
			func(n *MerchantStore, e *Logistic) { n.Edges.Logistics = append(n.Edges.Logistics, e) }); err != nil {
			return nil, err
		}
	}
	if query := msq.withRequests; query != nil {
		if err := msq.loadRequests(ctx, query, nodes,
			func(n *MerchantStore) { n.Edges.Requests = []*AgentRequest{} },
			func(n *MerchantStore, e *AgentRequest) { n.Edges.Requests = append(n.Edges.Requests, e) }); err != nil {
			return nil, err
		}
	}
	if query := msq.withOrders; query != nil {
		if err := msq.loadOrders(ctx, query, nodes,
			func(n *MerchantStore) { n.Edges.Orders = []*Order{} },
			func(n *MerchantStore, e *Order) { n.Edges.Orders = append(n.Edges.Orders, e) }); err != nil {
			return nil, err
		}
	}
	if query := msq.withOrderDetails; query != nil {
		if err := msq.loadOrderDetails(ctx, query, nodes,
			func(n *MerchantStore) { n.Edges.OrderDetails = []*OrderDetail{} },
			func(n *MerchantStore, e *OrderDetail) { n.Edges.OrderDetails = append(n.Edges.OrderDetails, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (msq *MerchantStoreQuery) loadMerchant(ctx context.Context, query *MerchantQuery, nodes []*MerchantStore, init func(*MerchantStore), assign func(*MerchantStore, *Merchant)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MerchantStore)
	for i := range nodes {
		if nodes[i].merchant_store == nil {
			continue
		}
		fk := *nodes[i].merchant_store
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(merchant.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "merchant_store" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (msq *MerchantStoreQuery) loadAgent(ctx context.Context, query *AgentQuery, nodes []*MerchantStore, init func(*MerchantStore), assign func(*MerchantStore, *Agent)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MerchantStore)
	for i := range nodes {
		if nodes[i].agent_store == nil {
			continue
		}
		fk := *nodes[i].agent_store
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(agent.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "agent_store" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (msq *MerchantStoreQuery) loadLogistics(ctx context.Context, query *LogisticQuery, nodes []*MerchantStore, init func(*MerchantStore), assign func(*MerchantStore, *Logistic)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*MerchantStore)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Logistic(func(s *sql.Selector) {
		s.Where(sql.InValues(merchantstore.LogisticsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.merchant_store_logistics
		if fk == nil {
			return fmt.Errorf(`foreign-key "merchant_store_logistics" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "merchant_store_logistics" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (msq *MerchantStoreQuery) loadRequests(ctx context.Context, query *AgentRequestQuery, nodes []*MerchantStore, init func(*MerchantStore), assign func(*MerchantStore, *AgentRequest)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*MerchantStore)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AgentRequest(func(s *sql.Selector) {
		s.Where(sql.InValues(merchantstore.RequestsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.merchant_store_requests
		if fk == nil {
			return fmt.Errorf(`foreign-key "merchant_store_requests" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "merchant_store_requests" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (msq *MerchantStoreQuery) loadOrders(ctx context.Context, query *OrderQuery, nodes []*MerchantStore, init func(*MerchantStore), assign func(*MerchantStore, *Order)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*MerchantStore)
	nids := make(map[int]map[*MerchantStore]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(merchantstore.OrdersTable)
		s.Join(joinT).On(s.C(order.FieldID), joinT.C(merchantstore.OrdersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(merchantstore.OrdersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(merchantstore.OrdersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*MerchantStore]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "orders" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (msq *MerchantStoreQuery) loadOrderDetails(ctx context.Context, query *OrderDetailQuery, nodes []*MerchantStore, init func(*MerchantStore), assign func(*MerchantStore, *OrderDetail)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*MerchantStore)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OrderDetail(func(s *sql.Selector) {
		s.Where(sql.InValues(merchantstore.OrderDetailsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.merchant_store_order_details
		if fk == nil {
			return fmt.Errorf(`foreign-key "merchant_store_order_details" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "merchant_store_order_details" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (msq *MerchantStoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := msq.querySpec()
	_spec.Node.Columns = msq.fields
	if len(msq.fields) > 0 {
		_spec.Unique = msq.unique != nil && *msq.unique
	}
	return sqlgraph.CountNodes(ctx, msq.driver, _spec)
}

func (msq *MerchantStoreQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := msq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (msq *MerchantStoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   merchantstore.Table,
			Columns: merchantstore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchantstore.FieldID,
			},
		},
		From:   msq.sql,
		Unique: true,
	}
	if unique := msq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := msq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchantstore.FieldID)
		for i := range fields {
			if fields[i] != merchantstore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := msq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := msq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := msq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := msq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (msq *MerchantStoreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(msq.driver.Dialect())
	t1 := builder.Table(merchantstore.Table)
	columns := msq.fields
	if len(columns) == 0 {
		columns = merchantstore.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if msq.sql != nil {
		selector = msq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if msq.unique != nil && *msq.unique {
		selector.Distinct()
	}
	for _, p := range msq.predicates {
		p(selector)
	}
	for _, p := range msq.order {
		p(selector)
	}
	if offset := msq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := msq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MerchantStoreGroupBy is the group-by builder for MerchantStore entities.
type MerchantStoreGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (msgb *MerchantStoreGroupBy) Aggregate(fns ...AggregateFunc) *MerchantStoreGroupBy {
	msgb.fns = append(msgb.fns, fns...)
	return msgb
}

// Scan applies the group-by query and scans the result into the given value.
func (msgb *MerchantStoreGroupBy) Scan(ctx context.Context, v any) error {
	query, err := msgb.path(ctx)
	if err != nil {
		return err
	}
	msgb.sql = query
	return msgb.sqlScan(ctx, v)
}

func (msgb *MerchantStoreGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range msgb.fields {
		if !merchantstore.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := msgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := msgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (msgb *MerchantStoreGroupBy) sqlQuery() *sql.Selector {
	selector := msgb.sql.Select()
	aggregation := make([]string, 0, len(msgb.fns))
	for _, fn := range msgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(msgb.fields)+len(msgb.fns))
		for _, f := range msgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(msgb.fields...)...)
}

// MerchantStoreSelect is the builder for selecting fields of MerchantStore entities.
type MerchantStoreSelect struct {
	*MerchantStoreQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mss *MerchantStoreSelect) Scan(ctx context.Context, v any) error {
	if err := mss.prepareQuery(ctx); err != nil {
		return err
	}
	mss.sql = mss.MerchantStoreQuery.sqlQuery(ctx)
	return mss.sqlScan(ctx, v)
}

func (mss *MerchantStoreSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := mss.sql.Query()
	if err := mss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
