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

// ChildCategory is an object representing the database table.
type ChildCategory struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title      string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	CategoryID int       `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	Meta       string    `boil:"meta" json:"meta" toml:"meta" yaml:"meta"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy  string    `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy  string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`
	IsDeleted  bool      `boil:"is_deleted" json:"is_deleted" toml:"is_deleted" yaml:"is_deleted"`

	R *childCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L childCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChildCategoryColumns = struct {
	ID         string
	Title      string
	CategoryID string
	Meta       string
	CreatedAt  string
	CreatedBy  string
	UpdatedAt  string
	UpdatedBy  string
	IsDeleted  string
}{
	ID:         "id",
	Title:      "title",
	CategoryID: "category_id",
	Meta:       "meta",
	CreatedAt:  "created_at",
	CreatedBy:  "created_by",
	UpdatedAt:  "updated_at",
	UpdatedBy:  "updated_by",
	IsDeleted:  "is_deleted",
}

// Generated where

var ChildCategoryWhere = struct {
	ID         whereHelperint
	Title      whereHelperstring
	CategoryID whereHelperint
	Meta       whereHelperstring
	CreatedAt  whereHelpertime_Time
	CreatedBy  whereHelperstring
	UpdatedAt  whereHelpertime_Time
	UpdatedBy  whereHelperstring
	IsDeleted  whereHelperbool
}{
	ID:         whereHelperint{field: "\"child_category\".\"id\""},
	Title:      whereHelperstring{field: "\"child_category\".\"title\""},
	CategoryID: whereHelperint{field: "\"child_category\".\"category_id\""},
	Meta:       whereHelperstring{field: "\"child_category\".\"meta\""},
	CreatedAt:  whereHelpertime_Time{field: "\"child_category\".\"created_at\""},
	CreatedBy:  whereHelperstring{field: "\"child_category\".\"created_by\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"child_category\".\"updated_at\""},
	UpdatedBy:  whereHelperstring{field: "\"child_category\".\"updated_by\""},
	IsDeleted:  whereHelperbool{field: "\"child_category\".\"is_deleted\""},
}

// ChildCategoryRels is where relationship names are stored.
var ChildCategoryRels = struct {
	Category string
	Articles string
}{
	Category: "Category",
	Articles: "Articles",
}

// childCategoryR is where relationships are stored.
type childCategoryR struct {
	Category *Category
	Articles ArticleSlice
}

// NewStruct creates a new relationship struct
func (*childCategoryR) NewStruct() *childCategoryR {
	return &childCategoryR{}
}

// childCategoryL is where Load methods for each relationship are stored.
type childCategoryL struct{}

var (
	childCategoryAllColumns            = []string{"id", "title", "category_id", "meta", "created_at", "created_by", "updated_at", "updated_by", "is_deleted"}
	childCategoryColumnsWithoutDefault = []string{"title", "category_id", "meta", "created_by", "updated_by"}
	childCategoryColumnsWithDefault    = []string{"id", "created_at", "updated_at", "is_deleted"}
	childCategoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChildCategorySlice is an alias for a slice of pointers to ChildCategory.
	// This should generally be used opposed to []ChildCategory.
	ChildCategorySlice []*ChildCategory
	// ChildCategoryHook is the signature for custom ChildCategory hook methods
	ChildCategoryHook func(context.Context, boil.ContextExecutor, *ChildCategory) error

	childCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	childCategoryType                 = reflect.TypeOf(&ChildCategory{})
	childCategoryMapping              = queries.MakeStructMapping(childCategoryType)
	childCategoryPrimaryKeyMapping, _ = queries.BindMapping(childCategoryType, childCategoryMapping, childCategoryPrimaryKeyColumns)
	childCategoryInsertCacheMut       sync.RWMutex
	childCategoryInsertCache          = make(map[string]insertCache)
	childCategoryUpdateCacheMut       sync.RWMutex
	childCategoryUpdateCache          = make(map[string]updateCache)
	childCategoryUpsertCacheMut       sync.RWMutex
	childCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var childCategoryBeforeInsertHooks []ChildCategoryHook
var childCategoryBeforeUpdateHooks []ChildCategoryHook
var childCategoryBeforeDeleteHooks []ChildCategoryHook
var childCategoryBeforeUpsertHooks []ChildCategoryHook

var childCategoryAfterInsertHooks []ChildCategoryHook
var childCategoryAfterSelectHooks []ChildCategoryHook
var childCategoryAfterUpdateHooks []ChildCategoryHook
var childCategoryAfterDeleteHooks []ChildCategoryHook
var childCategoryAfterUpsertHooks []ChildCategoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ChildCategory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ChildCategory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ChildCategory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ChildCategory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ChildCategory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ChildCategory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ChildCategory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ChildCategory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ChildCategory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childCategoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChildCategoryHook registers your hook function for all future operations.
func AddChildCategoryHook(hookPoint boil.HookPoint, childCategoryHook ChildCategoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		childCategoryBeforeInsertHooks = append(childCategoryBeforeInsertHooks, childCategoryHook)
	case boil.BeforeUpdateHook:
		childCategoryBeforeUpdateHooks = append(childCategoryBeforeUpdateHooks, childCategoryHook)
	case boil.BeforeDeleteHook:
		childCategoryBeforeDeleteHooks = append(childCategoryBeforeDeleteHooks, childCategoryHook)
	case boil.BeforeUpsertHook:
		childCategoryBeforeUpsertHooks = append(childCategoryBeforeUpsertHooks, childCategoryHook)
	case boil.AfterInsertHook:
		childCategoryAfterInsertHooks = append(childCategoryAfterInsertHooks, childCategoryHook)
	case boil.AfterSelectHook:
		childCategoryAfterSelectHooks = append(childCategoryAfterSelectHooks, childCategoryHook)
	case boil.AfterUpdateHook:
		childCategoryAfterUpdateHooks = append(childCategoryAfterUpdateHooks, childCategoryHook)
	case boil.AfterDeleteHook:
		childCategoryAfterDeleteHooks = append(childCategoryAfterDeleteHooks, childCategoryHook)
	case boil.AfterUpsertHook:
		childCategoryAfterUpsertHooks = append(childCategoryAfterUpsertHooks, childCategoryHook)
	}
}

