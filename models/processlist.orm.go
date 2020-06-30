
package models 

// !!! DO NOT EDIT THIS FILE

import (
	
	"context"
	"gopkg.in/guregu/null.v3"
	"github.com/mylxsw/eloquent/query"
	"github.com/iancoleman/strcase"
)

func init() {

}



// Processlist is a Processlist object
type Processlist struct {
	original *processlistOriginal
	processlistModel *ProcesslistModel

	
	Id int64 
	User string 
	Host string 
	Db string 
	Command string 
	Time int64 
	State string 
	Info string 
}

// SetModel set model for Processlist
func (inst *Processlist) SetModel(processlistModel *ProcesslistModel) {
	inst.processlistModel = processlistModel
}

// processlistOriginal is an object which stores original Processlist from database
type processlistOriginal struct {
	
	Id int64
	User string
	Host string
	Db string
	Command string
	Time int64
	State string
	Info string
}

// Staled identify whether the object has been modified
func (inst *Processlist) Staled() bool {
	if inst.original == nil {
		inst.original = &processlistOriginal {}
	}

	
	if inst.Id != inst.original.Id {
		return true
	}
	if inst.User != inst.original.User {
		return true
	}
	if inst.Host != inst.original.Host {
		return true
	}
	if inst.Db != inst.original.Db {
		return true
	}
	if inst.Command != inst.original.Command {
		return true
	}
	if inst.Time != inst.original.Time {
		return true
	}
	if inst.State != inst.original.State {
		return true
	}
	if inst.Info != inst.original.Info {
		return true
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *Processlist) StaledKV() query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &processlistOriginal {}
	}

	
	if inst.Id != inst.original.Id {
		kv["id"] = inst.Id
	}
	if inst.User != inst.original.User {
		kv["user"] = inst.User
	}
	if inst.Host != inst.original.Host {
		kv["host"] = inst.Host
	}
	if inst.Db != inst.original.Db {
		kv["db"] = inst.Db
	}
	if inst.Command != inst.original.Command {
		kv["command"] = inst.Command
	}
	if inst.Time != inst.original.Time {
		kv["time"] = inst.Time
	}
	if inst.State != inst.original.State {
		kv["state"] = inst.State
	}
	if inst.Info != inst.original.Info {
		kv["info"] = inst.Info
	}

	return kv
}

// Save create a new model or update it 
func (inst *Processlist) Save() error {
	if inst.processlistModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.processlistModel.SaveOrUpdate(*inst)
	if err != nil {
		return err 
	}

	inst.Id = id
	return nil
}

// Delete remove a processlist
func (inst *Processlist) Delete() error {
	if inst.processlistModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.processlistModel.DeleteById(inst.Id)
	if err != nil {
		return err 
	}

	return nil
}





type processlistScope struct {
	name  string
	apply func(builder query.Condition)
}

var processlistGlobalScopes = make([]processlistScope, 0)
var processlistLocalScopes = make([]processlistScope, 0)

// AddGlobalScopeForProcesslist assign a global scope to a model
func AddGlobalScopeForProcesslist(name string, apply func(builder query.Condition)) {
	processlistGlobalScopes = append(processlistGlobalScopes, processlistScope{name: name, apply: apply})
}

// AddLocalScopeForProcesslist assign a local scope to a model
func AddLocalScopeForProcesslist(name string, apply func(builder query.Condition)) {
	processlistLocalScopes = append(processlistLocalScopes, processlistScope{name: name, apply: apply})
}

func (m *ProcesslistModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range processlistGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range processlistLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *ProcesslistModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *ProcesslistModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}
	
	return true
}


type processlistWrap struct { 	
	Id null.Int	
	User null.String	
	Host null.String	
	Db null.String	
	Command null.String	
	Time null.Int	
	State null.String	
	Info null.String
}

func (w processlistWrap) ToProcesslist () Processlist {
	return Processlist {
		original: &processlistOriginal { 
			Id: w.Id.Int64,
			User: w.User.String,
			Host: w.Host.String,
			Db: w.Db.String,
			Command: w.Command.String,
			Time: w.Time.Int64,
			State: w.State.String,
			Info: w.Info.String,
		},
		
		Id: w.Id.Int64,
		User: w.User.String,
		Host: w.Host.String,
		Db: w.Db.String,
		Command: w.Command.String,
		Time: w.Time.Int64,
		State: w.State.String,
		Info: w.Info.String,
	}
}


// ProcesslistModel is a model which encapsulates the operations of the object
type ProcesslistModel struct {
	db *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes []string
	
	query query.SQLBuilder
}

var processlistTableName = "processlist"

func SetProcesslistTable (tableName string) {
	processlistTableName = tableName
}

// NewProcesslistModel create a ProcesslistModel
func NewProcesslistModel (db query.Database) *ProcesslistModel {
	return &ProcesslistModel {
		db: query.NewDatabaseWrap(db), 
		tableName: processlistTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes: make([]string, 0),
		query: query.Builder(),
	}
}

