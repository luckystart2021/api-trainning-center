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

// RankVehicle is an object representing the database table.
type RankVehicle struct {
	ID             int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name           string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Time           int    `boil:"time" json:"time" toml:"time" yaml:"time"`
	NumberQuestion int    `boil:"number_question" json:"number_question" toml:"number_question" yaml:"number_question"`
	PointPass      int    `boil:"point_pass" json:"point_pass" toml:"point_pass" yaml:"point_pass"`

	R *rankVehicleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rankVehicleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RankVehicleColumns = struct {
	ID             string
	Name           string
	Time           string
	NumberQuestion string
	PointPass      string
}{
	ID:             "id",
	Name:           "name",
	Time:           "time",
	NumberQuestion: "number_question",
	PointPass:      "point_pass",
}

// Generated where

var RankVehicleWhere = struct {
	ID             whereHelperint
	Name           whereHelperstring
	Time           whereHelperint
	NumberQuestion whereHelperint
	PointPass      whereHelperint
}{
	ID:             whereHelperint{field: "\"rank_vehicle\".\"id\""},
	Name:           whereHelperstring{field: "\"rank_vehicle\".\"name\""},
	Time:           whereHelperint{field: "\"rank_vehicle\".\"time\""},
	NumberQuestion: whereHelperint{field: "\"rank_vehicle\".\"number_question\""},
	PointPass:      whereHelperint{field: "\"rank_vehicle\".\"point_pass\""},
}

// RankVehicleRels is where relationship names are stored.
var RankVehicleRels = struct {
	RankTestsuites string
}{
	RankTestsuites: "RankTestsuites",
}

// rankVehicleR is where relationships are stored.
type rankVehicleR struct {
	RankTestsuites TestsuiteSlice
}

// NewStruct creates a new relationship struct
func (*rankVehicleR) NewStruct() *rankVehicleR {
	return &rankVehicleR{}
}

// rankVehicleL is where Load methods for each relationship are stored.
type rankVehicleL struct{}

var (
	rankVehicleAllColumns            = []string{"id", "name", "time", "number_question", "point_pass"}
	rankVehicleColumnsWithoutDefault = []string{"name", "time", "number_question", "point_pass"}
	rankVehicleColumnsWithDefault    = []string{"id"}
	rankVehiclePrimaryKeyColumns     = []string{"id"}
)

type (
	// RankVehicleSlice is an alias for a slice of pointers to RankVehicle.
	// This should generally be used opposed to []RankVehicle.
	RankVehicleSlice []*RankVehicle
	// RankVehicleHook is the signature for custom RankVehicle hook methods
	RankVehicleHook func(context.Context, boil.ContextExecutor, *RankVehicle) error

	rankVehicleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rankVehicleType                 = reflect.TypeOf(&RankVehicle{})
	rankVehicleMapping              = queries.MakeStructMapping(rankVehicleType)
	rankVehiclePrimaryKeyMapping, _ = queries.BindMapping(rankVehicleType, rankVehicleMapping, rankVehiclePrimaryKeyColumns)
	rankVehicleInsertCacheMut       sync.RWMutex
	rankVehicleInsertCache          = make(map[string]insertCache)
	rankVehicleUpdateCacheMut       sync.RWMutex
	rankVehicleUpdateCache          = make(map[string]updateCache)
	rankVehicleUpsertCacheMut       sync.RWMutex
	rankVehicleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var rankVehicleBeforeInsertHooks []RankVehicleHook
var rankVehicleBeforeUpdateHooks []RankVehicleHook
var rankVehicleBeforeDeleteHooks []RankVehicleHook
var rankVehicleBeforeUpsertHooks []RankVehicleHook

var rankVehicleAfterInsertHooks []RankVehicleHook
var rankVehicleAfterSelectHooks []RankVehicleHook
var rankVehicleAfterUpdateHooks []RankVehicleHook
var rankVehicleAfterDeleteHooks []RankVehicleHook
var rankVehicleAfterUpsertHooks []RankVehicleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RankVehicle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RankVehicle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RankVehicle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RankVehicle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RankVehicle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RankVehicle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RankVehicle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RankVehicle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RankVehicle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankVehicleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRankVehicleHook registers your hook function for all future operations.
func AddRankVehicleHook(hookPoint boil.HookPoint, rankVehicleHook RankVehicleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		rankVehicleBeforeInsertHooks = append(rankVehicleBeforeInsertHooks, rankVehicleHook)
	case boil.BeforeUpdateHook:
		rankVehicleBeforeUpdateHooks = append(rankVehicleBeforeUpdateHooks, rankVehicleHook)
	case boil.BeforeDeleteHook:
		rankVehicleBeforeDeleteHooks = append(rankVehicleBeforeDeleteHooks, rankVehicleHook)
	case boil.BeforeUpsertHook:
		rankVehicleBeforeUpsertHooks = append(rankVehicleBeforeUpsertHooks, rankVehicleHook)
	case boil.AfterInsertHook:
		rankVehicleAfterInsertHooks = append(rankVehicleAfterInsertHooks, rankVehicleHook)
	case boil.AfterSelectHook:
		rankVehicleAfterSelectHooks = append(rankVehicleAfterSelectHooks, rankVehicleHook)
	case boil.AfterUpdateHook:
		rankVehicleAfterUpdateHooks = append(rankVehicleAfterUpdateHooks, rankVehicleHook)
	case boil.AfterDeleteHook:
		rankVehicleAfterDeleteHooks = append(rankVehicleAfterDeleteHooks, rankVehicleHook)
	case boil.AfterUpsertHook:
		rankVehicleAfterUpsertHooks = append(rankVehicleAfterUpsertHooks, rankVehicleHook)
	}
}

// One returns a single rankVehicle record from the query.
func (q rankVehicleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RankVehicle, error) {
	o := &RankVehicle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for rank_vehicle")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RankVehicle records from the query.
func (q rankVehicleQuery) All(ctx context.Context, exec boil.ContextExecutor) (RankVehicleSlice, error) {
	var o []*RankVehicle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RankVehicle slice")
	}

	if len(rankVehicleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RankVehicle records in the query.
func (q rankVehicleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count rank_vehicle rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q rankVehicleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if rank_vehicle exists")
	}

	return count > 0, nil
}

// RankTestsuites retrieves all the testsuite's Testsuites with an executor via rank_id column.
func (o *RankVehicle) RankTestsuites(mods ...qm.QueryMod) testsuiteQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"testsuite\".\"rank_id\"=?", o.ID),
	)

	query := Testsuites(queryMods...)
	queries.SetFrom(query.Query, "\"testsuite\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"testsuite\".*"})
	}

	return query
}

