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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// TrainingCost is an object representing the database table.
type TrainingCost struct {
	ID        int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Amount    int64       `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Type      null.String `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	Note      null.String `boil:"note" json:"note,omitempty" toml:"note" yaml:"note,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy string      `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy string      `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`
	CourseID  int         `boil:"course_id" json:"course_id" toml:"course_id" yaml:"course_id"`
	ClassID   int         `boil:"class_id" json:"class_id" toml:"class_id" yaml:"class_id"`

	R *trainingCostR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L trainingCostL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TrainingCostColumns = struct {
	ID        string
	Amount    string
	Type      string
	Note      string
	CreatedAt string
	CreatedBy string
	UpdatedAt string
	UpdatedBy string
	CourseID  string
	ClassID   string
}{
	ID:        "id",
	Amount:    "amount",
	Type:      "type",
	Note:      "note",
	CreatedAt: "created_at",
	CreatedBy: "created_by",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
	CourseID:  "course_id",
	ClassID:   "class_id",
}

// Generated where

var TrainingCostWhere = struct {
	ID        whereHelperint64
	Amount    whereHelperint64
	Type      whereHelpernull_String
	Note      whereHelpernull_String
	CreatedAt whereHelpertime_Time
	CreatedBy whereHelperstring
	UpdatedAt whereHelpertime_Time
	UpdatedBy whereHelperstring
	CourseID  whereHelperint
	ClassID   whereHelperint
}{
	ID:        whereHelperint64{field: "\"training_costs\".\"id\""},
	Amount:    whereHelperint64{field: "\"training_costs\".\"amount\""},
	Type:      whereHelpernull_String{field: "\"training_costs\".\"type\""},
	Note:      whereHelpernull_String{field: "\"training_costs\".\"note\""},
	CreatedAt: whereHelpertime_Time{field: "\"training_costs\".\"created_at\""},
	CreatedBy: whereHelperstring{field: "\"training_costs\".\"created_by\""},
	UpdatedAt: whereHelpertime_Time{field: "\"training_costs\".\"updated_at\""},
	UpdatedBy: whereHelperstring{field: "\"training_costs\".\"updated_by\""},
	CourseID:  whereHelperint{field: "\"training_costs\".\"course_id\""},
	ClassID:   whereHelperint{field: "\"training_costs\".\"class_id\""},
}

// TrainingCostRels is where relationship names are stored.
var TrainingCostRels = struct {
}{}

// trainingCostR is where relationships are stored.
type trainingCostR struct {
}

// NewStruct creates a new relationship struct
func (*trainingCostR) NewStruct() *trainingCostR {
	return &trainingCostR{}
}

// trainingCostL is where Load methods for each relationship are stored.
type trainingCostL struct{}

var (
	trainingCostAllColumns            = []string{"id", "amount", "type", "note", "created_at", "created_by", "updated_at", "updated_by", "course_id", "class_id"}
	trainingCostColumnsWithoutDefault = []string{"amount", "type", "note", "created_by", "updated_by", "course_id", "class_id"}
	trainingCostColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	trainingCostPrimaryKeyColumns     = []string{"id"}
)

type (
	// TrainingCostSlice is an alias for a slice of pointers to TrainingCost.
	// This should generally be used opposed to []TrainingCost.
	TrainingCostSlice []*TrainingCost
	// TrainingCostHook is the signature for custom TrainingCost hook methods
	TrainingCostHook func(context.Context, boil.ContextExecutor, *TrainingCost) error

	trainingCostQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	trainingCostType                 = reflect.TypeOf(&TrainingCost{})
	trainingCostMapping              = queries.MakeStructMapping(trainingCostType)
	trainingCostPrimaryKeyMapping, _ = queries.BindMapping(trainingCostType, trainingCostMapping, trainingCostPrimaryKeyColumns)
	trainingCostInsertCacheMut       sync.RWMutex
	trainingCostInsertCache          = make(map[string]insertCache)
	trainingCostUpdateCacheMut       sync.RWMutex
	trainingCostUpdateCache          = make(map[string]updateCache)
	trainingCostUpsertCacheMut       sync.RWMutex
	trainingCostUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var trainingCostBeforeInsertHooks []TrainingCostHook
var trainingCostBeforeUpdateHooks []TrainingCostHook
var trainingCostBeforeDeleteHooks []TrainingCostHook
var trainingCostBeforeUpsertHooks []TrainingCostHook

var trainingCostAfterInsertHooks []TrainingCostHook
var trainingCostAfterSelectHooks []TrainingCostHook
var trainingCostAfterUpdateHooks []TrainingCostHook
var trainingCostAfterDeleteHooks []TrainingCostHook
var trainingCostAfterUpsertHooks []TrainingCostHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TrainingCost) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TrainingCost) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TrainingCost) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TrainingCost) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TrainingCost) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TrainingCost) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TrainingCost) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TrainingCost) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TrainingCost) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range trainingCostAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTrainingCostHook registers your hook function for all future operations.
func AddTrainingCostHook(hookPoint boil.HookPoint, trainingCostHook TrainingCostHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		trainingCostBeforeInsertHooks = append(trainingCostBeforeInsertHooks, trainingCostHook)
	case boil.BeforeUpdateHook:
		trainingCostBeforeUpdateHooks = append(trainingCostBeforeUpdateHooks, trainingCostHook)
	case boil.BeforeDeleteHook:
		trainingCostBeforeDeleteHooks = append(trainingCostBeforeDeleteHooks, trainingCostHook)
	case boil.BeforeUpsertHook:
		trainingCostBeforeUpsertHooks = append(trainingCostBeforeUpsertHooks, trainingCostHook)
	case boil.AfterInsertHook:
		trainingCostAfterInsertHooks = append(trainingCostAfterInsertHooks, trainingCostHook)
	case boil.AfterSelectHook:
		trainingCostAfterSelectHooks = append(trainingCostAfterSelectHooks, trainingCostHook)
	case boil.AfterUpdateHook:
		trainingCostAfterUpdateHooks = append(trainingCostAfterUpdateHooks, trainingCostHook)
	case boil.AfterDeleteHook:
		trainingCostAfterDeleteHooks = append(trainingCostAfterDeleteHooks, trainingCostHook)
	case boil.AfterUpsertHook:
		trainingCostAfterUpsertHooks = append(trainingCostAfterUpsertHooks, trainingCostHook)
	}
}

