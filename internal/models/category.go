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

// Category is an object representing the database table.
type Category struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title     string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	CreatedBy string    `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedBy string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Meta      string    `boil:"meta" json:"meta" toml:"meta" yaml:"meta"`

	R *categoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L categoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CategoryColumns = struct {
	ID        string
	Title     string
	CreatedBy string
	CreatedAt string
	UpdatedBy string
	UpdatedAt string
	Meta      string
}{
	ID:        "id",
	Title:     "title",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	Meta:      "meta",
}

// Generated where

var CategoryWhere = struct {
	ID        whereHelperint
	Title     whereHelperstring
	CreatedBy whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedBy whereHelperstring
	UpdatedAt whereHelpertime_Time
	Meta      whereHelperstring
}{
	ID:        whereHelperint{field: "\"category\".\"id\""},
	Title:     whereHelperstring{field: "\"category\".\"title\""},
	CreatedBy: whereHelperstring{field: "\"category\".\"created_by\""},
	CreatedAt: whereHelpertime_Time{field: "\"category\".\"created_at\""},
	UpdatedBy: whereHelperstring{field: "\"category\".\"updated_by\""},
	UpdatedAt: whereHelpertime_Time{field: "\"category\".\"updated_at\""},
	Meta:      whereHelperstring{field: "\"category\".\"meta\""},
}

// CategoryRels is where relationship names are stored.
var CategoryRels = struct {
	ChildCategories string
}{
	ChildCategories: "ChildCategories",
}

// categoryR is where relationships are stored.
type categoryR struct {
	ChildCategories ChildCategorySlice
}

// NewStruct creates a new relationship struct
func (*categoryR) NewStruct() *categoryR {
	return &categoryR{}
}

// categoryL is where Load methods for each relationship are stored.
type categoryL struct{}

var (
	categoryAllColumns            = []string{"id", "title", "created_by", "created_at", "updated_by", "updated_at", "meta"}
	categoryColumnsWithoutDefault = []string{"title", "created_by", "updated_by", "meta"}
	categoryColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	categoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// CategorySlice is an alias for a slice of pointers to Category.
	// This should generally be used opposed to []Category.
	CategorySlice []*Category
	// CategoryHook is the signature for custom Category hook methods
	CategoryHook func(context.Context, boil.ContextExecutor, *Category) error

	categoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	categoryType                 = reflect.TypeOf(&Category{})
	categoryMapping              = queries.MakeStructMapping(categoryType)
	categoryPrimaryKeyMapping, _ = queries.BindMapping(categoryType, categoryMapping, categoryPrimaryKeyColumns)
	categoryInsertCacheMut       sync.RWMutex
	categoryInsertCache          = make(map[string]insertCache)
	categoryUpdateCacheMut       sync.RWMutex
	categoryUpdateCache          = make(map[string]updateCache)
	categoryUpsertCacheMut       sync.RWMutex
	categoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var categoryBeforeInsertHooks []CategoryHook
var categoryBeforeUpdateHooks []CategoryHook
var categoryBeforeDeleteHooks []CategoryHook
var categoryBeforeUpsertHooks []CategoryHook

var categoryAfterInsertHooks []CategoryHook
var categoryAfterSelectHooks []CategoryHook
var categoryAfterUpdateHooks []CategoryHook
var categoryAfterDeleteHooks []CategoryHook
var categoryAfterUpsertHooks []CategoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Category) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Category) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Category) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Category) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Category) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Category) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Category) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Category) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Category) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range categoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCategoryHook registers your hook function for all future operations.
func AddCategoryHook(hookPoint boil.HookPoint, categoryHook CategoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		categoryBeforeInsertHooks = append(categoryBeforeInsertHooks, categoryHook)
	case boil.BeforeUpdateHook:
		categoryBeforeUpdateHooks = append(categoryBeforeUpdateHooks, categoryHook)
	case boil.BeforeDeleteHook:
		categoryBeforeDeleteHooks = append(categoryBeforeDeleteHooks, categoryHook)
	case boil.BeforeUpsertHook:
		categoryBeforeUpsertHooks = append(categoryBeforeUpsertHooks, categoryHook)
	case boil.AfterInsertHook:
		categoryAfterInsertHooks = append(categoryAfterInsertHooks, categoryHook)
	case boil.AfterSelectHook:
		categoryAfterSelectHooks = append(categoryAfterSelectHooks, categoryHook)
	case boil.AfterUpdateHook:
		categoryAfterUpdateHooks = append(categoryAfterUpdateHooks, categoryHook)
	case boil.AfterDeleteHook:
		categoryAfterDeleteHooks = append(categoryAfterDeleteHooks, categoryHook)
	case boil.AfterUpsertHook:
		categoryAfterUpsertHooks = append(categoryAfterUpsertHooks, categoryHook)
	}
}

// One returns a single category record from the query.
func (q categoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Category, error) {
	o := &Category{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for category")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Category records from the query.
func (q categoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (CategorySlice, error) {
	var o []*Category

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Category slice")
	}

	if len(categoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Category records in the query.
func (q categoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count category rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q categoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if category exists")
	}

	return count > 0, nil
}

// ChildCategories retrieves all the child_category's ChildCategories with an executor.
func (o *Category) ChildCategories(mods ...qm.QueryMod) childCategoryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"child_category\".\"category_id\"=?", o.ID),
	)

	query := ChildCategories(queryMods...)
	queries.SetFrom(query.Query, "\"child_category\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"child_category\".*"})
	}

	return query
}

// LoadChildCategories allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (categoryL) LoadChildCategories(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCategory interface{}, mods queries.Applicator) error {
	var slice []*Category
	var object *Category

	if singular {
		object = maybeCategory.(*Category)
	} else {
		slice = *maybeCategory.(*[]*Category)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &categoryR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &categoryR{}
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

	query := NewQuery(qm.From(`child_category`), qm.WhereIn(`child_category.category_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load child_category")
	}

	var resultSlice []*ChildCategory
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice child_category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on child_category")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for child_category")
	}

	if len(childCategoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChildCategories = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &childCategoryR{}
			}
			foreign.R.Category = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CategoryID {
				local.R.ChildCategories = append(local.R.ChildCategories, foreign)
				if foreign.R == nil {
					foreign.R = &childCategoryR{}
				}
				foreign.R.Category = local
				break
			}
		}
	}

	return nil
}

