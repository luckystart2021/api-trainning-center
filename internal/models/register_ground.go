// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// RegisterGround is an object representing the database table.
type RegisterGround struct {
	ID           int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ClassID      int       `boil:"class_id" json:"class_id" toml:"class_id" yaml:"class_id"`
	StartDate    time.Time `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	TeacherID    int       `boil:"teacher_id" json:"teacher_id" toml:"teacher_id" yaml:"teacher_id"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	GroundNumber string    `boil:"ground_number" json:"ground_number" toml:"ground_number" yaml:"ground_number"`
	EndDate      time.Time `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`

	R *registerGroundR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L registerGroundL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RegisterGroundColumns = struct {
	ID           string
	ClassID      string
	StartDate    string
	TeacherID    string
	CreatedAt    string
	UpdatedAt    string
	GroundNumber string
	EndDate      string
}{
	ID:           "id",
	ClassID:      "class_id",
	StartDate:    "start_date",
	TeacherID:    "teacher_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	GroundNumber: "ground_number",
	EndDate:      "end_date",
}

// Generated where

var RegisterGroundWhere = struct {
	ID           whereHelperint
	ClassID      whereHelperint
	StartDate    whereHelpertime_Time
	TeacherID    whereHelperint
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
	GroundNumber whereHelperstring
	EndDate      whereHelpertime_Time
}{
	ID:           whereHelperint{field: "\"register_ground\".\"id\""},
	ClassID:      whereHelperint{field: "\"register_ground\".\"class_id\""},
	StartDate:    whereHelpertime_Time{field: "\"register_ground\".\"start_date\""},
	TeacherID:    whereHelperint{field: "\"register_ground\".\"teacher_id\""},
	CreatedAt:    whereHelpertime_Time{field: "\"register_ground\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"register_ground\".\"updated_at\""},
	GroundNumber: whereHelperstring{field: "\"register_ground\".\"ground_number\""},
	EndDate:      whereHelpertime_Time{field: "\"register_ground\".\"end_date\""},
}

// RegisterGroundRels is where relationship names are stored.
var RegisterGroundRels = struct {
}{}

// registerGroundR is where relationships are stored.
type registerGroundR struct {
}

// NewStruct creates a new relationship struct
func (*registerGroundR) NewStruct() *registerGroundR {
	return &registerGroundR{}
}

// registerGroundL is where Load methods for each relationship are stored.
type registerGroundL struct{}

var (
	registerGroundAllColumns            = []string{"id", "class_id", "start_date", "teacher_id", "created_at", "updated_at", "ground_number", "end_date"}
	registerGroundColumnsWithoutDefault = []string{"class_id", "start_date", "teacher_id", "ground_number", "end_date"}
	registerGroundColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	registerGroundPrimaryKeyColumns     = []string{"id"}
)

type (
	// RegisterGroundSlice is an alias for a slice of pointers to RegisterGround.
	// This should generally be used opposed to []RegisterGround.
	RegisterGroundSlice []*RegisterGround
	// RegisterGroundHook is the signature for custom RegisterGround hook methods
	RegisterGroundHook func(context.Context, boil.ContextExecutor, *RegisterGround) error

	registerGroundQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	registerGroundType                 = reflect.TypeOf(&RegisterGround{})
	registerGroundMapping              = queries.MakeStructMapping(registerGroundType)
	registerGroundPrimaryKeyMapping, _ = queries.BindMapping(registerGroundType, registerGroundMapping, registerGroundPrimaryKeyColumns)
	registerGroundInsertCacheMut       sync.RWMutex
	registerGroundInsertCache          = make(map[string]insertCache)
	registerGroundUpdateCacheMut       sync.RWMutex
	registerGroundUpdateCache          = make(map[string]updateCache)
	registerGroundUpsertCacheMut       sync.RWMutex
	registerGroundUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var registerGroundBeforeInsertHooks []RegisterGroundHook
var registerGroundBeforeUpdateHooks []RegisterGroundHook
var registerGroundBeforeDeleteHooks []RegisterGroundHook
var registerGroundBeforeUpsertHooks []RegisterGroundHook

var registerGroundAfterInsertHooks []RegisterGroundHook
var registerGroundAfterSelectHooks []RegisterGroundHook
var registerGroundAfterUpdateHooks []RegisterGroundHook
var registerGroundAfterDeleteHooks []RegisterGroundHook
var registerGroundAfterUpsertHooks []RegisterGroundHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RegisterGround) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RegisterGround) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RegisterGround) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RegisterGround) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RegisterGround) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RegisterGround) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RegisterGround) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RegisterGround) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RegisterGround) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range registerGroundAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRegisterGroundHook registers your hook function for all future operations.
func AddRegisterGroundHook(hookPoint boil.HookPoint, registerGroundHook RegisterGroundHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		registerGroundBeforeInsertHooks = append(registerGroundBeforeInsertHooks, registerGroundHook)
	case boil.BeforeUpdateHook:
		registerGroundBeforeUpdateHooks = append(registerGroundBeforeUpdateHooks, registerGroundHook)
	case boil.BeforeDeleteHook:
		registerGroundBeforeDeleteHooks = append(registerGroundBeforeDeleteHooks, registerGroundHook)
	case boil.BeforeUpsertHook:
		registerGroundBeforeUpsertHooks = append(registerGroundBeforeUpsertHooks, registerGroundHook)
	case boil.AfterInsertHook:
		registerGroundAfterInsertHooks = append(registerGroundAfterInsertHooks, registerGroundHook)
	case boil.AfterSelectHook:
		registerGroundAfterSelectHooks = append(registerGroundAfterSelectHooks, registerGroundHook)
	case boil.AfterUpdateHook:
		registerGroundAfterUpdateHooks = append(registerGroundAfterUpdateHooks, registerGroundHook)
	case boil.AfterDeleteHook:
		registerGroundAfterDeleteHooks = append(registerGroundAfterDeleteHooks, registerGroundHook)
	case boil.AfterUpsertHook:
		registerGroundAfterUpsertHooks = append(registerGroundAfterUpsertHooks, registerGroundHook)
	}
}

