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

func newBlockMetum(db *gorm.DB, opts ...gen.DOOption) blockMetum {
	_blockMetum := blockMetum{}

	_blockMetum.blockMetumDo.UseDB(db, opts...)
	_blockMetum.blockMetumDo.UseModel(&model.BlockMetum{})

	tableName := _blockMetum.blockMetumDo.TableName()
	_blockMetum.ALL = field.NewAsterisk(tableName)
	_blockMetum.ID = field.NewString(tableName, "id")
	_blockMetum.CreatedAt = field.NewTime(tableName, "created_at")
	_blockMetum.UpdatedAt = field.NewTime(tableName, "updated_at")
	_blockMetum.DeletedAt = field.NewField(tableName, "deleted_at")
	_blockMetum.AccountID = field.NewString(tableName, "account_id")
	_blockMetum.BlockTxLimit = field.NewInt32(tableName, "block_tx_limit")
	_blockMetum.TransitionTxID = field.NewString(tableName, "transition_tx_id")
	_blockMetum.OpeningDate = field.NewTime(tableName, "opening_date")
	_blockMetum.ClosingDate = field.NewTime(tableName, "closing_date")

	_blockMetum.fillFieldMap()

	return _blockMetum
}

type blockMetum struct {
	blockMetumDo blockMetumDo

	ALL            field.Asterisk
	ID             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	AccountID      field.String
	BlockTxLimit   field.Int32
	TransitionTxID field.String
	OpeningDate    field.Time
	ClosingDate    field.Time

	fieldMap map[string]field.Expr
}

func (b blockMetum) Table(newTableName string) *blockMetum {
	b.blockMetumDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b blockMetum) As(alias string) *blockMetum {
	b.blockMetumDo.DO = *(b.blockMetumDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *blockMetum) updateTableName(table string) *blockMetum {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewString(table, "id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")
	b.AccountID = field.NewString(table, "account_id")
	b.BlockTxLimit = field.NewInt32(table, "block_tx_limit")
	b.TransitionTxID = field.NewString(table, "transition_tx_id")
	b.OpeningDate = field.NewTime(table, "opening_date")
	b.ClosingDate = field.NewTime(table, "closing_date")

	b.fillFieldMap()

	return b
}

func (b *blockMetum) WithContext(ctx context.Context) *blockMetumDo {
	return b.blockMetumDo.WithContext(ctx)
}

func (b blockMetum) TableName() string { return b.blockMetumDo.TableName() }

func (b blockMetum) Alias() string { return b.blockMetumDo.Alias() }

func (b blockMetum) Columns(cols ...field.Expr) gen.Columns { return b.blockMetumDo.Columns(cols...) }

func (b *blockMetum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *blockMetum) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 9)
	b.fieldMap["id"] = b.ID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["account_id"] = b.AccountID
	b.fieldMap["block_tx_limit"] = b.BlockTxLimit
	b.fieldMap["transition_tx_id"] = b.TransitionTxID
	b.fieldMap["opening_date"] = b.OpeningDate
	b.fieldMap["closing_date"] = b.ClosingDate
}

func (b blockMetum) clone(db *gorm.DB) blockMetum {
	b.blockMetumDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b blockMetum) replaceDB(db *gorm.DB) blockMetum {
	b.blockMetumDo.ReplaceDB(db)
	return b
}

type blockMetumDo struct{ gen.DO }

func (b blockMetumDo) Debug() *blockMetumDo {
	return b.withDO(b.DO.Debug())
}

func (b blockMetumDo) WithContext(ctx context.Context) *blockMetumDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b blockMetumDo) ReadDB() *blockMetumDo {
	return b.Clauses(dbresolver.Read)
}

func (b blockMetumDo) WriteDB() *blockMetumDo {
	return b.Clauses(dbresolver.Write)
}

func (b blockMetumDo) Session(config *gorm.Session) *blockMetumDo {
	return b.withDO(b.DO.Session(config))
}

func (b blockMetumDo) Clauses(conds ...clause.Expression) *blockMetumDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b blockMetumDo) Returning(value interface{}, columns ...string) *blockMetumDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b blockMetumDo) Not(conds ...gen.Condition) *blockMetumDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b blockMetumDo) Or(conds ...gen.Condition) *blockMetumDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b blockMetumDo) Select(conds ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b blockMetumDo) Where(conds ...gen.Condition) *blockMetumDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b blockMetumDo) Order(conds ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b blockMetumDo) Distinct(cols ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b blockMetumDo) Omit(cols ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b blockMetumDo) Join(table schema.Tabler, on ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b blockMetumDo) LeftJoin(table schema.Tabler, on ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b blockMetumDo) RightJoin(table schema.Tabler, on ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b blockMetumDo) Group(cols ...field.Expr) *blockMetumDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b blockMetumDo) Having(conds ...gen.Condition) *blockMetumDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b blockMetumDo) Limit(limit int) *blockMetumDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b blockMetumDo) Offset(offset int) *blockMetumDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b blockMetumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *blockMetumDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b blockMetumDo) Unscoped() *blockMetumDo {
	return b.withDO(b.DO.Unscoped())
}

func (b blockMetumDo) Create(values ...*model.BlockMetum) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b blockMetumDo) CreateInBatches(values []*model.BlockMetum, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b blockMetumDo) Save(values ...*model.BlockMetum) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b blockMetumDo) First() (*model.BlockMetum, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BlockMetum), nil
	}
}

func (b blockMetumDo) Take() (*model.BlockMetum, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BlockMetum), nil
	}
}

func (b blockMetumDo) Last() (*model.BlockMetum, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BlockMetum), nil
	}
}

func (b blockMetumDo) Find() ([]*model.BlockMetum, error) {
	result, err := b.DO.Find()
	return result.([]*model.BlockMetum), err
}

func (b blockMetumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BlockMetum, err error) {
	buf := make([]*model.BlockMetum, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b blockMetumDo) FindInBatches(result *[]*model.BlockMetum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b blockMetumDo) Attrs(attrs ...field.AssignExpr) *blockMetumDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b blockMetumDo) Assign(attrs ...field.AssignExpr) *blockMetumDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b blockMetumDo) Joins(fields ...field.RelationField) *blockMetumDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b blockMetumDo) Preload(fields ...field.RelationField) *blockMetumDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b blockMetumDo) FirstOrInit() (*model.BlockMetum, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BlockMetum), nil
	}
}

func (b blockMetumDo) FirstOrCreate() (*model.BlockMetum, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BlockMetum), nil
	}
}

func (b blockMetumDo) FindByPage(offset int, limit int) (result []*model.BlockMetum, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b blockMetumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b blockMetumDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b blockMetumDo) Delete(models ...*model.BlockMetum) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *blockMetumDo) withDO(do gen.Dao) *blockMetumDo {
	b.DO = *do.(*gen.DO)
	return b
}
