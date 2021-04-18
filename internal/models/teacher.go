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

// Teacher is an object representing the database table.
type Teacher struct {
	ID               int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Fullname         string      `boil:"fullname" json:"fullname" toml:"fullname" yaml:"fullname"`
	Sex              string      `boil:"sex" json:"sex" toml:"sex" yaml:"sex"`
	Dateofbirth      string      `boil:"dateofbirth" json:"dateofbirth" toml:"dateofbirth" yaml:"dateofbirth"`
	Phone            string      `boil:"phone" json:"phone" toml:"phone" yaml:"phone"`
	Address          string      `boil:"address" json:"address" toml:"address" yaml:"address"`
	CMND             string      `boil:"cmnd" json:"cmnd" toml:"cmnd" yaml:"cmnd"`
	CNSK             bool        `boil:"cnsk" json:"cnsk" toml:"cnsk" yaml:"cnsk"`
	GPLX             null.String `boil:"gplx" json:"gplx,omitempty" toml:"gplx" yaml:"gplx,omitempty"`
	ExperienceDriver int         `boil:"experience_driver" json:"experience_driver" toml:"experience_driver" yaml:"experience_driver"`
	KMSafe           int         `boil:"km_safe" json:"km_safe" toml:"km_safe" yaml:"km_safe"`
	CreatedBy        string      `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	CreatedAt        time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedBy        string      `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`
	UpdatedAt        time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	IsDeleted        bool        `boil:"is_deleted" json:"is_deleted" toml:"is_deleted" yaml:"is_deleted"`
	Email            null.String `boil:"email" json:"email,omitempty" toml:"email" yaml:"email,omitempty"`
	IsContract       bool        `boil:"is_contract" json:"is_contract" toml:"is_contract" yaml:"is_contract"`
	IsPractice       bool        `boil:"is_practice" json:"is_practice" toml:"is_practice" yaml:"is_practice"`

	R *teacherR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L teacherL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TeacherColumns = struct {
	ID               string
	Fullname         string
	Sex              string
	Dateofbirth      string
	Phone            string
	Address          string
	CMND             string
	CNSK             string
	GPLX             string
	ExperienceDriver string
	KMSafe           string
	CreatedBy        string
	CreatedAt        string
	UpdatedBy        string
	UpdatedAt        string
	IsDeleted        string
	Email            string
	IsContract       string
	IsPractice       string
}{
	ID:               "id",
	Fullname:         "fullname",
	Sex:              "sex",
	Dateofbirth:      "dateofbirth",
	Phone:            "phone",
	Address:          "address",
	CMND:             "cmnd",
	CNSK:             "cnsk",
	GPLX:             "gplx",
	ExperienceDriver: "experience_driver",
	KMSafe:           "km_safe",
	CreatedBy:        "created_by",
	CreatedAt:        "created_at",
	UpdatedBy:        "updated_by",
	UpdatedAt:        "updated_at",
	IsDeleted:        "is_deleted",
	Email:            "email",
	IsContract:       "is_contract",
	IsPractice:       "is_practice",
}

// Generated where

var TeacherWhere = struct {
	ID               whereHelperint
	Fullname         whereHelperstring
	Sex              whereHelperstring
	Dateofbirth      whereHelperstring
	Phone            whereHelperstring
	Address          whereHelperstring
	CMND             whereHelperstring
	CNSK             whereHelperbool
	GPLX             whereHelpernull_String
	ExperienceDriver whereHelperint
	KMSafe           whereHelperint
	CreatedBy        whereHelperstring
	CreatedAt        whereHelpertime_Time
	UpdatedBy        whereHelperstring
	UpdatedAt        whereHelpertime_Time
	IsDeleted        whereHelperbool
	Email            whereHelpernull_String
	IsContract       whereHelperbool
	IsPractice       whereHelperbool
}{
	ID:               whereHelperint{field: "\"teacher\".\"id\""},
	Fullname:         whereHelperstring{field: "\"teacher\".\"fullname\""},
	Sex:              whereHelperstring{field: "\"teacher\".\"sex\""},
	Dateofbirth:      whereHelperstring{field: "\"teacher\".\"dateofbirth\""},
	Phone:            whereHelperstring{field: "\"teacher\".\"phone\""},
	Address:          whereHelperstring{field: "\"teacher\".\"address\""},
	CMND:             whereHelperstring{field: "\"teacher\".\"cmnd\""},
	CNSK:             whereHelperbool{field: "\"teacher\".\"cnsk\""},
	GPLX:             whereHelpernull_String{field: "\"teacher\".\"gplx\""},
	ExperienceDriver: whereHelperint{field: "\"teacher\".\"experience_driver\""},
	KMSafe:           whereHelperint{field: "\"teacher\".\"km_safe\""},
	CreatedBy:        whereHelperstring{field: "\"teacher\".\"created_by\""},
	CreatedAt:        whereHelpertime_Time{field: "\"teacher\".\"created_at\""},
	UpdatedBy:        whereHelperstring{field: "\"teacher\".\"updated_by\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"teacher\".\"updated_at\""},
	IsDeleted:        whereHelperbool{field: "\"teacher\".\"is_deleted\""},
	Email:            whereHelpernull_String{field: "\"teacher\".\"email\""},
	IsContract:       whereHelperbool{field: "\"teacher\".\"is_contract\""},
	IsPractice:       whereHelperbool{field: "\"teacher\".\"is_practice\""},
}

// TeacherRels is where relationship names are stored.
var TeacherRels = struct {
	Subjects string
}{
	Subjects: "Subjects",
}

// teacherR is where relationships are stored.
type teacherR struct {
	Subjects SubjectSlice
}

// NewStruct creates a new relationship struct
func (*teacherR) NewStruct() *teacherR {
	return &teacherR{}
}

// teacherL is where Load methods for each relationship are stored.
type teacherL struct{}

var (
	teacherAllColumns            = []string{"id", "fullname", "sex", "dateofbirth", "phone", "address", "cmnd", "cnsk", "gplx", "experience_driver", "km_safe", "created_by", "created_at", "updated_by", "updated_at", "is_deleted", "email", "is_contract", "is_practice"}
	teacherColumnsWithoutDefault = []string{"fullname", "sex", "dateofbirth", "phone", "address", "cmnd", "gplx", "experience_driver", "km_safe", "created_by", "updated_by", "email"}
	teacherColumnsWithDefault    = []string{"id", "cnsk", "created_at", "updated_at", "is_deleted", "is_contract", "is_practice"}
	teacherPrimaryKeyColumns     = []string{"id"}
)

type (
	// TeacherSlice is an alias for a slice of pointers to Teacher.
	// This should generally be used opposed to []Teacher.
	TeacherSlice []*Teacher
	// TeacherHook is the signature for custom Teacher hook methods
	TeacherHook func(context.Context, boil.ContextExecutor, *Teacher) error

	teacherQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	teacherType                 = reflect.TypeOf(&Teacher{})
	teacherMapping              = queries.MakeStructMapping(teacherType)
	teacherPrimaryKeyMapping, _ = queries.BindMapping(teacherType, teacherMapping, teacherPrimaryKeyColumns)
	teacherInsertCacheMut       sync.RWMutex
	teacherInsertCache          = make(map[string]insertCache)
	teacherUpdateCacheMut       sync.RWMutex
	teacherUpdateCache          = make(map[string]updateCache)
	teacherUpsertCacheMut       sync.RWMutex
	teacherUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var teacherBeforeInsertHooks []TeacherHook
var teacherBeforeUpdateHooks []TeacherHook
var teacherBeforeDeleteHooks []TeacherHook
var teacherBeforeUpsertHooks []TeacherHook

var teacherAfterInsertHooks []TeacherHook
var teacherAfterSelectHooks []TeacherHook
var teacherAfterUpdateHooks []TeacherHook
var teacherAfterDeleteHooks []TeacherHook
var teacherAfterUpsertHooks []TeacherHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Teacher) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Teacher) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Teacher) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Teacher) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Teacher) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Teacher) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Teacher) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Teacher) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Teacher) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teacherAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTeacherHook registers your hook function for all future operations.
func AddTeacherHook(hookPoint boil.HookPoint, teacherHook TeacherHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		teacherBeforeInsertHooks = append(teacherBeforeInsertHooks, teacherHook)
	case boil.BeforeUpdateHook:
		teacherBeforeUpdateHooks = append(teacherBeforeUpdateHooks, teacherHook)
	case boil.BeforeDeleteHook:
		teacherBeforeDeleteHooks = append(teacherBeforeDeleteHooks, teacherHook)
	case boil.BeforeUpsertHook:
		teacherBeforeUpsertHooks = append(teacherBeforeUpsertHooks, teacherHook)
	case boil.AfterInsertHook:
		teacherAfterInsertHooks = append(teacherAfterInsertHooks, teacherHook)
	case boil.AfterSelectHook:
		teacherAfterSelectHooks = append(teacherAfterSelectHooks, teacherHook)
	case boil.AfterUpdateHook:
		teacherAfterUpdateHooks = append(teacherAfterUpdateHooks, teacherHook)
	case boil.AfterDeleteHook:
		teacherAfterDeleteHooks = append(teacherAfterDeleteHooks, teacherHook)
	case boil.AfterUpsertHook:
		teacherAfterUpsertHooks = append(teacherAfterUpsertHooks, teacherHook)
	}
}

// One returns a single teacher record from the query.
func (q teacherQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Teacher, error) {
	o := &Teacher{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for teacher")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Teacher records from the query.
func (q teacherQuery) All(ctx context.Context, exec boil.ContextExecutor) (TeacherSlice, error) {
	var o []*Teacher

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Teacher slice")
	}

	if len(teacherAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Teacher records in the query.
func (q teacherQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count teacher rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q teacherQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if teacher exists")
	}

	return count > 0, nil
}

// Subjects retrieves all the subject's Subjects with an executor.
func (o *Teacher) Subjects(mods ...qm.QueryMod) subjectQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"subject\".\"teacher_id\"=?", o.ID),
	)

	query := Subjects(queryMods...)
	queries.SetFrom(query.Query, "\"subject\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"subject\".*"})
	}

	return query
}

// LoadSubjects allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (teacherL) LoadSubjects(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTeacher interface{}, mods queries.Applicator) error {
	var slice []*Teacher
	var object *Teacher

	if singular {
		object = maybeTeacher.(*Teacher)
	} else {
		slice = *maybeTeacher.(*[]*Teacher)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &teacherR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &teacherR{}
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

	query := NewQuery(qm.From(`subject`), qm.WhereIn(`subject.teacher_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load subject")
	}

	var resultSlice []*Subject
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice subject")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on subject")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subject")
	}

	if len(subjectAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Subjects = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &subjectR{}
			}
			foreign.R.Teacher = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TeacherID {
				local.R.Subjects = append(local.R.Subjects, foreign)
				if foreign.R == nil {
					foreign.R = &subjectR{}
				}
				foreign.R.Teacher = local
				break
			}
		}
	}

	return nil
}

