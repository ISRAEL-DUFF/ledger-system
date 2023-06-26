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

func newTransaction(db *gorm.DB, opts ...gen.DOOption) transaction {
	_transaction := transaction{}

	_transaction.transactionDo.UseDB(db, opts...)
	_transaction.transactionDo.UseModel(&model.Transaction{})

	tableName := _transaction.transactionDo.TableName()
	_transaction.ALL = field.NewAsterisk(tableName)
	_transaction.ID = field.NewString(tableName, "id")
	_transaction.CreatedAt = field.NewTime(tableName, "created_at")
	_transaction.UpdatedAt = field.NewTime(tableName, "updated_at")
	_transaction.DeletedAt = field.NewField(tableName, "deleted_at")
	_transaction.Description = field.NewString(tableName, "description")
	_transaction.Status = field.NewString(tableName, "status")

	_transaction.fillFieldMap()

	return _transaction
}

type transaction struct {
	transactionDo transactionDo

	ALL         field.Asterisk
	ID          field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	Description field.String
	Status      field.String

	fieldMap map[string]field.Expr
}

func (t transaction) Table(newTableName string) *transaction {
	t.transactionDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t transaction) As(alias string) *transaction {
	t.transactionDo.DO = *(t.transactionDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *transaction) updateTableName(table string) *transaction {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewString(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.Description = field.NewString(table, "description")
	t.Status = field.NewString(table, "status")

	t.fillFieldMap()

	return t
}

func (t *transaction) WithContext(ctx context.Context) *transactionDo {
	return t.transactionDo.WithContext(ctx)
}

func (t transaction) TableName() string { return t.transactionDo.TableName() }

func (t transaction) Alias() string { return t.transactionDo.Alias() }

func (t *transaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *transaction) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["description"] = t.Description
	t.fieldMap["status"] = t.Status
}

func (t transaction) clone(db *gorm.DB) transaction {
	t.transactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t transaction) replaceDB(db *gorm.DB) transaction {
	t.transactionDo.ReplaceDB(db)
	return t
}

type transactionDo struct{ gen.DO }

func (t transactionDo) Debug() *transactionDo {
	return t.withDO(t.DO.Debug())
}

func (t transactionDo) WithContext(ctx context.Context) *transactionDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t transactionDo) ReadDB() *transactionDo {
	return t.Clauses(dbresolver.Read)
}

func (t transactionDo) WriteDB() *transactionDo {
	return t.Clauses(dbresolver.Write)
}

func (t transactionDo) Session(config *gorm.Session) *transactionDo {
	return t.withDO(t.DO.Session(config))
}

func (t transactionDo) Clauses(conds ...clause.Expression) *transactionDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t transactionDo) Returning(value interface{}, columns ...string) *transactionDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t transactionDo) Not(conds ...gen.Condition) *transactionDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t transactionDo) Or(conds ...gen.Condition) *transactionDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t transactionDo) Select(conds ...field.Expr) *transactionDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t transactionDo) Where(conds ...gen.Condition) *transactionDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t transactionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *transactionDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t transactionDo) Order(conds ...field.Expr) *transactionDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t transactionDo) Distinct(cols ...field.Expr) *transactionDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t transactionDo) Omit(cols ...field.Expr) *transactionDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t transactionDo) Join(table schema.Tabler, on ...field.Expr) *transactionDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t transactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *transactionDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t transactionDo) RightJoin(table schema.Tabler, on ...field.Expr) *transactionDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t transactionDo) Group(cols ...field.Expr) *transactionDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t transactionDo) Having(conds ...gen.Condition) *transactionDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t transactionDo) Limit(limit int) *transactionDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t transactionDo) Offset(offset int) *transactionDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t transactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *transactionDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t transactionDo) Unscoped() *transactionDo {
	return t.withDO(t.DO.Unscoped())
}

func (t transactionDo) Create(values ...*model.Transaction) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t transactionDo) CreateInBatches(values []*model.Transaction, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t transactionDo) Save(values ...*model.Transaction) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t transactionDo) First() (*model.Transaction, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transaction), nil
	}
}

func (t transactionDo) Take() (*model.Transaction, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transaction), nil
	}
}

func (t transactionDo) Last() (*model.Transaction, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transaction), nil
	}
}

func (t transactionDo) Find() ([]*model.Transaction, error) {
	result, err := t.DO.Find()
	return result.([]*model.Transaction), err
}

func (t transactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Transaction, err error) {
	buf := make([]*model.Transaction, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t transactionDo) FindInBatches(result *[]*model.Transaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t transactionDo) Attrs(attrs ...field.AssignExpr) *transactionDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t transactionDo) Assign(attrs ...field.AssignExpr) *transactionDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t transactionDo) Joins(fields ...field.RelationField) *transactionDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t transactionDo) Preload(fields ...field.RelationField) *transactionDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t transactionDo) FirstOrInit() (*model.Transaction, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transaction), nil
	}
}

func (t transactionDo) FirstOrCreate() (*model.Transaction, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Transaction), nil
	}
}

func (t transactionDo) FindByPage(offset int, limit int) (result []*model.Transaction, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t transactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t transactionDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t transactionDo) Delete(models ...*model.Transaction) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *transactionDo) withDO(do gen.Dao) *transactionDo {
	t.DO = *do.(*gen.DO)
	return t
}