// One returns a single childCategory record from the query.
func (q childCategoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ChildCategory, error) {
	o := &ChildCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for child_category")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ChildCategory records from the query.
func (q childCategoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChildCategorySlice, error) {
	var o []*ChildCategory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ChildCategory slice")
	}

	if len(childCategoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ChildCategory records in the query.
func (q childCategoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count child_category rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q childCategoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if child_category exists")
	}

	return count > 0, nil
}

// Category pointed to by the foreign key.
func (o *ChildCategory) Category(mods ...qm.QueryMod) categoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CategoryID),
	}

	queryMods = append(queryMods, mods...)

	query := Categories(queryMods...)
	queries.SetFrom(query.Query, "\"category\"")

	return query
}

// Articles retrieves all the article's Articles with an executor.
func (o *ChildCategory) Articles(mods ...qm.QueryMod) articleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"articles\".\"child_category_id\"=?", o.ID),
	)

	query := Articles(queryMods...)
	queries.SetFrom(query.Query, "\"articles\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"articles\".*"})
	}

	return query
}

// LoadCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (childCategoryL) LoadCategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChildCategory interface{}, mods queries.Applicator) error {
	var slice []*ChildCategory
	var object *ChildCategory

	if singular {
		object = maybeChildCategory.(*ChildCategory)
	} else {
		slice = *maybeChildCategory.(*[]*ChildCategory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &childCategoryR{}
		}
		args = append(args, object.CategoryID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &childCategoryR{}
			}

			for _, a := range args {
				if a == obj.CategoryID {
					continue Outer
				}
			}

			args = append(args, obj.CategoryID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`category`), qm.WhereIn(`category.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Category")
	}

	var resultSlice []*Category
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for category")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for category")
	}

	if len(childCategoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Category = foreign
		if foreign.R == nil {
			foreign.R = &categoryR{}
		}
		foreign.R.ChildCategories = append(foreign.R.ChildCategories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CategoryID == foreign.ID {
				local.R.Category = foreign
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.ChildCategories = append(foreign.R.ChildCategories, local)
				break
			}
		}
	}

	return nil
}

// LoadArticles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (childCategoryL) LoadArticles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChildCategory interface{}, mods queries.Applicator) error {
	var slice []*ChildCategory
	var object *ChildCategory

	if singular {
		object = maybeChildCategory.(*ChildCategory)
	} else {
		slice = *maybeChildCategory.(*[]*ChildCategory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &childCategoryR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &childCategoryR{}
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

	query := NewQuery(qm.From(`articles`), qm.WhereIn(`articles.child_category_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load articles")
	}

	var resultSlice []*Article
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice articles")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on articles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for articles")
	}

	if len(articleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Articles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &articleR{}
			}
			foreign.R.ChildCategory = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChildCategoryID {
				local.R.Articles = append(local.R.Articles, foreign)
				if foreign.R == nil {
					foreign.R = &articleR{}
				}
				foreign.R.ChildCategory = local
				break
			}
		}
	}

	return nil
}