// AddSubjects adds the given related objects to the existing relationships
// of the teacher, optionally inserting them as new records.
// Appends related to o.R.Subjects.
// Sets related.R.Teacher appropriately.
func (o *Teacher) AddSubjects(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Subject) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TeacherID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"subject\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"teacher_id"}),
				strmangle.WhereClause("\"", "\"", 2, subjectPrimaryKeyColumns),
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

			rel.TeacherID = o.ID
		}
	}

	if o.R == nil {
		o.R = &teacherR{
			Subjects: related,
		}
	} else {
		o.R.Subjects = append(o.R.Subjects, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &subjectR{
				Teacher: o,
			}
		} else {
			rel.R.Teacher = o
		}
	}
	return nil
}

// Teachers retrieves all the records using an executor.
func Teachers(mods ...qm.QueryMod) teacherQuery {
	mods = append(mods, qm.From("\"teacher\""))
	return teacherQuery{NewQuery(mods...)}
}

// FindTeacher retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTeacher(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Teacher, error) {
	teacherObj := &Teacher{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"teacher\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, teacherObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from teacher")
	}

	return teacherObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Teacher) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no teacher provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(teacherColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	teacherInsertCacheMut.RLock()
	cache, cached := teacherInsertCache[key]
	teacherInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			teacherAllColumns,
			teacherColumnsWithDefault,
			teacherColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(teacherType, teacherMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(teacherType, teacherMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"teacher\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"teacher\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into teacher")
	}

	if !cached {
		teacherInsertCacheMut.Lock()
		teacherInsertCache[key] = cache
		teacherInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Teacher.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Teacher) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	teacherUpdateCacheMut.RLock()
	cache, cached := teacherUpdateCache[key]
	teacherUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			teacherAllColumns,
			teacherPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update teacher, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"teacher\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, teacherPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(teacherType, teacherMapping, append(wl, teacherPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update teacher row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for teacher")
	}

	if !cached {
		teacherUpdateCacheMut.Lock()
		teacherUpdateCache[key] = cache
		teacherUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q teacherQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for teacher")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TeacherSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"teacher\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, teacherPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in teacher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all teacher")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Teacher) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no teacher provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(teacherColumnsWithDefault, o)

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

	teacherUpsertCacheMut.RLock()
	cache, cached := teacherUpsertCache[key]
	teacherUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			teacherAllColumns,
			teacherColumnsWithDefault,
			teacherColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			teacherAllColumns,
			teacherPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert teacher, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(teacherPrimaryKeyColumns))
			copy(conflict, teacherPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"teacher\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(teacherType, teacherMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(teacherType, teacherMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert teacher")
	}

	if !cached {
		teacherUpsertCacheMut.Lock()
		teacherUpsertCache[key] = cache
		teacherUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Teacher record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Teacher) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Teacher provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), teacherPrimaryKeyMapping)
	sql := "DELETE FROM \"teacher\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for teacher")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q teacherQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no teacherQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from teacher")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for teacher")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TeacherSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(teacherBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"teacher\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teacherPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from teacher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for teacher")
	}

	if len(teacherAfterDeleteHooks) != 0 {
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
func (o *Teacher) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTeacher(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TeacherSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TeacherSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teacherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"teacher\".* FROM \"teacher\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teacherPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TeacherSlice")
	}

	*o = slice

	return nil
}

// TeacherExists checks if the Teacher row exists.
func TeacherExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"teacher\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if teacher exists")
	}

	return exists, nil
}
