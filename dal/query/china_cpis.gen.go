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

func newChinaCPI(db *gorm.DB, opts ...gen.DOOption) chinaCPI {
	_chinaCPI := chinaCPI{}

	_chinaCPI.chinaCPIDo.UseDB(db, opts...)
	_chinaCPI.chinaCPIDo.UseModel(&model.ChinaCPI{})

	tableName := _chinaCPI.chinaCPIDo.TableName()
	_chinaCPI.ALL = field.NewAsterisk(tableName)
	_chinaCPI.Date = field.NewString(tableName, "date")
	_chinaCPI.Time = field.NewString(tableName, "time")
	_chinaCPI.National = field.NewFloat64(tableName, "national")
	_chinaCPI.NationalYearOnYear = field.NewFloat64(tableName, "national_year_on_year")
	_chinaCPI.NationalYearOverYear = field.NewFloat64(tableName, "national_year_over_year")
	_chinaCPI.NationalAccumulative = field.NewFloat64(tableName, "national_accumulative")
	_chinaCPI.City = field.NewFloat64(tableName, "city")
	_chinaCPI.CityYearOnYear = field.NewFloat64(tableName, "city_year_on_year")
	_chinaCPI.CityYearOverYear = field.NewFloat64(tableName, "city_year_over_year")
	_chinaCPI.CityAccumulative = field.NewFloat64(tableName, "city_accumulative")
	_chinaCPI.Rural = field.NewFloat64(tableName, "rural")
	_chinaCPI.RuralYearOnYear = field.NewFloat64(tableName, "rural_year_on_year")
	_chinaCPI.RuralYearOverYear = field.NewFloat64(tableName, "rural_year_over_year")
	_chinaCPI.RuralAccumulative = field.NewFloat64(tableName, "rural_accumulative")

	_chinaCPI.fillFieldMap()

	return _chinaCPI
}

type chinaCPI struct {
	chinaCPIDo chinaCPIDo

	ALL                  field.Asterisk
	Date                 field.String
	Time                 field.String
	National             field.Float64
	NationalYearOnYear   field.Float64
	NationalYearOverYear field.Float64
	NationalAccumulative field.Float64
	City                 field.Float64
	CityYearOnYear       field.Float64
	CityYearOverYear     field.Float64
	CityAccumulative     field.Float64
	Rural                field.Float64
	RuralYearOnYear      field.Float64
	RuralYearOverYear    field.Float64
	RuralAccumulative    field.Float64

	fieldMap map[string]field.Expr
}

func (c chinaCPI) Table(newTableName string) *chinaCPI {
	c.chinaCPIDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chinaCPI) As(alias string) *chinaCPI {
	c.chinaCPIDo.DO = *(c.chinaCPIDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chinaCPI) updateTableName(table string) *chinaCPI {
	c.ALL = field.NewAsterisk(table)
	c.Date = field.NewString(table, "date")
	c.Time = field.NewString(table, "time")
	c.National = field.NewFloat64(table, "national")
	c.NationalYearOnYear = field.NewFloat64(table, "national_year_on_year")
	c.NationalYearOverYear = field.NewFloat64(table, "national_year_over_year")
	c.NationalAccumulative = field.NewFloat64(table, "national_accumulative")
	c.City = field.NewFloat64(table, "city")
	c.CityYearOnYear = field.NewFloat64(table, "city_year_on_year")
	c.CityYearOverYear = field.NewFloat64(table, "city_year_over_year")
	c.CityAccumulative = field.NewFloat64(table, "city_accumulative")
	c.Rural = field.NewFloat64(table, "rural")
	c.RuralYearOnYear = field.NewFloat64(table, "rural_year_on_year")
	c.RuralYearOverYear = field.NewFloat64(table, "rural_year_over_year")
	c.RuralAccumulative = field.NewFloat64(table, "rural_accumulative")

	c.fillFieldMap()

	return c
}

func (c *chinaCPI) WithContext(ctx context.Context) *chinaCPIDo { return c.chinaCPIDo.WithContext(ctx) }

func (c chinaCPI) TableName() string { return c.chinaCPIDo.TableName() }

func (c chinaCPI) Alias() string { return c.chinaCPIDo.Alias() }

func (c *chinaCPI) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chinaCPI) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 14)
	c.fieldMap["date"] = c.Date
	c.fieldMap["time"] = c.Time
	c.fieldMap["national"] = c.National
	c.fieldMap["national_year_on_year"] = c.NationalYearOnYear
	c.fieldMap["national_year_over_year"] = c.NationalYearOverYear
	c.fieldMap["national_accumulative"] = c.NationalAccumulative
	c.fieldMap["city"] = c.City
	c.fieldMap["city_year_on_year"] = c.CityYearOnYear
	c.fieldMap["city_year_over_year"] = c.CityYearOverYear
	c.fieldMap["city_accumulative"] = c.CityAccumulative
	c.fieldMap["rural"] = c.Rural
	c.fieldMap["rural_year_on_year"] = c.RuralYearOnYear
	c.fieldMap["rural_year_over_year"] = c.RuralYearOverYear
	c.fieldMap["rural_accumulative"] = c.RuralAccumulative
}

