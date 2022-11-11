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

func newMacroChinaMoneySupplyModel(db *gorm.DB, opts ...gen.DOOption) macroChinaMoneySupplyModel {
	_macroChinaMoneySupplyModel := macroChinaMoneySupplyModel{}

	_macroChinaMoneySupplyModel.macroChinaMoneySupplyModelDo.UseDB(db, opts...)
	_macroChinaMoneySupplyModel.macroChinaMoneySupplyModelDo.UseModel(&model.MacroChinaMoneySupplyModel{})

	tableName := _macroChinaMoneySupplyModel.macroChinaMoneySupplyModelDo.TableName()
	_macroChinaMoneySupplyModel.ALL = field.NewAsterisk(tableName)
	_macroChinaMoneySupplyModel.Date = field.NewTime(tableName, "date")
	_macroChinaMoneySupplyModel.M2 = field.NewFloat64(tableName, "m2")
	_macroChinaMoneySupplyModel.M2YearOnYear = field.NewFloat64(tableName, "m2_year_on_year")
	_macroChinaMoneySupplyModel.M2YearOverYear = field.NewFloat64(tableName, "m2_year_over_year")
	_macroChinaMoneySupplyModel.M1 = field.NewFloat64(tableName, "m1")
	_macroChinaMoneySupplyModel.M1YearOnYear = field.NewFloat64(tableName, "m1_year_on_year")
	_macroChinaMoneySupplyModel.M1YearOverYear = field.NewFloat64(tableName, "m1_year_over_year")
	_macroChinaMoneySupplyModel.M0 = field.NewFloat64(tableName, "m0")
	_macroChinaMoneySupplyModel.M0YearOnYear = field.NewFloat64(tableName, "m0_year_on_year")
	_macroChinaMoneySupplyModel.M0YearOverYear = field.NewFloat64(tableName, "m0_year_over_year")

	_macroChinaMoneySupplyModel.fillFieldMap()

	return _macroChinaMoneySupplyModel
}

type macroChinaMoneySupplyModel struct {
	macroChinaMoneySupplyModelDo macroChinaMoneySupplyModelDo

	ALL            field.Asterisk
	Date           field.Time
	M2             field.Float64
	M2YearOnYear   field.Float64
	M2YearOverYear field.Float64
	M1             field.Float64
	M1YearOnYear   field.Float64
	M1YearOverYear field.Float64
	M0             field.Float64
	M0YearOnYear   field.Float64
	M0YearOverYear field.Float64

	fieldMap map[string]field.Expr
}

func (m macroChinaMoneySupplyModel) Table(newTableName string) *macroChinaMoneySupplyModel {
	m.macroChinaMoneySupplyModelDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m macroChinaMoneySupplyModel) As(alias string) *macroChinaMoneySupplyModel {
	m.macroChinaMoneySupplyModelDo.DO = *(m.macroChinaMoneySupplyModelDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *macroChinaMoneySupplyModel) updateTableName(table string) *macroChinaMoneySupplyModel {
	m.ALL = field.NewAsterisk(table)
	m.Date = field.NewTime(table, "date")
	m.M2 = field.NewFloat64(table, "m2")
	m.M2YearOnYear = field.NewFloat64(table, "m2_year_on_year")
	m.M2YearOverYear = field.NewFloat64(table, "m2_year_over_year")
	m.M1 = field.NewFloat64(table, "m1")
	m.M1YearOnYear = field.NewFloat64(table, "m1_year_on_year")
	m.M1YearOverYear = field.NewFloat64(table, "m1_year_over_year")
	m.M0 = field.NewFloat64(table, "m0")
	m.M0YearOnYear = field.NewFloat64(table, "m0_year_on_year")
	m.M0YearOverYear = field.NewFloat64(table, "m0_year_over_year")

	m.fillFieldMap()

	return m
}

func (m *macroChinaMoneySupplyModel) WithContext(ctx context.Context) *macroChinaMoneySupplyModelDo {
	return m.macroChinaMoneySupplyModelDo.WithContext(ctx)
}

func (m macroChinaMoneySupplyModel) TableName() string {
	return m.macroChinaMoneySupplyModelDo.TableName()
}

func (m macroChinaMoneySupplyModel) Alias() string { return m.macroChinaMoneySupplyModelDo.Alias() }

func (m *macroChinaMoneySupplyModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *macroChinaMoneySupplyModel) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 10)
	m.fieldMap["date"] = m.Date
	m.fieldMap["m2"] = m.M2
	m.fieldMap["m2_year_on_year"] = m.M2YearOnYear
	m.fieldMap["m2_year_over_year"] = m.M2YearOverYear
	m.fieldMap["m1"] = m.M1
	m.fieldMap["m1_year_on_year"] = m.M1YearOnYear
	m.fieldMap["m1_year_over_year"] = m.M1YearOverYear
	m.fieldMap["m0"] = m.M0
	m.fieldMap["m0_year_on_year"] = m.M0YearOnYear
	m.fieldMap["m0_year_over_year"] = m.M0YearOverYear
}