// AddChildCategories adds the given related objects to the existing relationships
// of the category, optionally inserting them as new records.
// Appends related to o.R.ChildCategories.
// Sets related.R.Category appropriately.
func (o *Category) AddChildCategories(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChildCategory) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CategoryID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"child_category\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"category_id"}),
				strmangle.WhereClause("\"", "\"", 2, childCategoryPrimaryKeyColumns),
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

			rel.CategoryID = o.ID
		}
	}

	if o.R == nil {
		o.R = &categoryR{
			ChildCategories: related,
		}
	} else {
		o.R.ChildCategories = append(o.R.ChildCategories, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &childCategoryR{
				Category: o,
			}
		} else {
			rel.R.Category = o
		}
	}
	return nil
}

// Categories retrieves all the records using an executor.
func Categories(mods ...qm.QueryMod) categoryQuery {
	mods = append(mods, qm.From("\"category\""))
	return categoryQuery{NewQuery(mods...)}
}

// FindCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCategory(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Category, error) {
	categoryObj := &Category{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"category\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, categoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from category")
	}

	return categoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Category) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no category provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(categoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	categoryInsertCacheMut.RLock()
	cache, cached := categoryInsertCache[key]
	categoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			categoryAllColumns,
			categoryColumnsWithDefault,
			categoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(categoryType, categoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(categoryType, categoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"category\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"category\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into category")
	}

	if !cached {
		categoryInsertCacheMut.Lock()
		categoryInsertCache[key] = cache
		categoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Category.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Category) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	categoryUpdateCacheMut.RLock()
	cache, cached := categoryUpdateCache[key]
	categoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			categoryAllColumns,
			categoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update category, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"category\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, categoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(categoryType, categoryMapping, append(wl, categoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update category row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for category")
	}

	if !cached {
		categoryUpdateCacheMut.Lock()
		categoryUpdateCache[key] = cache
		categoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q categoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for category")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), categoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"category\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, categoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in category slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all category")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Category) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no category provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(categoryColumnsWithDefault, o)

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

	categoryUpsertCacheMut.RLock()
	cache, cached := categoryUpsertCache[key]
	categoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			categoryAllColumns,
			categoryColumnsWithDefault,
			categoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			categoryAllColumns,
			categoryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert category, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(categoryPrimaryKeyColumns))
			copy(conflict, categoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"category\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(categoryType, categoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(categoryType, categoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert category")
	}

	if !cached {
		categoryUpsertCacheMut.Lock()
		categoryUpsertCache[key] = cache
		categoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Category record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Category) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Category provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), categoryPrimaryKeyMapping)
	sql := "DELETE FROM \"category\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for category")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q categoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no categoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for category")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(categoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), categoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, categoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from category slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for category")
	}

	if len(categoryAfterDeleteHooks) != 0 {
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
func (o *Category) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCategory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), categoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"category\".* FROM \"category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, categoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CategorySlice")
	}

	*o = slice

	return nil
}

// CategoryExists checks if the Category row exists.
func CategoryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"category\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if category exists")
	}

	return exists, nil
}