// One returns a single registerGround record from the query.
func (q registerGroundQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RegisterGround, error) {
	o := &RegisterGround{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for register_ground")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RegisterGround records from the query.
func (q registerGroundQuery) All(ctx context.Context, exec boil.ContextExecutor) (RegisterGroundSlice, error) {
	var o []*RegisterGround

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RegisterGround slice")
	}

	if len(registerGroundAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RegisterGround records in the query.
func (q registerGroundQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count register_ground rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q registerGroundQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if register_ground exists")
	}

	return count > 0, nil
}

// RegisterGrounds retrieves all the records using an executor.
func RegisterGrounds(mods ...qm.QueryMod) registerGroundQuery {
	mods = append(mods, qm.From("\"register_ground\""))
	return registerGroundQuery{NewQuery(mods...)}
}

// FindRegisterGround retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRegisterGround(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RegisterGround, error) {
	registerGroundObj := &RegisterGround{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"register_ground\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, registerGroundObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from register_ground")
	}

	return registerGroundObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RegisterGround) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no register_ground provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(registerGroundColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	registerGroundInsertCacheMut.RLock()
	cache, cached := registerGroundInsertCache[key]
	registerGroundInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			registerGroundAllColumns,
			registerGroundColumnsWithDefault,
			registerGroundColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(registerGroundType, registerGroundMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(registerGroundType, registerGroundMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"register_ground\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"register_ground\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into register_ground")
	}

	if !cached {
		registerGroundInsertCacheMut.Lock()
		registerGroundInsertCache[key] = cache
		registerGroundInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RegisterGround.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RegisterGround) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	registerGroundUpdateCacheMut.RLock()
	cache, cached := registerGroundUpdateCache[key]
	registerGroundUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			registerGroundAllColumns,
			registerGroundPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update register_ground, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"register_ground\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, registerGroundPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(registerGroundType, registerGroundMapping, append(wl, registerGroundPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update register_ground row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for register_ground")
	}

	if !cached {
		registerGroundUpdateCacheMut.Lock()
		registerGroundUpdateCache[key] = cache
		registerGroundUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q registerGroundQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for register_ground")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for register_ground")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RegisterGroundSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerGroundPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"register_ground\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, registerGroundPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in registerGround slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all registerGround")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RegisterGround) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no register_ground provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(registerGroundColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	registerGroundUpsertCacheMut.RLock()
	cache, cached := registerGroundUpsertCache[key]
	registerGroundUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			registerGroundAllColumns,
			registerGroundColumnsWithDefault,
			registerGroundColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			registerGroundAllColumns,
			registerGroundPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert register_ground, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(registerGroundPrimaryKeyColumns))
			copy(conflict, registerGroundPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"register_ground\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(registerGroundType, registerGroundMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(registerGroundType, registerGroundMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert register_ground")
	}

	if !cached {
		registerGroundUpsertCacheMut.Lock()
		registerGroundUpsertCache[key] = cache
		registerGroundUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RegisterGround record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RegisterGround) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RegisterGround provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), registerGroundPrimaryKeyMapping)
	sql := "DELETE FROM \"register_ground\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from register_ground")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for register_ground")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q registerGroundQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no registerGroundQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from register_ground")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for register_ground")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RegisterGroundSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(registerGroundBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerGroundPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"register_ground\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, registerGroundPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from registerGround slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for register_ground")
	}

	if len(registerGroundAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RegisterGround) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRegisterGround(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RegisterGroundSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RegisterGroundSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), registerGroundPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"register_ground\".* FROM \"register_ground\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, registerGroundPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RegisterGroundSlice")
	}

	*o = slice

	return nil
}

// RegisterGroundExists checks if the RegisterGround row exists.
func RegisterGroundExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"register_ground\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if register_ground exists")
	}

	return exists, nil
}