func (c chinaCPI) clone(db *gorm.DB) chinaCPI {
	c.chinaCPIDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chinaCPI) replaceDB(db *gorm.DB) chinaCPI {
	c.chinaCPIDo.ReplaceDB(db)
	return c
}

type chinaCPIDo struct{ gen.DO }

func (c chinaCPIDo) Debug() *chinaCPIDo {
	return c.withDO(c.DO.Debug())
}

func (c chinaCPIDo) WithContext(ctx context.Context) *chinaCPIDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chinaCPIDo) ReadDB() *chinaCPIDo {
	return c.Clauses(dbresolver.Read)
}

func (c chinaCPIDo) WriteDB() *chinaCPIDo {
	return c.Clauses(dbresolver.Write)
}

func (c chinaCPIDo) Session(config *gorm.Session) *chinaCPIDo {
	return c.withDO(c.DO.Session(config))
}

func (c chinaCPIDo) Clauses(conds ...clause.Expression) *chinaCPIDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chinaCPIDo) Returning(value interface{}, columns ...string) *chinaCPIDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chinaCPIDo) Not(conds ...gen.Condition) *chinaCPIDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chinaCPIDo) Or(conds ...gen.Condition) *chinaCPIDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chinaCPIDo) Select(conds ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chinaCPIDo) Where(conds ...gen.Condition) *chinaCPIDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chinaCPIDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *chinaCPIDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chinaCPIDo) Order(conds ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chinaCPIDo) Distinct(cols ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chinaCPIDo) Omit(cols ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chinaCPIDo) Join(table schema.Tabler, on ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chinaCPIDo) LeftJoin(table schema.Tabler, on ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chinaCPIDo) RightJoin(table schema.Tabler, on ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chinaCPIDo) Group(cols ...field.Expr) *chinaCPIDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chinaCPIDo) Having(conds ...gen.Condition) *chinaCPIDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chinaCPIDo) Limit(limit int) *chinaCPIDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chinaCPIDo) Offset(offset int) *chinaCPIDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chinaCPIDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *chinaCPIDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chinaCPIDo) Unscoped() *chinaCPIDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chinaCPIDo) Create(values ...*model.ChinaCPI) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chinaCPIDo) CreateInBatches(values []*model.ChinaCPI, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chinaCPIDo) Save(values ...*model.ChinaCPI) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chinaCPIDo) First() (*model.ChinaCPI, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaCPI), nil
	}
}

func (c chinaCPIDo) Take() (*model.ChinaCPI, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaCPI), nil
	}
}

func (c chinaCPIDo) Last() (*model.ChinaCPI, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaCPI), nil
	}
}

func (c chinaCPIDo) Find() ([]*model.ChinaCPI, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChinaCPI), err
}

func (c chinaCPIDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChinaCPI, err error) {
	buf := make([]*model.ChinaCPI, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chinaCPIDo) FindInBatches(result *[]*model.ChinaCPI, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chinaCPIDo) Attrs(attrs ...field.AssignExpr) *chinaCPIDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chinaCPIDo) Assign(attrs ...field.AssignExpr) *chinaCPIDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chinaCPIDo) Joins(fields ...field.RelationField) *chinaCPIDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chinaCPIDo) Preload(fields ...field.RelationField) *chinaCPIDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chinaCPIDo) FirstOrInit() (*model.ChinaCPI, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaCPI), nil
	}
}

func (c chinaCPIDo) FirstOrCreate() (*model.ChinaCPI, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChinaCPI), nil
	}
}

func (c chinaCPIDo) FindByPage(offset int, limit int) (result []*model.ChinaCPI, count int64, err error) {
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

func (c chinaCPIDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chinaCPIDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chinaCPIDo) Delete(models ...*model.ChinaCPI) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chinaCPIDo) withDO(do gen.Dao) *chinaCPIDo {
	c.DO = *do.(*gen.DO)
	return c
}
