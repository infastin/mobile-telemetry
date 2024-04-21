// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"mobile-telemetry/server/service/repo/db/entimpl/ent/device"
	"mobile-telemetry/server/service/repo/db/entimpl/ent/predicate"
	"mobile-telemetry/server/service/repo/db/entimpl/ent/telemetry"
	"mobile-telemetry/server/service/repo/db/entimpl/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TelemetryQuery is the builder for querying Telemetry entities.
type TelemetryQuery struct {
	config
	ctx        *QueryContext
	order      []telemetry.OrderOption
	inters     []Interceptor
	predicates []predicate.Telemetry
	withUser   *UserQuery
	withDevice *DeviceQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TelemetryQuery builder.
func (tq *TelemetryQuery) Where(ps ...predicate.Telemetry) *TelemetryQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TelemetryQuery) Limit(limit int) *TelemetryQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TelemetryQuery) Offset(offset int) *TelemetryQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TelemetryQuery) Unique(unique bool) *TelemetryQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TelemetryQuery) Order(o ...telemetry.OrderOption) *TelemetryQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryUser chains the current query on the "user" edge.
func (tq *TelemetryQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(telemetry.Table, telemetry.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, telemetry.UserTable, telemetry.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDevice chains the current query on the "device" edge.
func (tq *TelemetryQuery) QueryDevice() *DeviceQuery {
	query := (&DeviceClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(telemetry.Table, telemetry.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, telemetry.DeviceTable, telemetry.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Telemetry entity from the query.
// Returns a *NotFoundError when no Telemetry was found.
func (tq *TelemetryQuery) First(ctx context.Context) (*Telemetry, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{telemetry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TelemetryQuery) FirstX(ctx context.Context) *Telemetry {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Telemetry ID from the query.
// Returns a *NotFoundError when no Telemetry ID was found.
func (tq *TelemetryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{telemetry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TelemetryQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Telemetry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Telemetry entity is found.
// Returns a *NotFoundError when no Telemetry entities are found.
func (tq *TelemetryQuery) Only(ctx context.Context) (*Telemetry, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{telemetry.Label}
	default:
		return nil, &NotSingularError{telemetry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TelemetryQuery) OnlyX(ctx context.Context) *Telemetry {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Telemetry ID in the query.
// Returns a *NotSingularError when more than one Telemetry ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TelemetryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{telemetry.Label}
	default:
		err = &NotSingularError{telemetry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TelemetryQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Telemetries.
func (tq *TelemetryQuery) All(ctx context.Context) ([]*Telemetry, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Telemetry, *TelemetryQuery]()
	return withInterceptors[[]*Telemetry](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TelemetryQuery) AllX(ctx context.Context) []*Telemetry {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Telemetry IDs.
func (tq *TelemetryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(telemetry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TelemetryQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TelemetryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TelemetryQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TelemetryQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TelemetryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TelemetryQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TelemetryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TelemetryQuery) Clone() *TelemetryQuery {
	if tq == nil {
		return nil
	}
	return &TelemetryQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]telemetry.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Telemetry{}, tq.predicates...),
		withUser:   tq.withUser.Clone(),
		withDevice: tq.withDevice.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TelemetryQuery) WithUser(opts ...func(*UserQuery)) *TelemetryQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUser = query
	return tq
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TelemetryQuery) WithDevice(opts ...func(*DeviceQuery)) *TelemetryQuery {
	query := (&DeviceClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withDevice = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppVersion string `json:"app_version,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Telemetry.Query().
//		GroupBy(telemetry.FieldAppVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TelemetryQuery) GroupBy(field string, fields ...string) *TelemetryGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TelemetryGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = telemetry.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppVersion string `json:"app_version,omitempty"`
//	}
//
//	client.Telemetry.Query().
//		Select(telemetry.FieldAppVersion).
//		Scan(ctx, &v)
func (tq *TelemetryQuery) Select(fields ...string) *TelemetrySelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TelemetrySelect{TelemetryQuery: tq}
	sbuild.label = telemetry.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TelemetrySelect configured with the given aggregations.
func (tq *TelemetryQuery) Aggregate(fns ...AggregateFunc) *TelemetrySelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TelemetryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !telemetry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TelemetryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Telemetry, error) {
	var (
		nodes       = []*Telemetry{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withUser != nil,
			tq.withDevice != nil,
		}
	)
	if tq.withUser != nil || tq.withDevice != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, telemetry.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Telemetry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Telemetry{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withUser; query != nil {
		if err := tq.loadUser(ctx, query, nodes, nil,
			func(n *Telemetry, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withDevice; query != nil {
		if err := tq.loadDevice(ctx, query, nodes, nil,
			func(n *Telemetry, e *Device) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TelemetryQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Telemetry, init func(*Telemetry), assign func(*Telemetry, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Telemetry)
	for i := range nodes {
		if nodes[i].user_telemetries == nil {
			continue
		}
		fk := *nodes[i].user_telemetries
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_telemetries" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TelemetryQuery) loadDevice(ctx context.Context, query *DeviceQuery, nodes []*Telemetry, init func(*Telemetry), assign func(*Telemetry, *Device)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Telemetry)
	for i := range nodes {
		if nodes[i].device_telemetries == nil {
			continue
		}
		fk := *nodes[i].device_telemetries
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(device.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_telemetries" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TelemetryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TelemetryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(telemetry.Table, telemetry.Columns, sqlgraph.NewFieldSpec(telemetry.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, telemetry.FieldID)
		for i := range fields {
			if fields[i] != telemetry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TelemetryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(telemetry.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = telemetry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TelemetryGroupBy is the group-by builder for Telemetry entities.
type TelemetryGroupBy struct {
	selector
	build *TelemetryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TelemetryGroupBy) Aggregate(fns ...AggregateFunc) *TelemetryGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TelemetryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TelemetryQuery, *TelemetryGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TelemetryGroupBy) sqlScan(ctx context.Context, root *TelemetryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TelemetrySelect is the builder for selecting fields of Telemetry entities.
type TelemetrySelect struct {
	*TelemetryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TelemetrySelect) Aggregate(fns ...AggregateFunc) *TelemetrySelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TelemetrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TelemetryQuery, *TelemetrySelect](ctx, ts.TelemetryQuery, ts, ts.inters, v)
}

func (ts *TelemetrySelect) sqlScan(ctx context.Context, root *TelemetryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