// One returns a single trainingCost record from the query.
func (q trainingCostQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TrainingCost, error) {
	o := &TrainingCost{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for training_costs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TrainingCost records from the query.
func (q trainingCostQuery) All(ctx context.Context, exec boil.ContextExecutor) (TrainingCostSlice, error) {
	var o []*TrainingCost

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TrainingCost slice")
	}

	if len(trainingCostAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TrainingCost records in the query.
func (q trainingCostQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count training_costs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q trainingCostQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if training_costs exists")
	}

	return count > 0, nil
}

// TrainingCosts retrieves all the records using an executor.
func TrainingCosts(mods ...qm.QueryMod) trainingCostQuery {
	mods = append(mods, qm.From("\"training_costs\""))
	return trainingCostQuery{NewQuery(mods...)}
}

// FindTrainingCost retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTrainingCost(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*TrainingCost, error) {
	trainingCostObj := &TrainingCost{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"training_costs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, trainingCostObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from training_costs")
	}

	return trainingCostObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TrainingCost) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no training_costs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(trainingCostColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	trainingCostInsertCacheMut.RLock()
	cache, cached := trainingCostInsertCache[key]
	trainingCostInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			trainingCostAllColumns,
			trainingCostColumnsWithDefault,
			trainingCostColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(trainingCostType, trainingCostMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(trainingCostType, trainingCostMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"training_costs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"training_costs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into training_costs")
	}

	if !cached {
		trainingCostInsertCacheMut.Lock()
		trainingCostInsertCache[key] = cache
		trainingCostInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TrainingCost.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TrainingCost) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	trainingCostUpdateCacheMut.RLock()
	cache, cached := trainingCostUpdateCache[key]
	trainingCostUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			trainingCostAllColumns,
			trainingCostPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update training_costs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"training_costs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, trainingCostPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(trainingCostType, trainingCostMapping, append(wl, trainingCostPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update training_costs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for training_costs")
	}

	if !cached {
		trainingCostUpdateCacheMut.Lock()
		trainingCostUpdateCache[key] = cache
		trainingCostUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q trainingCostQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for training_costs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for training_costs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TrainingCostSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingCostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"training_costs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, trainingCostPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in trainingCost slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all trainingCost")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TrainingCost) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no training_costs provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(trainingCostColumnsWithDefault, o)

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

	trainingCostUpsertCacheMut.RLock()
	cache, cached := trainingCostUpsertCache[key]
	trainingCostUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			trainingCostAllColumns,
			trainingCostColumnsWithDefault,
			trainingCostColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			trainingCostAllColumns,
			trainingCostPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert training_costs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(trainingCostPrimaryKeyColumns))
			copy(conflict, trainingCostPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"training_costs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(trainingCostType, trainingCostMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(trainingCostType, trainingCostMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert training_costs")
	}

	if !cached {
		trainingCostUpsertCacheMut.Lock()
		trainingCostUpsertCache[key] = cache
		trainingCostUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TrainingCost record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TrainingCost) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TrainingCost provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), trainingCostPrimaryKeyMapping)
	sql := "DELETE FROM \"training_costs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from training_costs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for training_costs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q trainingCostQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no trainingCostQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from training_costs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for training_costs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TrainingCostSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(trainingCostBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingCostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"training_costs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, trainingCostPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from trainingCost slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for training_costs")
	}

	if len(trainingCostAfterDeleteHooks) != 0 {
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
func (o *TrainingCost) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTrainingCost(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TrainingCostSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TrainingCostSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), trainingCostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"training_costs\".* FROM \"training_costs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, trainingCostPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TrainingCostSlice")
	}

	*o = slice

	return nil
}

// TrainingCostExists checks if the TrainingCost row exists.
func TrainingCostExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"training_costs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if training_costs exists")
	}

	return exists, nil
}
