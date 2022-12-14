// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/chenhaonan-eth/dao/dal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newBondZhUsRate(db *gorm.DB, opts ...gen.DOOption) bondZhUsRate {
	_bondZhUsRate := bondZhUsRate{}

	_bondZhUsRate.bondZhUsRateDo.UseDB(db, opts...)
	_bondZhUsRate.bondZhUsRateDo.UseModel(&model.BondZhUsRate{})

	tableName := _bondZhUsRate.bondZhUsRateDo.TableName()
	_bondZhUsRate.ALL = field.NewAsterisk(tableName)
	_bondZhUsRate.Date = field.NewString(tableName, "date")
	_bondZhUsRate.CN2Years = field.NewFloat64(tableName, "cn2_years")
	_bondZhUsRate.CN5Years = field.NewFloat64(tableName, "cn5_years")
	_bondZhUsRate.CN10Years = field.NewFloat64(tableName, "cn10_years")
	_bondZhUsRate.CN30Years = field.NewFloat64(tableName, "cn30_years")
	_bondZhUsRate.CN10_2Years = field.NewFloat64(tableName, "cn10_2_years")
	_bondZhUsRate.US2Years = field.NewFloat64(tableName, "us2_years")
	_bondZhUsRate.US5Years = field.NewFloat64(tableName, "us5_years")
	_bondZhUsRate.US10Years = field.NewFloat64(tableName, "us10_years")
	_bondZhUsRate.US30Years = field.NewFloat64(tableName, "us30_years")
	_bondZhUsRate.US10_2Years = field.NewFloat64(tableName, "us10_2_years")

	_bondZhUsRate.fillFieldMap()

	return _bondZhUsRate
}

type bondZhUsRate struct {
	bondZhUsRateDo bondZhUsRateDo

	ALL         field.Asterisk
	Date        field.String
	CN2Years    field.Float64
	CN5Years    field.Float64
	CN10Years   field.Float64
	CN30Years   field.Float64
	CN10_2Years field.Float64
	US2Years    field.Float64
	US5Years    field.Float64
	US10Years   field.Float64
	US30Years   field.Float64
	US10_2Years field.Float64

	fieldMap map[string]field.Expr
}

func (b bondZhUsRate) Table(newTableName string) *bondZhUsRate {
	b.bondZhUsRateDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bondZhUsRate) As(alias string) *bondZhUsRate {
	b.bondZhUsRateDo.DO = *(b.bondZhUsRateDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bondZhUsRate) updateTableName(table string) *bondZhUsRate {
	b.ALL = field.NewAsterisk(table)
	b.Date = field.NewString(table, "date")
	b.CN2Years = field.NewFloat64(table, "cn2_years")
	b.CN5Years = field.NewFloat64(table, "cn5_years")
	b.CN10Years = field.NewFloat64(table, "cn10_years")
	b.CN30Years = field.NewFloat64(table, "cn30_years")
	b.CN10_2Years = field.NewFloat64(table, "cn10_2_years")
	b.US2Years = field.NewFloat64(table, "us2_years")
	b.US5Years = field.NewFloat64(table, "us5_years")
	b.US10Years = field.NewFloat64(table, "us10_years")
	b.US30Years = field.NewFloat64(table, "us30_years")
	b.US10_2Years = field.NewFloat64(table, "us10_2_years")

	b.fillFieldMap()

	return b
}

func (b *bondZhUsRate) WithContext(ctx context.Context) *bondZhUsRateDo {
	return b.bondZhUsRateDo.WithContext(ctx)
}

func (b bondZhUsRate) TableName() string { return b.bondZhUsRateDo.TableName() }

func (b bondZhUsRate) Alias() string { return b.bondZhUsRateDo.Alias() }

func (b *bondZhUsRate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bondZhUsRate) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 11)
	b.fieldMap["date"] = b.Date
	b.fieldMap["cn2_years"] = b.CN2Years
	b.fieldMap["cn5_years"] = b.CN5Years
	b.fieldMap["cn10_years"] = b.CN10Years
	b.fieldMap["cn30_years"] = b.CN30Years
	b.fieldMap["cn10_2_years"] = b.CN10_2Years
	b.fieldMap["us2_years"] = b.US2Years
	b.fieldMap["us5_years"] = b.US5Years
	b.fieldMap["us10_years"] = b.US10Years
	b.fieldMap["us30_years"] = b.US30Years
	b.fieldMap["us10_2_years"] = b.US10_2Years
}

