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

func newAPIKey(db *gorm.DB, opts ...gen.DOOption) aPIKey {
	_aPIKey := aPIKey{}

	_aPIKey.aPIKeyDo.UseDB(db, opts...)
	_aPIKey.aPIKeyDo.UseModel(&model.APIKey{})

	tableName := _aPIKey.aPIKeyDo.TableName()
	_aPIKey.ALL = field.NewAsterisk(tableName)
	_aPIKey.ID = field.NewString(tableName, "id")
	_aPIKey.CreatedAt = field.NewTime(tableName, "created_at")
	_aPIKey.UpdatedAt = field.NewTime(tableName, "updated_at")
	_aPIKey.DeletedAt = field.NewField(tableName, "deleted_at")
	_aPIKey.TestSecretKey = field.NewString(tableName, "test_secret_key")
	_aPIKey.TestPublicKey = field.NewString(tableName, "test_public_key")
	_aPIKey.LiveSecretKey = field.NewString(tableName, "live_secret_key")
	_aPIKey.LivePublicKey = field.NewString(tableName, "live_public_key")
	_aPIKey.OwnerID = field.NewString(tableName, "owner_id")

	_aPIKey.fillFieldMap()

	return _aPIKey
}

type aPIKey struct {
	aPIKeyDo aPIKeyDo

	ALL           field.Asterisk
	ID            field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	TestSecretKey field.String
	TestPublicKey field.String
	LiveSecretKey field.String
	LivePublicKey field.String
	OwnerID       field.String

	fieldMap map[string]field.Expr
}

func (a aPIKey) Table(newTableName string) *aPIKey {
	a.aPIKeyDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aPIKey) As(alias string) *aPIKey {
	a.aPIKeyDo.DO = *(a.aPIKeyDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aPIKey) updateTableName(table string) *aPIKey {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.TestSecretKey = field.NewString(table, "test_secret_key")
	a.TestPublicKey = field.NewString(table, "test_public_key")
	a.LiveSecretKey = field.NewString(table, "live_secret_key")
	a.LivePublicKey = field.NewString(table, "live_public_key")
	a.OwnerID = field.NewString(table, "owner_id")

	a.fillFieldMap()

	return a
}

func (a *aPIKey) WithContext(ctx context.Context) *aPIKeyDo { return a.aPIKeyDo.WithContext(ctx) }

func (a aPIKey) TableName() string { return a.aPIKeyDo.TableName() }

func (a aPIKey) Alias() string { return a.aPIKeyDo.Alias() }

func (a aPIKey) Columns(cols ...field.Expr) gen.Columns { return a.aPIKeyDo.Columns(cols...) }

func (a *aPIKey) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aPIKey) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["test_secret_key"] = a.TestSecretKey
	a.fieldMap["test_public_key"] = a.TestPublicKey
	a.fieldMap["live_secret_key"] = a.LiveSecretKey
	a.fieldMap["live_public_key"] = a.LivePublicKey
	a.fieldMap["owner_id"] = a.OwnerID
}

func (a aPIKey) clone(db *gorm.DB) aPIKey {
	a.aPIKeyDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aPIKey) replaceDB(db *gorm.DB) aPIKey {
	a.aPIKeyDo.ReplaceDB(db)
	return a
}

type aPIKeyDo struct{ gen.DO }

func (a aPIKeyDo) Debug() *aPIKeyDo {
	return a.withDO(a.DO.Debug())
}

func (a aPIKeyDo) WithContext(ctx context.Context) *aPIKeyDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aPIKeyDo) ReadDB() *aPIKeyDo {
	return a.Clauses(dbresolver.Read)
}

func (a aPIKeyDo) WriteDB() *aPIKeyDo {
	return a.Clauses(dbresolver.Write)
}

func (a aPIKeyDo) Session(config *gorm.Session) *aPIKeyDo {
	return a.withDO(a.DO.Session(config))
}

func (a aPIKeyDo) Clauses(conds ...clause.Expression) *aPIKeyDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aPIKeyDo) Returning(value interface{}, columns ...string) *aPIKeyDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aPIKeyDo) Not(conds ...gen.Condition) *aPIKeyDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aPIKeyDo) Or(conds ...gen.Condition) *aPIKeyDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aPIKeyDo) Select(conds ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aPIKeyDo) Where(conds ...gen.Condition) *aPIKeyDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aPIKeyDo) Order(conds ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aPIKeyDo) Distinct(cols ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aPIKeyDo) Omit(cols ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aPIKeyDo) Join(table schema.Tabler, on ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aPIKeyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aPIKeyDo) RightJoin(table schema.Tabler, on ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aPIKeyDo) Group(cols ...field.Expr) *aPIKeyDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aPIKeyDo) Having(conds ...gen.Condition) *aPIKeyDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aPIKeyDo) Limit(limit int) *aPIKeyDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aPIKeyDo) Offset(offset int) *aPIKeyDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aPIKeyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *aPIKeyDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aPIKeyDo) Unscoped() *aPIKeyDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aPIKeyDo) Create(values ...*model.APIKey) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aPIKeyDo) CreateInBatches(values []*model.APIKey, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aPIKeyDo) Save(values ...*model.APIKey) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aPIKeyDo) First() (*model.APIKey, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.APIKey), nil
	}
}

func (a aPIKeyDo) Take() (*model.APIKey, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.APIKey), nil
	}
}

func (a aPIKeyDo) Last() (*model.APIKey, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.APIKey), nil
	}
}

func (a aPIKeyDo) Find() ([]*model.APIKey, error) {
	result, err := a.DO.Find()
	return result.([]*model.APIKey), err
}

func (a aPIKeyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.APIKey, err error) {
	buf := make([]*model.APIKey, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aPIKeyDo) FindInBatches(result *[]*model.APIKey, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aPIKeyDo) Attrs(attrs ...field.AssignExpr) *aPIKeyDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aPIKeyDo) Assign(attrs ...field.AssignExpr) *aPIKeyDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aPIKeyDo) Joins(fields ...field.RelationField) *aPIKeyDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aPIKeyDo) Preload(fields ...field.RelationField) *aPIKeyDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aPIKeyDo) FirstOrInit() (*model.APIKey, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.APIKey), nil
	}
}

func (a aPIKeyDo) FirstOrCreate() (*model.APIKey, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.APIKey), nil
	}
}

func (a aPIKeyDo) FindByPage(offset int, limit int) (result []*model.APIKey, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a aPIKeyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aPIKeyDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aPIKeyDo) Delete(models ...*model.APIKey) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aPIKeyDo) withDO(do gen.Dao) *aPIKeyDo {
	a.DO = *do.(*gen.DO)
	return a
}
