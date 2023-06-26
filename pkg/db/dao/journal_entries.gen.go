// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/israel-duff/ledger-system/pkg/db/model"
)

func newJournalEntry(db *gorm.DB, opts ...gen.DOOption) journalEntry {
	_journalEntry := journalEntry{}

	_journalEntry.journalEntryDo.UseDB(db, opts...)
	_journalEntry.journalEntryDo.UseModel(&model.JournalEntry{})

	tableName := _journalEntry.journalEntryDo.TableName()
	_journalEntry.ALL = field.NewAsterisk(tableName)
	_journalEntry.ID = field.NewString(tableName, "id")
	_journalEntry.CreatedAt = field.NewTime(tableName, "created_at")
	_journalEntry.UpdatedAt = field.NewTime(tableName, "updated_at")
	_journalEntry.DeletedAt = field.NewField(tableName, "deleted_at")
	_journalEntry.Name = field.NewString(tableName, "name")
	_journalEntry.Type = field.NewString(tableName, "type")
	_journalEntry.Amount = field.NewFloat64(tableName, "amount")
	_journalEntry.BlockID = field.NewString(tableName, "block_id")
	_journalEntry.TransactionID = field.NewString(tableName, "transaction_id")

	_journalEntry.fillFieldMap()

	return _journalEntry
}

type journalEntry struct {
	journalEntryDo journalEntryDo

	ALL           field.Asterisk
	ID            field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	Name          field.String
	Type          field.String
	Amount        field.Float64
	BlockID       field.String
	TransactionID field.String

	fieldMap map[string]field.Expr
}

func (j journalEntry) Table(newTableName string) *journalEntry {
	j.journalEntryDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j journalEntry) As(alias string) *journalEntry {
	j.journalEntryDo.DO = *(j.journalEntryDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *journalEntry) updateTableName(table string) *journalEntry {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewString(table, "id")
	j.CreatedAt = field.NewTime(table, "created_at")
	j.UpdatedAt = field.NewTime(table, "updated_at")
	j.DeletedAt = field.NewField(table, "deleted_at")
	j.Name = field.NewString(table, "name")
	j.Type = field.NewString(table, "type")
	j.Amount = field.NewFloat64(table, "amount")
	j.BlockID = field.NewString(table, "block_id")
	j.TransactionID = field.NewString(table, "transaction_id")

	j.fillFieldMap()

	return j
}

func (j *journalEntry) WithContext(ctx context.Context) *journalEntryDo {
	return j.journalEntryDo.WithContext(ctx)
}

func (j journalEntry) TableName() string { return j.journalEntryDo.TableName() }

func (j journalEntry) Alias() string { return j.journalEntryDo.Alias() }

func (j *journalEntry) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *journalEntry) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 9)
	j.fieldMap["id"] = j.ID
	j.fieldMap["created_at"] = j.CreatedAt
	j.fieldMap["updated_at"] = j.UpdatedAt
	j.fieldMap["deleted_at"] = j.DeletedAt
	j.fieldMap["name"] = j.Name
	j.fieldMap["type"] = j.Type
	j.fieldMap["amount"] = j.Amount
	j.fieldMap["block_id"] = j.BlockID
	j.fieldMap["transaction_id"] = j.TransactionID
}

func (j journalEntry) clone(db *gorm.DB) journalEntry {
	j.journalEntryDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j journalEntry) replaceDB(db *gorm.DB) journalEntry {
	j.journalEntryDo.ReplaceDB(db)
	return j
}

type journalEntryDo struct{ gen.DO }

func (j journalEntryDo) Debug() *journalEntryDo {
	return j.withDO(j.DO.Debug())
}

func (j journalEntryDo) WithContext(ctx context.Context) *journalEntryDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j journalEntryDo) ReadDB() *journalEntryDo {
	return j.Clauses(dbresolver.Read)
}

func (j journalEntryDo) WriteDB() *journalEntryDo {
	return j.Clauses(dbresolver.Write)
}

func (j journalEntryDo) Session(config *gorm.Session) *journalEntryDo {
	return j.withDO(j.DO.Session(config))
}

func (j journalEntryDo) Clauses(conds ...clause.Expression) *journalEntryDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j journalEntryDo) Returning(value interface{}, columns ...string) *journalEntryDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j journalEntryDo) Not(conds ...gen.Condition) *journalEntryDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j journalEntryDo) Or(conds ...gen.Condition) *journalEntryDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j journalEntryDo) Select(conds ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j journalEntryDo) Where(conds ...gen.Condition) *journalEntryDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j journalEntryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *journalEntryDo {
	return j.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (j journalEntryDo) Order(conds ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j journalEntryDo) Distinct(cols ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j journalEntryDo) Omit(cols ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j journalEntryDo) Join(table schema.Tabler, on ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j journalEntryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j journalEntryDo) RightJoin(table schema.Tabler, on ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j journalEntryDo) Group(cols ...field.Expr) *journalEntryDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j journalEntryDo) Having(conds ...gen.Condition) *journalEntryDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j journalEntryDo) Limit(limit int) *journalEntryDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j journalEntryDo) Offset(offset int) *journalEntryDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j journalEntryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *journalEntryDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j journalEntryDo) Unscoped() *journalEntryDo {
	return j.withDO(j.DO.Unscoped())
}

func (j journalEntryDo) Create(values ...*model.JournalEntry) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j journalEntryDo) CreateInBatches(values []*model.JournalEntry, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j journalEntryDo) Save(values ...*model.JournalEntry) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j journalEntryDo) First() (*model.JournalEntry, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JournalEntry), nil
	}
}

func (j journalEntryDo) Take() (*model.JournalEntry, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JournalEntry), nil
	}
}

func (j journalEntryDo) Last() (*model.JournalEntry, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JournalEntry), nil
	}
}

func (j journalEntryDo) Find() ([]*model.JournalEntry, error) {
	result, err := j.DO.Find()
	return result.([]*model.JournalEntry), err
}

func (j journalEntryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JournalEntry, err error) {
	buf := make([]*model.JournalEntry, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j journalEntryDo) FindInBatches(result *[]*model.JournalEntry, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j journalEntryDo) Attrs(attrs ...field.AssignExpr) *journalEntryDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j journalEntryDo) Assign(attrs ...field.AssignExpr) *journalEntryDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j journalEntryDo) Joins(fields ...field.RelationField) *journalEntryDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j journalEntryDo) Preload(fields ...field.RelationField) *journalEntryDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j journalEntryDo) FirstOrInit() (*model.JournalEntry, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JournalEntry), nil
	}
}

func (j journalEntryDo) FirstOrCreate() (*model.JournalEntry, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JournalEntry), nil
	}
}

func (j journalEntryDo) FindByPage(offset int, limit int) (result []*model.JournalEntry, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j journalEntryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j journalEntryDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j journalEntryDo) Delete(models ...*model.JournalEntry) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *journalEntryDo) withDO(do gen.Dao) *journalEntryDo {
	j.DO = *do.(*gen.DO)
	return j
}