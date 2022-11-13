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

func newMacroPpi(db *gorm.DB, opts ...gen.DOOption) macroPpi {
	_macroPpi := macroPpi{}

	_macroPpi.macroPpiDo.UseDB(db, opts...)
	_macroPpi.macroPpiDo.UseModel(&model.MacroPpi{})

	tableName := _macroPpi.macroPpiDo.TableName()
	_macroPpi.ALL = field.NewAsterisk(tableName)
	_macroPpi.Date = field.NewString(tableName, "date")
	_macroPpi.Country = field.NewString(tableName, "country")
	_macroPpi.Ppi = field.NewString(tableName, "ppi")

	_macroPpi.fillFieldMap()

	return _macroPpi
}

type macroPpi struct {
	macroPpiDo macroPpiDo

	ALL     field.Asterisk
	Date    field.String
	Country field.String
	Ppi     field.String

	fieldMap map[string]field.Expr
}

func (m macroPpi) Table(newTableName string) *macroPpi {
	m.macroPpiDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m macroPpi) As(alias string) *macroPpi {
	m.macroPpiDo.DO = *(m.macroPpiDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *macroPpi) updateTableName(table string) *macroPpi {
	m.ALL = field.NewAsterisk(table)
	m.Date = field.NewString(table, "date")
	m.Country = field.NewString(table, "country")
	m.Ppi = field.NewString(table, "ppi")

	m.fillFieldMap()

	return m
}

func (m *macroPpi) WithContext(ctx context.Context) *macroPpiDo { return m.macroPpiDo.WithContext(ctx) }

func (m macroPpi) TableName() string { return m.macroPpiDo.TableName() }

func (m macroPpi) Alias() string { return m.macroPpiDo.Alias() }

func (m *macroPpi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *macroPpi) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 3)
	m.fieldMap["date"] = m.Date
	m.fieldMap["country"] = m.Country
	m.fieldMap["ppi"] = m.Ppi
}

func (m macroPpi) clone(db *gorm.DB) macroPpi {
	m.macroPpiDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m macroPpi) replaceDB(db *gorm.DB) macroPpi {
	m.macroPpiDo.ReplaceDB(db)
	return m
}

type macroPpiDo struct{ gen.DO }

func (m macroPpiDo) Debug() *macroPpiDo {
	return m.withDO(m.DO.Debug())
}

func (m macroPpiDo) WithContext(ctx context.Context) *macroPpiDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m macroPpiDo) ReadDB() *macroPpiDo {
	return m.Clauses(dbresolver.Read)
}

func (m macroPpiDo) WriteDB() *macroPpiDo {
	return m.Clauses(dbresolver.Write)
}

func (m macroPpiDo) Session(config *gorm.Session) *macroPpiDo {
	return m.withDO(m.DO.Session(config))
}

func (m macroPpiDo) Clauses(conds ...clause.Expression) *macroPpiDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m macroPpiDo) Returning(value interface{}, columns ...string) *macroPpiDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m macroPpiDo) Not(conds ...gen.Condition) *macroPpiDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m macroPpiDo) Or(conds ...gen.Condition) *macroPpiDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m macroPpiDo) Select(conds ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m macroPpiDo) Where(conds ...gen.Condition) *macroPpiDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m macroPpiDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *macroPpiDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m macroPpiDo) Order(conds ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m macroPpiDo) Distinct(cols ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m macroPpiDo) Omit(cols ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m macroPpiDo) Join(table schema.Tabler, on ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m macroPpiDo) LeftJoin(table schema.Tabler, on ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m macroPpiDo) RightJoin(table schema.Tabler, on ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m macroPpiDo) Group(cols ...field.Expr) *macroPpiDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m macroPpiDo) Having(conds ...gen.Condition) *macroPpiDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m macroPpiDo) Limit(limit int) *macroPpiDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m macroPpiDo) Offset(offset int) *macroPpiDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m macroPpiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *macroPpiDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m macroPpiDo) Unscoped() *macroPpiDo {
	return m.withDO(m.DO.Unscoped())
}

func (m macroPpiDo) Create(values ...*model.MacroPpi) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m macroPpiDo) CreateInBatches(values []*model.MacroPpi, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m macroPpiDo) Save(values ...*model.MacroPpi) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m macroPpiDo) First() (*model.MacroPpi, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPpi), nil
	}
}

func (m macroPpiDo) Take() (*model.MacroPpi, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPpi), nil
	}
}

func (m macroPpiDo) Last() (*model.MacroPpi, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPpi), nil
	}
}

func (m macroPpiDo) Find() ([]*model.MacroPpi, error) {
	result, err := m.DO.Find()
	return result.([]*model.MacroPpi), err
}

func (m macroPpiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MacroPpi, err error) {
	buf := make([]*model.MacroPpi, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m macroPpiDo) FindInBatches(result *[]*model.MacroPpi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m macroPpiDo) Attrs(attrs ...field.AssignExpr) *macroPpiDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m macroPpiDo) Assign(attrs ...field.AssignExpr) *macroPpiDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m macroPpiDo) Joins(fields ...field.RelationField) *macroPpiDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m macroPpiDo) Preload(fields ...field.RelationField) *macroPpiDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m macroPpiDo) FirstOrInit() (*model.MacroPpi, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPpi), nil
	}
}

func (m macroPpiDo) FirstOrCreate() (*model.MacroPpi, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPpi), nil
	}
}

func (m macroPpiDo) FindByPage(offset int, limit int) (result []*model.MacroPpi, count int64, err error) {
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

func (m macroPpiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m macroPpiDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m macroPpiDo) Delete(models ...*model.MacroPpi) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *macroPpiDo) withDO(do gen.Dao) *macroPpiDo {
	m.DO = *do.(*gen.DO)
	return m
}