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

func newMacroChinaShrzgm(db *gorm.DB, opts ...gen.DOOption) macroChinaShrzgm {
	_macroChinaShrzgm := macroChinaShrzgm{}

	_macroChinaShrzgm.macroChinaShrzgmDo.UseDB(db, opts...)
	_macroChinaShrzgm.macroChinaShrzgmDo.UseModel(&model.MacroChinaShrzgm{})

	tableName := _macroChinaShrzgm.macroChinaShrzgmDo.TableName()
	_macroChinaShrzgm.ALL = field.NewAsterisk(tableName)
	_macroChinaShrzgm.Date = field.NewString(tableName, "date")
	_macroChinaShrzgm.Ndbab = field.NewFloat32(tableName, "ndbab")
	_macroChinaShrzgm.Entrustloan = field.NewFloat32(tableName, "entrustloan")
	_macroChinaShrzgm.Forcloan = field.NewFloat32(tableName, "forcloan")
	_macroChinaShrzgm.Rmblaon = field.NewFloat32(tableName, "rmblaon")
	_macroChinaShrzgm.Bibae = field.NewFloat32(tableName, "bibae")
	_macroChinaShrzgm.Tiosfs = field.NewFloat32(tableName, "tiosfs")
	_macroChinaShrzgm.Sfinfe = field.NewFloat32(tableName, "sfinfe")
	_macroChinaShrzgm.Trustloan = field.NewFloat32(tableName, "trustloan")

	_macroChinaShrzgm.fillFieldMap()

	return _macroChinaShrzgm
}

type macroChinaShrzgm struct {
	macroChinaShrzgmDo macroChinaShrzgmDo

	ALL         field.Asterisk
	Date        field.String
	Ndbab       field.Float32
	Entrustloan field.Float32
	Forcloan    field.Float32
	Rmblaon     field.Float32
	Bibae       field.Float32
	Tiosfs      field.Float32
	Sfinfe      field.Float32
	Trustloan   field.Float32

	fieldMap map[string]field.Expr
}

func (m macroChinaShrzgm) Table(newTableName string) *macroChinaShrzgm {
	m.macroChinaShrzgmDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m macroChinaShrzgm) As(alias string) *macroChinaShrzgm {
	m.macroChinaShrzgmDo.DO = *(m.macroChinaShrzgmDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *macroChinaShrzgm) updateTableName(table string) *macroChinaShrzgm {
	m.ALL = field.NewAsterisk(table)
	m.Date = field.NewString(table, "date")
	m.Ndbab = field.NewFloat32(table, "ndbab")
	m.Entrustloan = field.NewFloat32(table, "entrustloan")
	m.Forcloan = field.NewFloat32(table, "forcloan")
	m.Rmblaon = field.NewFloat32(table, "rmblaon")
	m.Bibae = field.NewFloat32(table, "bibae")
	m.Tiosfs = field.NewFloat32(table, "tiosfs")
	m.Sfinfe = field.NewFloat32(table, "sfinfe")
	m.Trustloan = field.NewFloat32(table, "trustloan")

	m.fillFieldMap()

	return m
}

func (m *macroChinaShrzgm) WithContext(ctx context.Context) *macroChinaShrzgmDo {
	return m.macroChinaShrzgmDo.WithContext(ctx)
}

func (m macroChinaShrzgm) TableName() string { return m.macroChinaShrzgmDo.TableName() }

func (m macroChinaShrzgm) Alias() string { return m.macroChinaShrzgmDo.Alias() }

func (m *macroChinaShrzgm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *macroChinaShrzgm) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["date"] = m.Date
	m.fieldMap["ndbab"] = m.Ndbab
	m.fieldMap["entrustloan"] = m.Entrustloan
	m.fieldMap["forcloan"] = m.Forcloan
	m.fieldMap["rmblaon"] = m.Rmblaon
	m.fieldMap["bibae"] = m.Bibae
	m.fieldMap["tiosfs"] = m.Tiosfs
	m.fieldMap["sfinfe"] = m.Sfinfe
	m.fieldMap["trustloan"] = m.Trustloan
}

func (m macroChinaShrzgm) clone(db *gorm.DB) macroChinaShrzgm {
	m.macroChinaShrzgmDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m macroChinaShrzgm) replaceDB(db *gorm.DB) macroChinaShrzgm {
	m.macroChinaShrzgmDo.ReplaceDB(db)
	return m
}

