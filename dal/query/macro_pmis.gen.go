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

func newMacroPMI(db *gorm.DB, opts ...gen.DOOption) macroPMI {
	_macroPMI := macroPMI{}

	_macroPMI.macroPMIDo.UseDB(db, opts...)
	_macroPMI.macroPMIDo.UseModel(&model.MacroPMI{})

	tableName := _macroPMI.macroPMIDo.TableName()
	_macroPMI.ALL = field.NewAsterisk(tableName)
	_macroPMI.Date = field.NewString(tableName, "date")
	_macroPMI.Country = field.NewString(tableName, "country")
	_macroPMI.Pmi = field.NewString(tableName, "pmi")

	_macroPMI.fillFieldMap()

	return _macroPMI
}

type macroPMI struct {
	macroPMIDo macroPMIDo

	ALL     field.Asterisk
	Date    field.String
	Country field.String
	Pmi     field.String

	fieldMap map[string]field.Expr
}

func (m macroPMI) Table(newTableName string) *macroPMI {
	m.macroPMIDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m macroPMI) As(alias string) *macroPMI {
	m.macroPMIDo.DO = *(m.macroPMIDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *macroPMI) updateTableName(table string) *macroPMI {
	m.ALL = field.NewAsterisk(table)
	m.Date = field.NewString(table, "date")
	m.Country = field.NewString(table, "country")
	m.Pmi = field.NewString(table, "pmi")

	m.fillFieldMap()

	return m
}

func (m *macroPMI) WithContext(ctx context.Context) *macroPMIDo { return m.macroPMIDo.WithContext(ctx) }

func (m macroPMI) TableName() string { return m.macroPMIDo.TableName() }

func (m macroPMI) Alias() string { return m.macroPMIDo.Alias() }

func (m *macroPMI) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *macroPMI) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 3)
	m.fieldMap["date"] = m.Date
	m.fieldMap["country"] = m.Country
	m.fieldMap["pmi"] = m.Pmi
}

func (m macroPMI) clone(db *gorm.DB) macroPMI {
	m.macroPMIDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m macroPMI) replaceDB(db *gorm.DB) macroPMI {
	m.macroPMIDo.ReplaceDB(db)
	return m
}

type macroPMIDo struct{ gen.DO }

func (m macroPMIDo) Debug() *macroPMIDo {
	return m.withDO(m.DO.Debug())
}

func (m macroPMIDo) WithContext(ctx context.Context) *macroPMIDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m macroPMIDo) ReadDB() *macroPMIDo {
	return m.Clauses(dbresolver.Read)
}

func (m macroPMIDo) WriteDB() *macroPMIDo {
	return m.Clauses(dbresolver.Write)
}

func (m macroPMIDo) Session(config *gorm.Session) *macroPMIDo {
	return m.withDO(m.DO.Session(config))
}

func (m macroPMIDo) Clauses(conds ...clause.Expression) *macroPMIDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m macroPMIDo) Returning(value interface{}, columns ...string) *macroPMIDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m macroPMIDo) Not(conds ...gen.Condition) *macroPMIDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m macroPMIDo) Or(conds ...gen.Condition) *macroPMIDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m macroPMIDo) Select(conds ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m macroPMIDo) Where(conds ...gen.Condition) *macroPMIDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m macroPMIDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *macroPMIDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m macroPMIDo) Order(conds ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m macroPMIDo) Distinct(cols ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m macroPMIDo) Omit(cols ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m macroPMIDo) Join(table schema.Tabler, on ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m macroPMIDo) LeftJoin(table schema.Tabler, on ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m macroPMIDo) RightJoin(table schema.Tabler, on ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m macroPMIDo) Group(cols ...field.Expr) *macroPMIDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m macroPMIDo) Having(conds ...gen.Condition) *macroPMIDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m macroPMIDo) Limit(limit int) *macroPMIDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m macroPMIDo) Offset(offset int) *macroPMIDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m macroPMIDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *macroPMIDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m macroPMIDo) Unscoped() *macroPMIDo {
	return m.withDO(m.DO.Unscoped())
}

func (m macroPMIDo) Create(values ...*model.MacroPMI) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m macroPMIDo) CreateInBatches(values []*model.MacroPMI, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m macroPMIDo) Save(values ...*model.MacroPMI) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m macroPMIDo) First() (*model.MacroPMI, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPMI), nil
	}
}

func (m macroPMIDo) Take() (*model.MacroPMI, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPMI), nil
	}
}

func (m macroPMIDo) Last() (*model.MacroPMI, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPMI), nil
	}
}

func (m macroPMIDo) Find() ([]*model.MacroPMI, error) {
	result, err := m.DO.Find()
	return result.([]*model.MacroPMI), err
}

func (m macroPMIDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MacroPMI, err error) {
	buf := make([]*model.MacroPMI, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m macroPMIDo) FindInBatches(result *[]*model.MacroPMI, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m macroPMIDo) Attrs(attrs ...field.AssignExpr) *macroPMIDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m macroPMIDo) Assign(attrs ...field.AssignExpr) *macroPMIDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m macroPMIDo) Joins(fields ...field.RelationField) *macroPMIDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m macroPMIDo) Preload(fields ...field.RelationField) *macroPMIDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m macroPMIDo) FirstOrInit() (*model.MacroPMI, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPMI), nil
	}
}

func (m macroPMIDo) FirstOrCreate() (*model.MacroPMI, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroPMI), nil
	}
}

func (m macroPMIDo) FindByPage(offset int, limit int) (result []*model.MacroPMI, count int64, err error) {
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

func (m macroPMIDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m macroPMIDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m macroPMIDo) Delete(models ...*model.MacroPMI) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *macroPMIDo) withDO(do gen.Dao) *macroPMIDo {
	m.DO = *do.(*gen.DO)
	return m
}