// GetDB return database instance
func (m *ProcesslistModel) GetDB() query.Database {
	return m.db.GetDB()
}



func (m *ProcesslistModel) clone() *ProcesslistModel {
	return &ProcesslistModel{
		db: m.db, 
		tableName: m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes: append([]string{}, m.includeLocalScopes...),
		query: m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *ProcesslistModel) WithoutGlobalScopes(names ...string) *ProcesslistModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *ProcesslistModel) WithLocalScopes(names ...string) *ProcesslistModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Query add query builder to model
func (m *ProcesslistModel) Query(builder query.SQLBuilder) *ProcesslistModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *ProcesslistModel) Find(id int64) (Processlist, error) {
	return m.First(m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *ProcesslistModel) Exists(builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *ProcesslistModel) Count(builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()
	
	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *ProcesslistModel) Paginate(page int64, perPage int64, builders ...query.SQLBuilder) ([]Processlist, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta {
		PerPage: perPage,
		Page: page,
	}

	count, err := m.Count(builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count % perPage != 0 {
		meta.LastPage += 1
	}


	res, err := m.Get(append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *ProcesslistModel) Get(builders ...query.SQLBuilder) ([]Processlist, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"user",
			"host",
			"db",
			"command",
			"time",
			"state",
			"info",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {
		 
		case "id":
			selectFields = append(selectFields, f) 
		case "user":
			selectFields = append(selectFields, f) 
		case "host":
			selectFields = append(selectFields, f) 
		case "db":
			selectFields = append(selectFields, f) 
		case "command":
			selectFields = append(selectFields, f) 
		case "time":
			selectFields = append(selectFields, f) 
		case "state":
			selectFields = append(selectFields, f) 
		case "info":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*processlistWrap, []interface{}) {
		var processlistVar processlistWrap
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {
			 
			case "id":
				scanFields = append(scanFields, &processlistVar.Id) 
			case "user":
				scanFields = append(scanFields, &processlistVar.User) 
			case "host":
				scanFields = append(scanFields, &processlistVar.Host) 
			case "db":
				scanFields = append(scanFields, &processlistVar.Db) 
			case "command":
				scanFields = append(scanFields, &processlistVar.Command) 
			case "time":
				scanFields = append(scanFields, &processlistVar.Time) 
			case "state":
				scanFields = append(scanFields, &processlistVar.State) 
			case "info":
				scanFields = append(scanFields, &processlistVar.Info)
			}
		}

		return &processlistVar, scanFields
	}
	
	sqlStr, params := b.Fields(selectFields...).ResolveQuery()
	
	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	processlists := make([]Processlist, 0)
	for rows.Next() {
		processlistVar, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		processlistReal := processlistVar.ToProcesslist()
		processlistReal.SetModel(m)
		processlists = append(processlists, processlistReal)
	}

	return processlists, nil
}

// First return first result for given query
func (m *ProcesslistModel) First(builders ...query.SQLBuilder) (Processlist, error) {
	res, err := m.Get(append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return Processlist{}, err 
	}

	if len(res) == 0 {
		return Processlist{}, query.ErrNoResult
	}

	return res[0], nil
}

// Create save a new processlist to database
func (m *ProcesslistModel) Create(kv query.KV) (int64, error) {
	
	

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all processlists to database
func (m *ProcesslistModel) SaveAll(processlists []Processlist) ([]int64, error) {
	ids := make([]int64, 0)
	for _, processlist := range processlists {
		id, err := m.Save(processlist)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a processlist to database
func (m *ProcesslistModel) Save(processlist Processlist) (int64, error) {
	return m.Create(processlist.StaledKV())
}

// SaveOrUpdate save a new processlist or update it when it has a id > 0
func (m *ProcesslistModel) SaveOrUpdate(processlist Processlist) (id int64, updated bool, err error) {
	if processlist.Id > 0 {
		_, _err := m.UpdateById(processlist.Id, processlist)
		return processlist.Id, true, _err
	}

	_id, _err := m.Save(processlist)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *ProcesslistModel) UpdateFields(kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *ProcesslistModel) Update(processlist Processlist) (int64, error) {
	return m.UpdateFields(processlist.StaledKV())
}

// UpdateById update a model by id
func (m *ProcesslistModel) UpdateById(id int64, processlist Processlist) (int64, error) {
	return m.Query(query.Builder().Where("id", "=", id)).Update(processlist)
}



// Delete remove a model
func (m *ProcesslistModel) Delete(builders ...query.SQLBuilder) (int64, error) {
	
	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).Table(m.tableName).ResolveDelete()

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
	
}

// DeleteById remove a model by id
func (m *ProcesslistModel) DeleteById(id int64) (int64, error) {
	return m.Query(query.Builder().Where("id", "=", id)).Delete()
}