func (m macroChinaMoneySupplyModel) clone(db *gorm.DB) macroChinaMoneySupplyModel {
	m.macroChinaMoneySupplyModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m macroChinaMoneySupplyModel) replaceDB(db *gorm.DB) macroChinaMoneySupplyModel {
	m.macroChinaMoneySupplyModelDo.ReplaceDB(db)
	return m
}

type macroChinaMoneySupplyModelDo struct{ gen.DO }

func (m macroChinaMoneySupplyModelDo) Debug() *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Debug())
}

func (m macroChinaMoneySupplyModelDo) WithContext(ctx context.Context) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m macroChinaMoneySupplyModelDo) ReadDB() *macroChinaMoneySupplyModelDo {
	return m.Clauses(dbresolver.Read)
}

func (m macroChinaMoneySupplyModelDo) WriteDB() *macroChinaMoneySupplyModelDo {
	return m.Clauses(dbresolver.Write)
}

func (m macroChinaMoneySupplyModelDo) Session(config *gorm.Session) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Session(config))
}

func (m macroChinaMoneySupplyModelDo) Clauses(conds ...clause.Expression) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m macroChinaMoneySupplyModelDo) Returning(value interface{}, columns ...string) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m macroChinaMoneySupplyModelDo) Not(conds ...gen.Condition) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m macroChinaMoneySupplyModelDo) Or(conds ...gen.Condition) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m macroChinaMoneySupplyModelDo) Select(conds ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m macroChinaMoneySupplyModelDo) Where(conds ...gen.Condition) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m macroChinaMoneySupplyModelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *macroChinaMoneySupplyModelDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m macroChinaMoneySupplyModelDo) Order(conds ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m macroChinaMoneySupplyModelDo) Distinct(cols ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m macroChinaMoneySupplyModelDo) Omit(cols ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m macroChinaMoneySupplyModelDo) Join(table schema.Tabler, on ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m macroChinaMoneySupplyModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m macroChinaMoneySupplyModelDo) RightJoin(table schema.Tabler, on ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m macroChinaMoneySupplyModelDo) Group(cols ...field.Expr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m macroChinaMoneySupplyModelDo) Having(conds ...gen.Condition) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m macroChinaMoneySupplyModelDo) Limit(limit int) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m macroChinaMoneySupplyModelDo) Offset(offset int) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m macroChinaMoneySupplyModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m macroChinaMoneySupplyModelDo) Unscoped() *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Unscoped())
}

func (m macroChinaMoneySupplyModelDo) Create(values ...*model.MacroChinaMoneySupplyModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m macroChinaMoneySupplyModelDo) CreateInBatches(values []*model.MacroChinaMoneySupplyModel, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m macroChinaMoneySupplyModelDo) Save(values ...*model.MacroChinaMoneySupplyModel) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m macroChinaMoneySupplyModelDo) First() (*model.MacroChinaMoneySupplyModel, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaMoneySupplyModel), nil
	}
}

func (m macroChinaMoneySupplyModelDo) Take() (*model.MacroChinaMoneySupplyModel, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaMoneySupplyModel), nil
	}
}

func (m macroChinaMoneySupplyModelDo) Last() (*model.MacroChinaMoneySupplyModel, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaMoneySupplyModel), nil
	}
}

func (m macroChinaMoneySupplyModelDo) Find() ([]*model.MacroChinaMoneySupplyModel, error) {
	result, err := m.DO.Find()
	return result.([]*model.MacroChinaMoneySupplyModel), err
}

func (m macroChinaMoneySupplyModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MacroChinaMoneySupplyModel, err error) {
	buf := make([]*model.MacroChinaMoneySupplyModel, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m macroChinaMoneySupplyModelDo) FindInBatches(result *[]*model.MacroChinaMoneySupplyModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m macroChinaMoneySupplyModelDo) Attrs(attrs ...field.AssignExpr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m macroChinaMoneySupplyModelDo) Assign(attrs ...field.AssignExpr) *macroChinaMoneySupplyModelDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m macroChinaMoneySupplyModelDo) Joins(fields ...field.RelationField) *macroChinaMoneySupplyModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m macroChinaMoneySupplyModelDo) Preload(fields ...field.RelationField) *macroChinaMoneySupplyModelDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m macroChinaMoneySupplyModelDo) FirstOrInit() (*model.MacroChinaMoneySupplyModel, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaMoneySupplyModel), nil
	}
}

func (m macroChinaMoneySupplyModelDo) FirstOrCreate() (*model.MacroChinaMoneySupplyModel, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaMoneySupplyModel), nil
	}
}

func (m macroChinaMoneySupplyModelDo) FindByPage(offset int, limit int) (result []*model.MacroChinaMoneySupplyModel, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m macroChinaMoneySupplyModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m macroChinaMoneySupplyModelDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m macroChinaMoneySupplyModelDo) Delete(models ...*model.MacroChinaMoneySupplyModel) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *macroChinaMoneySupplyModelDo) withDO(do gen.Dao) *macroChinaMoneySupplyModelDo {
	m.DO = *do.(*gen.DO)
	return m
}