func (b bondZhUsRate) clone(db *gorm.DB) bondZhUsRate {
	b.bondZhUsRateDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bondZhUsRate) replaceDB(db *gorm.DB) bondZhUsRate {
	b.bondZhUsRateDo.ReplaceDB(db)
	return b
}

type bondZhUsRateDo struct{ gen.DO }

func (b bondZhUsRateDo) Debug() *bondZhUsRateDo {
	return b.withDO(b.DO.Debug())
}

func (b bondZhUsRateDo) WithContext(ctx context.Context) *bondZhUsRateDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bondZhUsRateDo) ReadDB() *bondZhUsRateDo {
	return b.Clauses(dbresolver.Read)
}

func (b bondZhUsRateDo) WriteDB() *bondZhUsRateDo {
	return b.Clauses(dbresolver.Write)
}

func (b bondZhUsRateDo) Session(config *gorm.Session) *bondZhUsRateDo {
	return b.withDO(b.DO.Session(config))
}

func (b bondZhUsRateDo) Clauses(conds ...clause.Expression) *bondZhUsRateDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bondZhUsRateDo) Returning(value interface{}, columns ...string) *bondZhUsRateDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bondZhUsRateDo) Not(conds ...gen.Condition) *bondZhUsRateDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bondZhUsRateDo) Or(conds ...gen.Condition) *bondZhUsRateDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bondZhUsRateDo) Select(conds ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bondZhUsRateDo) Where(conds ...gen.Condition) *bondZhUsRateDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bondZhUsRateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *bondZhUsRateDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b bondZhUsRateDo) Order(conds ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bondZhUsRateDo) Distinct(cols ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bondZhUsRateDo) Omit(cols ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bondZhUsRateDo) Join(table schema.Tabler, on ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bondZhUsRateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bondZhUsRateDo) RightJoin(table schema.Tabler, on ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bondZhUsRateDo) Group(cols ...field.Expr) *bondZhUsRateDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bondZhUsRateDo) Having(conds ...gen.Condition) *bondZhUsRateDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bondZhUsRateDo) Limit(limit int) *bondZhUsRateDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bondZhUsRateDo) Offset(offset int) *bondZhUsRateDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bondZhUsRateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bondZhUsRateDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bondZhUsRateDo) Unscoped() *bondZhUsRateDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bondZhUsRateDo) Create(values ...*model.BondZhUsRate) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bondZhUsRateDo) CreateInBatches(values []*model.BondZhUsRate, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bondZhUsRateDo) Save(values ...*model.BondZhUsRate) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bondZhUsRateDo) First() (*model.BondZhUsRate, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BondZhUsRate), nil
	}
}

func (b bondZhUsRateDo) Take() (*model.BondZhUsRate, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BondZhUsRate), nil
	}
}

func (b bondZhUsRateDo) Last() (*model.BondZhUsRate, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BondZhUsRate), nil
	}
}

func (b bondZhUsRateDo) Find() ([]*model.BondZhUsRate, error) {
	result, err := b.DO.Find()
	return result.([]*model.BondZhUsRate), err
}

func (b bondZhUsRateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BondZhUsRate, err error) {
	buf := make([]*model.BondZhUsRate, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bondZhUsRateDo) FindInBatches(result *[]*model.BondZhUsRate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bondZhUsRateDo) Attrs(attrs ...field.AssignExpr) *bondZhUsRateDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bondZhUsRateDo) Assign(attrs ...field.AssignExpr) *bondZhUsRateDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bondZhUsRateDo) Joins(fields ...field.RelationField) *bondZhUsRateDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bondZhUsRateDo) Preload(fields ...field.RelationField) *bondZhUsRateDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bondZhUsRateDo) FirstOrInit() (*model.BondZhUsRate, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BondZhUsRate), nil
	}
}

func (b bondZhUsRateDo) FirstOrCreate() (*model.BondZhUsRate, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BondZhUsRate), nil
	}
}

func (b bondZhUsRateDo) FindByPage(offset int, limit int) (result []*model.BondZhUsRate, count int64, err error) {
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

func (b bondZhUsRateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bondZhUsRateDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bondZhUsRateDo) Delete(models ...*model.BondZhUsRate) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bondZhUsRateDo) withDO(do gen.Dao) *bondZhUsRateDo {
	b.DO = *do.(*gen.DO)
	return b
}