// SetCategory of the childCategory to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.ChildCategories.
func (o *ChildCategory) SetCategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Category) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"child_category\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"category_id"}),
		strmangle.WhereClause("\"", "\"", 2, childCategoryPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CategoryID = related.ID
	if o.R == nil {
		o.R = &childCategoryR{
			Category: related,
		}
	} else {
		o.R.Category = related
	}

	if related.R == nil {
		related.R = &categoryR{
			ChildCategories: ChildCategorySlice{o},
		}
	} else {
		related.R.ChildCategories = append(related.R.ChildCategories, o)
	}

	return nil
}

// AddArticles adds the given related objects to the existing relationships
// of the child_category, optionally inserting them as new records.
// Appends related to o.R.Articles.
// Sets related.R.ChildCategory appropriately.
func (o *ChildCategory) AddArticles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Article) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChildCategoryID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"articles\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"child_category_id"}),
				strmangle.WhereClause("\"", "\"", 2, articlePrimaryKeyColumns),
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

			rel.ChildCategoryID = o.ID
		}
	}

	if o.R == nil {
		o.R = &childCategoryR{
			Articles: related,
		}
	} else {
		o.R.Articles = append(o.R.Articles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &articleR{
				ChildCategory: o,
			}
		} else {
			rel.R.ChildCategory = o
		}
	}
	return nil
}

// ChildCategories retrieves all the records using an executor.
func ChildCategories(mods ...qm.QueryMod) childCategoryQuery {
	mods = append(mods, qm.From("\"child_category\""))
	return childCategoryQuery{NewQuery(mods...)}
}

// FindChildCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChildCategory(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ChildCategory, error) {
	childCategoryObj := &ChildCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"child_category\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, childCategoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from child_category")
	}

	return childCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ChildCategory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no child_category provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(childCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	childCategoryInsertCacheMut.RLock()
	cache, cached := childCategoryInsertCache[key]
	childCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			childCategoryAllColumns,
			childCategoryColumnsWithDefault,
			childCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(childCategoryType, childCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(childCategoryType, childCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"child_category\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"child_category\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into child_category")
	}

	if !cached {
		childCategoryInsertCacheMut.Lock()
		childCategoryInsertCache[key] = cache
		childCategoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ChildCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ChildCategory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	childCategoryUpdateCacheMut.RLock()
	cache, cached := childCategoryUpdateCache[key]
	childCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			childCategoryAllColumns,
			childCategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update child_category, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"child_category\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, childCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(childCategoryType, childCategoryMapping, append(wl, childCategoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update child_category row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for child_category")
	}

	if !cached {
		childCategoryUpdateCacheMut.Lock()
		childCategoryUpdateCache[key] = cache
		childCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q childCategoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for child_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for child_category")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChildCategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"child_category\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, childCategoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in childCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all childCategory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ChildCategory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no child_category provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(childCategoryColumnsWithDefault, o)

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

	childCategoryUpsertCacheMut.RLock()
	cache, cached := childCategoryUpsertCache[key]
	childCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			childCategoryAllColumns,
			childCategoryColumnsWithDefault,
			childCategoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			childCategoryAllColumns,
			childCategoryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert child_category, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(childCategoryPrimaryKeyColumns))
			copy(conflict, childCategoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"child_category\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(childCategoryType, childCategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(childCategoryType, childCategoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert child_category")
	}

	if !cached {
		childCategoryUpsertCacheMut.Lock()
		childCategoryUpsertCache[key] = cache
		childCategoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ChildCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ChildCategory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ChildCategory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), childCategoryPrimaryKeyMapping)
	sql := "DELETE FROM \"child_category\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from child_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for child_category")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q childCategoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no childCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from child_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for child_category")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChildCategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(childCategoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"child_category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, childCategoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from childCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for child_category")
	}

	if len(childCategoryAfterDeleteHooks) != 0 {
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
func (o *ChildCategory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChildCategory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChildCategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChildCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"child_category\".* FROM \"child_category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, childCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChildCategorySlice")
	}

	*o = slice

	return nil
}

// ChildCategoryExists checks if the ChildCategory row exists.
func ChildCategoryExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"child_category\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if child_category exists")
	}

	return exists, nil
}