type macroChinaShrzgmDo struct{ gen.DO }

func (m macroChinaShrzgmDo) Debug() *macroChinaShrzgmDo {
	return m.withDO(m.DO.Debug())
}

func (m macroChinaShrzgmDo) WithContext(ctx context.Context) *macroChinaShrzgmDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m macroChinaShrzgmDo) ReadDB() *macroChinaShrzgmDo {
	return m.Clauses(dbresolver.Read)
}

func (m macroChinaShrzgmDo) WriteDB() *macroChinaShrzgmDo {
	return m.Clauses(dbresolver.Write)
}

func (m macroChinaShrzgmDo) Session(config *gorm.Session) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Session(config))
}

func (m macroChinaShrzgmDo) Clauses(conds ...clause.Expression) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m macroChinaShrzgmDo) Returning(value interface{}, columns ...string) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m macroChinaShrzgmDo) Not(conds ...gen.Condition) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m macroChinaShrzgmDo) Or(conds ...gen.Condition) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m macroChinaShrzgmDo) Select(conds ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m macroChinaShrzgmDo) Where(conds ...gen.Condition) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m macroChinaShrzgmDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *macroChinaShrzgmDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m macroChinaShrzgmDo) Order(conds ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m macroChinaShrzgmDo) Distinct(cols ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m macroChinaShrzgmDo) Omit(cols ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m macroChinaShrzgmDo) Join(table schema.Tabler, on ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m macroChinaShrzgmDo) LeftJoin(table schema.Tabler, on ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m macroChinaShrzgmDo) RightJoin(table schema.Tabler, on ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m macroChinaShrzgmDo) Group(cols ...field.Expr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m macroChinaShrzgmDo) Having(conds ...gen.Condition) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m macroChinaShrzgmDo) Limit(limit int) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m macroChinaShrzgmDo) Offset(offset int) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m macroChinaShrzgmDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m macroChinaShrzgmDo) Unscoped() *macroChinaShrzgmDo {
	return m.withDO(m.DO.Unscoped())
}

func (m macroChinaShrzgmDo) Create(values ...*model.MacroChinaShrzgm) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m macroChinaShrzgmDo) CreateInBatches(values []*model.MacroChinaShrzgm, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m macroChinaShrzgmDo) Save(values ...*model.MacroChinaShrzgm) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m macroChinaShrzgmDo) First() (*model.MacroChinaShrzgm, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaShrzgm), nil
	}
}

func (m macroChinaShrzgmDo) Take() (*model.MacroChinaShrzgm, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaShrzgm), nil
	}
}

func (m macroChinaShrzgmDo) Last() (*model.MacroChinaShrzgm, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaShrzgm), nil
	}
}

func (m macroChinaShrzgmDo) Find() ([]*model.MacroChinaShrzgm, error) {
	result, err := m.DO.Find()
	return result.([]*model.MacroChinaShrzgm), err
}

func (m macroChinaShrzgmDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MacroChinaShrzgm, err error) {
	buf := make([]*model.MacroChinaShrzgm, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m macroChinaShrzgmDo) FindInBatches(result *[]*model.MacroChinaShrzgm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m macroChinaShrzgmDo) Attrs(attrs ...field.AssignExpr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m macroChinaShrzgmDo) Assign(attrs ...field.AssignExpr) *macroChinaShrzgmDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m macroChinaShrzgmDo) Joins(fields ...field.RelationField) *macroChinaShrzgmDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m macroChinaShrzgmDo) Preload(fields ...field.RelationField) *macroChinaShrzgmDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m macroChinaShrzgmDo) FirstOrInit() (*model.MacroChinaShrzgm, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaShrzgm), nil
	}
}

func (m macroChinaShrzgmDo) FirstOrCreate() (*model.MacroChinaShrzgm, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MacroChinaShrzgm), nil
	}
}

func (m macroChinaShrzgmDo) FindByPage(offset int, limit int) (result []*model.MacroChinaShrzgm, count int64, err error) {
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

func (m macroChinaShrzgmDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m macroChinaShrzgmDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m macroChinaShrzgmDo) Delete(models ...*model.MacroChinaShrzgm) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *macroChinaShrzgmDo) withDO(do gen.Dao) *macroChinaShrzgmDo {
	m.DO = *do.(*gen.DO)
	return m
}