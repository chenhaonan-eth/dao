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

func newChinaPPI(db *gorm.DB, opts ...gen.DOOption) chinaPPI {
	_chinaPPI := chinaPPI{}

	_chinaPPI.chinaPPIDo.UseDB(db, opts...)
	_chinaPPI.chinaPPIDo.UseModel(&model.ChinaPPI{})

	tableName := _chinaPPI.chinaPPIDo.TableName()
	_chinaPPI.ALL = field.NewAsterisk(tableName)
	_chinaPPI.Date = field.NewString(tableName, "date")
	_chinaPPI.Time = field.NewString(tableName, "time")
	_chinaPPI.PPI = field.NewFloat64(tableName, "ppi")
	_chinaPPI.YearOnYear = field.NewFloat64(tableName, "year_on_year")
	_chinaPPI.Accumulative = field.NewFloat64(tableName, "accumulative")

	_chinaPPI.fillFieldMap()

	return _chinaPPI
}

type chinaPPI struct {
	chinaPPIDo chinaPPIDo

	ALL          field.Asterisk
	Date         field.String
	Time         field.String
	PPI          field.Float64
	YearOnYear   field.Float64
	Accumulative field.Float64

	fieldMap map[string]field.Expr
}

func (c chinaPPI) Table(newTableName string) *chinaPPI {
	c.chinaPPIDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chinaPPI) As(alias string) *chinaPPI {
	c.chinaPPIDo.DO = *(c.chinaPPIDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chinaPPI) updateTableName(table string) *chinaPPI {
	c.ALL = field.NewAsterisk(table)
	c.Date = field.NewString(table, "date")
	c.Time = field.NewString(table, "time")
	c.PPI = field.NewFloat64(table, "ppi")
	c.YearOnYear = field.NewFloat64(table, "year_on_year")
	c.Accumulative = field.NewFloat64(table, "accumulative")

	c.fillFieldMap()

	return c
}

func (c *chinaPPI) WithContext(ctx context.Context) *chinaPPIDo { return c.chinaPPIDo.WithContext(ctx) }

func (c chinaPPI) TableName() string { return c.chinaPPIDo.TableName() }

func (c chinaPPI) Alias() string { return c.chinaPPIDo.Alias() }

func (c *chinaPPI) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chinaPPI) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["date"] = c.Date
	c.fieldMap["time"] = c.Time
	c.fieldMap["ppi"] = c.PPI
	c.fieldMap["year_on_year"] = c.YearOnYear
	c.fieldMap["accumulative"] = c.Accumulative
}

func (c chinaPPI) clone(db *gorm.DB) chinaPPI {
	c.chinaPPIDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chinaPPI) replaceDB(db *gorm.DB) chinaPPI {
	c.chinaPPIDo.ReplaceDB(db)
	return c
}

type chinaPPIDo struct{ gen.DO }

func (c chinaPPIDo) Debug() *chinaPPIDo {
	return c.withDO(c.DO.Debug())
}

func (c chinaPPIDo) WithContext(ctx context.Context) *chinaPPIDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chinaPPIDo) ReadDB() *chinaPPIDo {
	return c.Clauses(dbresolver.Read)
}

func (c chinaPPIDo) WriteDB() *chinaPPIDo {
	return c.Clauses(dbresolver.Write)
}

func (c chinaPPIDo) Session(config *gorm.Session) *chinaPPIDo {
	return c.withDO(c.DO.Session(config))
}

func (c chinaPPIDo) Clauses(conds ...clause.Expression) *chinaPPIDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chinaPPIDo) Returning(value interface{}, columns ...string) *chinaPPIDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chinaPPIDo) Not(conds ...gen.Condition) *chinaPPIDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chinaPPIDo) Or(conds ...gen.Condition) *chinaPPIDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chinaPPIDo) Select(conds ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chinaPPIDo) Where(conds ...gen.Condition) *chinaPPIDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chinaPPIDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *chinaPPIDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chinaPPIDo) Order(conds ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chinaPPIDo) Distinct(cols ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chinaPPIDo) Omit(cols ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chinaPPIDo) Join(table schema.Tabler, on ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chinaPPIDo) LeftJoin(table schema.Tabler, on ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chinaPPIDo) RightJoin(table schema.Tabler, on ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chinaPPIDo) Group(cols ...field.Expr) *chinaPPIDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chinaPPIDo) Having(conds ...gen.Condition) *chinaPPIDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chinaPPIDo) Limit(limit int) *chinaPPIDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chinaPPIDo) Offset(offset int) *chinaPPIDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chinaPPIDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *chinaPPIDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chinaPPIDo) Unscoped() *chinaPPIDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chinaPPIDo) Create(values ...*model.ChinaPPI) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chinaPPIDo) CreateInBatches(values []*model.ChinaPPI, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chinaPPIDo) Save(values ...*model.ChinaPPI) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chinaPPIDo) First() (*model.ChinaPPI, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaPPI), nil
	}
}

func (c chinaPPIDo) Take() (*model.ChinaPPI, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaPPI), nil
	}
}

func (c chinaPPIDo) Last() (*model.ChinaPPI, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaPPI), nil
	}
}

func (c chinaPPIDo) Find() ([]*model.ChinaPPI, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChinaPPI), err
}

func (c chinaPPIDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChinaPPI, err error) {
	buf := make([]*model.ChinaPPI, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chinaPPIDo) FindInBatches(result *[]*model.ChinaPPI, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chinaPPIDo) Attrs(attrs ...field.AssignExpr) *chinaPPIDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chinaPPIDo) Assign(attrs ...field.AssignExpr) *chinaPPIDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chinaPPIDo) Joins(fields ...field.RelationField) *chinaPPIDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chinaPPIDo) Preload(fields ...field.RelationField) *chinaPPIDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chinaPPIDo) FirstOrInit() (*model.ChinaPPI, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaPPI), nil
	}
}

func (c chinaPPIDo) FirstOrCreate() (*model.ChinaPPI, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaPPI), nil
	}
}

func (c chinaPPIDo) FindByPage(offset int, limit int) (result []*model.ChinaPPI, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c chinaPPIDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chinaPPIDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chinaPPIDo) Delete(models ...*model.ChinaPPI) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chinaPPIDo) withDO(do gen.Dao) *chinaPPIDo {
	c.DO = *do.(*gen.DO)
	return c
}
