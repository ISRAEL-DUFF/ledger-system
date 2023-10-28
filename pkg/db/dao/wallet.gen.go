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

func newWallet(db *gorm.DB, opts ...gen.DOOption) wallet {
	_wallet := wallet{}

	_wallet.walletDo.UseDB(db, opts...)
	_wallet.walletDo.UseModel(&model.Wallet{})

	tableName := _wallet.walletDo.TableName()
	_wallet.ALL = field.NewAsterisk(tableName)
	_wallet.ID = field.NewString(tableName, "id")
	_wallet.CreatedAt = field.NewTime(tableName, "created_at")
	_wallet.UpdatedAt = field.NewTime(tableName, "updated_at")
	_wallet.DeletedAt = field.NewField(tableName, "deleted_at")
	_wallet.Name = field.NewString(tableName, "name")
	_wallet.Type = field.NewString(tableName, "type")
	_wallet.AccountNumber = field.NewString(tableName, "account_number")
	_wallet.LedgerAccounts = field.NewString(tableName, "ledger_accounts")
	_wallet.OwnerID = field.NewString(tableName, "owner_id")

	_wallet.fillFieldMap()

	return _wallet
}

type wallet struct {
	walletDo walletDo

	ALL            field.Asterisk
	ID             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	Name           field.String
	Type           field.String
	AccountNumber  field.String
	LedgerAccounts field.String
	OwnerID        field.String

	fieldMap map[string]field.Expr
}

func (w wallet) Table(newTableName string) *wallet {
	w.walletDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wallet) As(alias string) *wallet {
	w.walletDo.DO = *(w.walletDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wallet) updateTableName(table string) *wallet {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewString(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.Name = field.NewString(table, "name")
	w.Type = field.NewString(table, "type")
	w.AccountNumber = field.NewString(table, "account_number")
	w.LedgerAccounts = field.NewString(table, "ledger_accounts")
	w.OwnerID = field.NewString(table, "owner_id")

	w.fillFieldMap()

	return w
}

func (w *wallet) WithContext(ctx context.Context) *walletDo { return w.walletDo.WithContext(ctx) }

func (w wallet) TableName() string { return w.walletDo.TableName() }

func (w wallet) Alias() string { return w.walletDo.Alias() }

func (w wallet) Columns(cols ...field.Expr) gen.Columns { return w.walletDo.Columns(cols...) }

func (w *wallet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wallet) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["name"] = w.Name
	w.fieldMap["type"] = w.Type
	w.fieldMap["account_number"] = w.AccountNumber
	w.fieldMap["ledger_accounts"] = w.LedgerAccounts
	w.fieldMap["owner_id"] = w.OwnerID
}

func (w wallet) clone(db *gorm.DB) wallet {
	w.walletDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wallet) replaceDB(db *gorm.DB) wallet {
	w.walletDo.ReplaceDB(db)
	return w
}

type walletDo struct{ gen.DO }

func (w walletDo) Debug() *walletDo {
	return w.withDO(w.DO.Debug())
}

func (w walletDo) WithContext(ctx context.Context) *walletDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w walletDo) ReadDB() *walletDo {
	return w.Clauses(dbresolver.Read)
}

func (w walletDo) WriteDB() *walletDo {
	return w.Clauses(dbresolver.Write)
}

func (w walletDo) Session(config *gorm.Session) *walletDo {
	return w.withDO(w.DO.Session(config))
}

func (w walletDo) Clauses(conds ...clause.Expression) *walletDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w walletDo) Returning(value interface{}, columns ...string) *walletDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w walletDo) Not(conds ...gen.Condition) *walletDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w walletDo) Or(conds ...gen.Condition) *walletDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w walletDo) Select(conds ...field.Expr) *walletDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w walletDo) Where(conds ...gen.Condition) *walletDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w walletDo) Order(conds ...field.Expr) *walletDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w walletDo) Distinct(cols ...field.Expr) *walletDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w walletDo) Omit(cols ...field.Expr) *walletDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w walletDo) Join(table schema.Tabler, on ...field.Expr) *walletDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w walletDo) LeftJoin(table schema.Tabler, on ...field.Expr) *walletDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w walletDo) RightJoin(table schema.Tabler, on ...field.Expr) *walletDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w walletDo) Group(cols ...field.Expr) *walletDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w walletDo) Having(conds ...gen.Condition) *walletDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w walletDo) Limit(limit int) *walletDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w walletDo) Offset(offset int) *walletDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w walletDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *walletDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w walletDo) Unscoped() *walletDo {
	return w.withDO(w.DO.Unscoped())
}

func (w walletDo) Create(values ...*model.Wallet) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w walletDo) CreateInBatches(values []*model.Wallet, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w walletDo) Save(values ...*model.Wallet) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w walletDo) First() (*model.Wallet, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) Take() (*model.Wallet, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) Last() (*model.Wallet, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) Find() ([]*model.Wallet, error) {
	result, err := w.DO.Find()
	return result.([]*model.Wallet), err
}

func (w walletDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Wallet, err error) {
	buf := make([]*model.Wallet, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w walletDo) FindInBatches(result *[]*model.Wallet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w walletDo) Attrs(attrs ...field.AssignExpr) *walletDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w walletDo) Assign(attrs ...field.AssignExpr) *walletDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w walletDo) Joins(fields ...field.RelationField) *walletDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w walletDo) Preload(fields ...field.RelationField) *walletDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w walletDo) FirstOrInit() (*model.Wallet, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) FirstOrCreate() (*model.Wallet, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Wallet), nil
	}
}

func (w walletDo) FindByPage(offset int, limit int) (result []*model.Wallet, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w walletDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w walletDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w walletDo) Delete(models ...*model.Wallet) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *walletDo) withDO(do gen.Dao) *walletDo {
	w.DO = *do.(*gen.DO)
	return w
}