// LoadRankTestsuites allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (rankVehicleL) LoadRankTestsuites(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRankVehicle interface{}, mods queries.Applicator) error {
	var slice []*RankVehicle
	var object *RankVehicle

	if singular {
		object = maybeRankVehicle.(*RankVehicle)
	} else {
		slice = *maybeRankVehicle.(*[]*RankVehicle)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &rankVehicleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &rankVehicleR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`testsuite`), qm.WhereIn(`testsuite.rank_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load testsuite")
	}

	var resultSlice []*Testsuite
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice testsuite")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on testsuite")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for testsuite")
	}

	if len(testsuiteAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RankTestsuites = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &testsuiteR{}
			}
			foreign.R.Rank = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RankID {
				local.R.RankTestsuites = append(local.R.RankTestsuites, foreign)
				if foreign.R == nil {
					foreign.R = &testsuiteR{}
				}
				foreign.R.Rank = local
				break
			}
		}
	}

	return nil
}

// AddRankTestsuites adds the given related objects to the existing relationships
// of the rank_vehicle, optionally inserting them as new records.
// Appends related to o.R.RankTestsuites.
// Sets related.R.Rank appropriately.
func (o *RankVehicle) AddRankTestsuites(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Testsuite) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RankID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"testsuite\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"rank_id"}),
				strmangle.WhereClause("\"", "\"", 2, testsuitePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RankID = o.ID
		}
	}

	if o.R == nil {
		o.R = &rankVehicleR{
			RankTestsuites: related,
		}
	} else {
		o.R.RankTestsuites = append(o.R.RankTestsuites, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &testsuiteR{
				Rank: o,
			}
		} else {
			rel.R.Rank = o
		}
	}
	return nil
}

// RankVehicles retrieves all the records using an executor.
func RankVehicles(mods ...qm.QueryMod) rankVehicleQuery {
	mods = append(mods, qm.From("\"rank_vehicle\""))
	return rankVehicleQuery{NewQuery(mods...)}
}

// FindRankVehicle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRankVehicle(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RankVehicle, error) {
	rankVehicleObj := &RankVehicle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"rank_vehicle\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, rankVehicleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from rank_vehicle")
	}

	return rankVehicleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RankVehicle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rank_vehicle provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rankVehicleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rankVehicleInsertCacheMut.RLock()
	cache, cached := rankVehicleInsertCache[key]
	rankVehicleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rankVehicleAllColumns,
			rankVehicleColumnsWithDefault,
			rankVehicleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rankVehicleType, rankVehicleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rankVehicleType, rankVehicleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"rank_vehicle\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"rank_vehicle\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into rank_vehicle")
	}

	if !cached {
		rankVehicleInsertCacheMut.Lock()
		rankVehicleInsertCache[key] = cache
		rankVehicleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RankVehicle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RankVehicle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	rankVehicleUpdateCacheMut.RLock()
	cache, cached := rankVehicleUpdateCache[key]
	rankVehicleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rankVehicleAllColumns,
			rankVehiclePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update rank_vehicle, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"rank_vehicle\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rankVehiclePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rankVehicleType, rankVehicleMapping, append(wl, rankVehiclePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update rank_vehicle row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for rank_vehicle")
	}

	if !cached {
		rankVehicleUpdateCacheMut.Lock()
		rankVehicleUpdateCache[key] = cache
		rankVehicleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q rankVehicleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for rank_vehicle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for rank_vehicle")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RankVehicleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rankVehiclePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"rank_vehicle\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rankVehiclePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in rankVehicle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all rankVehicle")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RankVehicle) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rank_vehicle provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rankVehicleColumnsWithDefault, o)

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

	rankVehicleUpsertCacheMut.RLock()
	cache, cached := rankVehicleUpsertCache[key]
	rankVehicleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rankVehicleAllColumns,
			rankVehicleColumnsWithDefault,
			rankVehicleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rankVehicleAllColumns,
			rankVehiclePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert rank_vehicle, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rankVehiclePrimaryKeyColumns))
			copy(conflict, rankVehiclePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"rank_vehicle\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rankVehicleType, rankVehicleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rankVehicleType, rankVehicleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert rank_vehicle")
	}

	if !cached {
		rankVehicleUpsertCacheMut.Lock()
		rankVehicleUpsertCache[key] = cache
		rankVehicleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RankVehicle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RankVehicle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RankVehicle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rankVehiclePrimaryKeyMapping)
	sql := "DELETE FROM \"rank_vehicle\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from rank_vehicle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for rank_vehicle")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rankVehicleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rankVehicleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rank_vehicle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rank_vehicle")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RankVehicleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(rankVehicleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rankVehiclePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"rank_vehicle\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rankVehiclePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rankVehicle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rank_vehicle")
	}

	if len(rankVehicleAfterDeleteHooks) != 0 {
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
func (o *RankVehicle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRankVehicle(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RankVehicleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RankVehicleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rankVehiclePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"rank_vehicle\".* FROM \"rank_vehicle\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rankVehiclePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RankVehicleSlice")
	}

	*o = slice

	return nil
}

// RankVehicleExists checks if the RankVehicle row exists.
func RankVehicleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"rank_vehicle\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if rank_vehicle exists")
	}

	return exists, nil
}
