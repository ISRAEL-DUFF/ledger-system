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

func newGooseDbVersion(db *gorm.DB, opts ...gen.DOOption) gooseDbVersion {
	_gooseDbVersion := gooseDbVersion{}

	_gooseDbVersion.gooseDbVersionDo.UseDB(db, opts...)
	_gooseDbVersion.gooseDbVersionDo.UseModel(&model.GooseDbVersion{})

	tableName := _gooseDbVersion.gooseDbVersionDo.TableName()
	_gooseDbVersion.ALL = field.NewAsterisk(tableName)
	_gooseDbVersion.ID = field.NewInt32(tableName, "id")
	_gooseDbVersion.VersionID = field.NewInt64(tableName, "version_id")
	_gooseDbVersion.IsApplied = field.NewBool(tableName, "is_applied")
	_gooseDbVersion.Tstamp = field.NewTime(tableName, "tstamp")

	_gooseDbVersion.fillFieldMap()

	return _gooseDbVersion
}

type gooseDbVersion struct {
	gooseDbVersionDo gooseDbVersionDo

	ALL       field.Asterisk
	ID        field.Int32
	VersionID field.Int64
	IsApplied field.Bool
	Tstamp    field.Time

	fieldMap map[string]field.Expr
}

func (g gooseDbVersion) Table(newTableName string) *gooseDbVersion {
	g.gooseDbVersionDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gooseDbVersion) As(alias string) *gooseDbVersion {
	g.gooseDbVersionDo.DO = *(g.gooseDbVersionDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gooseDbVersion) updateTableName(table string) *gooseDbVersion {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.VersionID = field.NewInt64(table, "version_id")
	g.IsApplied = field.NewBool(table, "is_applied")
	g.Tstamp = field.NewTime(table, "tstamp")

	g.fillFieldMap()

	return g
}

func (g *gooseDbVersion) WithContext(ctx context.Context) *gooseDbVersionDo {
	return g.gooseDbVersionDo.WithContext(ctx)
}

func (g gooseDbVersion) TableName() string { return g.gooseDbVersionDo.TableName() }

func (g gooseDbVersion) Alias() string { return g.gooseDbVersionDo.Alias() }

func (g *gooseDbVersion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gooseDbVersion) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 4)
	g.fieldMap["id"] = g.ID
	g.fieldMap["version_id"] = g.VersionID
	g.fieldMap["is_applied"] = g.IsApplied
	g.fieldMap["tstamp"] = g.Tstamp
}

func (g gooseDbVersion) clone(db *gorm.DB) gooseDbVersion {
	g.gooseDbVersionDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gooseDbVersion) replaceDB(db *gorm.DB) gooseDbVersion {
	g.gooseDbVersionDo.ReplaceDB(db)
	return g
}

type gooseDbVersionDo struct{ gen.DO }

func (g gooseDbVersionDo) Debug() *gooseDbVersionDo {
	return g.withDO(g.DO.Debug())
}

func (g gooseDbVersionDo) WithContext(ctx context.Context) *gooseDbVersionDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gooseDbVersionDo) ReadDB() *gooseDbVersionDo {
	return g.Clauses(dbresolver.Read)
}

func (g gooseDbVersionDo) WriteDB() *gooseDbVersionDo {
	return g.Clauses(dbresolver.Write)
}

func (g gooseDbVersionDo) Session(config *gorm.Session) *gooseDbVersionDo {
	return g.withDO(g.DO.Session(config))
}

func (g gooseDbVersionDo) Clauses(conds ...clause.Expression) *gooseDbVersionDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gooseDbVersionDo) Returning(value interface{}, columns ...string) *gooseDbVersionDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gooseDbVersionDo) Not(conds ...gen.Condition) *gooseDbVersionDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gooseDbVersionDo) Or(conds ...gen.Condition) *gooseDbVersionDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gooseDbVersionDo) Select(conds ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gooseDbVersionDo) Where(conds ...gen.Condition) *gooseDbVersionDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gooseDbVersionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *gooseDbVersionDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g gooseDbVersionDo) Order(conds ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gooseDbVersionDo) Distinct(cols ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gooseDbVersionDo) Omit(cols ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gooseDbVersionDo) Join(table schema.Tabler, on ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gooseDbVersionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gooseDbVersionDo) RightJoin(table schema.Tabler, on ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gooseDbVersionDo) Group(cols ...field.Expr) *gooseDbVersionDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gooseDbVersionDo) Having(conds ...gen.Condition) *gooseDbVersionDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gooseDbVersionDo) Limit(limit int) *gooseDbVersionDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gooseDbVersionDo) Offset(offset int) *gooseDbVersionDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gooseDbVersionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *gooseDbVersionDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gooseDbVersionDo) Unscoped() *gooseDbVersionDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gooseDbVersionDo) Create(values ...*model.GooseDbVersion) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gooseDbVersionDo) CreateInBatches(values []*model.GooseDbVersion, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gooseDbVersionDo) Save(values ...*model.GooseDbVersion) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gooseDbVersionDo) First() (*model.GooseDbVersion, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GooseDbVersion), nil
	}
}

func (g gooseDbVersionDo) Take() (*model.GooseDbVersion, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GooseDbVersion), nil
	}
}

func (g gooseDbVersionDo) Last() (*model.GooseDbVersion, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GooseDbVersion), nil
	}
}

func (g gooseDbVersionDo) Find() ([]*model.GooseDbVersion, error) {
	result, err := g.DO.Find()
	return result.([]*model.GooseDbVersion), err
}

func (g gooseDbVersionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GooseDbVersion, err error) {
	buf := make([]*model.GooseDbVersion, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gooseDbVersionDo) FindInBatches(result *[]*model.GooseDbVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gooseDbVersionDo) Attrs(attrs ...field.AssignExpr) *gooseDbVersionDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gooseDbVersionDo) Assign(attrs ...field.AssignExpr) *gooseDbVersionDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gooseDbVersionDo) Joins(fields ...field.RelationField) *gooseDbVersionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gooseDbVersionDo) Preload(fields ...field.RelationField) *gooseDbVersionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gooseDbVersionDo) FirstOrInit() (*model.GooseDbVersion, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GooseDbVersion), nil
	}
}

func (g gooseDbVersionDo) FirstOrCreate() (*model.GooseDbVersion, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GooseDbVersion), nil
	}
}

func (g gooseDbVersionDo) FindByPage(offset int, limit int) (result []*model.GooseDbVersion, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gooseDbVersionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gooseDbVersionDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gooseDbVersionDo) Delete(models ...*model.GooseDbVersion) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gooseDbVersionDo) withDO(do gen.Dao) *gooseDbVersionDo {
	g.DO = *do.(*gen.DO)
	return g
}
