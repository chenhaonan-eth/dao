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

func newSocialFinancingFlow(db *gorm.DB, opts ...gen.DOOption) socialFinancingFlow {
	_socialFinancingFlow := socialFinancingFlow{}

	_socialFinancingFlow.socialFinancingFlowDo.UseDB(db, opts...)
	_socialFinancingFlow.socialFinancingFlowDo.UseModel(&model.SocialFinancingFlow{})

	tableName := _socialFinancingFlow.socialFinancingFlowDo.TableName()
	_socialFinancingFlow.ALL = field.NewAsterisk(tableName)
	_socialFinancingFlow.Date = field.NewString(tableName, "date")
	_socialFinancingFlow.Ndbab = field.NewFloat32(tableName, "ndbab")
	_socialFinancingFlow.Entrustloan = field.NewFloat32(tableName, "entrustloan")
	_socialFinancingFlow.Forcloan = field.NewFloat32(tableName, "forcloan")
	_socialFinancingFlow.Rmblaon = field.NewFloat32(tableName, "rmblaon")
	_socialFinancingFlow.Bibae = field.NewFloat32(tableName, "bibae")
	_socialFinancingFlow.Tiosfs = field.NewFloat32(tableName, "tiosfs")
	_socialFinancingFlow.Sfinfe = field.NewFloat32(tableName, "sfinfe")
	_socialFinancingFlow.Trustloan = field.NewFloat32(tableName, "trustloan")

	_socialFinancingFlow.fillFieldMap()

	return _socialFinancingFlow
}

type socialFinancingFlow struct {
	socialFinancingFlowDo socialFinancingFlowDo

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

func (s socialFinancingFlow) Table(newTableName string) *socialFinancingFlow {
	s.socialFinancingFlowDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s socialFinancingFlow) As(alias string) *socialFinancingFlow {
	s.socialFinancingFlowDo.DO = *(s.socialFinancingFlowDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *socialFinancingFlow) updateTableName(table string) *socialFinancingFlow {
	s.ALL = field.NewAsterisk(table)
	s.Date = field.NewString(table, "date")
	s.Ndbab = field.NewFloat32(table, "ndbab")
	s.Entrustloan = field.NewFloat32(table, "entrustloan")
	s.Forcloan = field.NewFloat32(table, "forcloan")
	s.Rmblaon = field.NewFloat32(table, "rmblaon")
	s.Bibae = field.NewFloat32(table, "bibae")
	s.Tiosfs = field.NewFloat32(table, "tiosfs")
	s.Sfinfe = field.NewFloat32(table, "sfinfe")
	s.Trustloan = field.NewFloat32(table, "trustloan")

	s.fillFieldMap()

	return s
}

func (s *socialFinancingFlow) WithContext(ctx context.Context) *socialFinancingFlowDo {
	return s.socialFinancingFlowDo.WithContext(ctx)
}

func (s socialFinancingFlow) TableName() string { return s.socialFinancingFlowDo.TableName() }

func (s socialFinancingFlow) Alias() string { return s.socialFinancingFlowDo.Alias() }

func (s *socialFinancingFlow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *socialFinancingFlow) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["date"] = s.Date
	s.fieldMap["ndbab"] = s.Ndbab
	s.fieldMap["entrustloan"] = s.Entrustloan
	s.fieldMap["forcloan"] = s.Forcloan
	s.fieldMap["rmblaon"] = s.Rmblaon
	s.fieldMap["bibae"] = s.Bibae
	s.fieldMap["tiosfs"] = s.Tiosfs
	s.fieldMap["sfinfe"] = s.Sfinfe
	s.fieldMap["trustloan"] = s.Trustloan
}

func (s socialFinancingFlow) clone(db *gorm.DB) socialFinancingFlow {
	s.socialFinancingFlowDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s socialFinancingFlow) replaceDB(db *gorm.DB) socialFinancingFlow {
	s.socialFinancingFlowDo.ReplaceDB(db)
	return s
}

type socialFinancingFlowDo struct{ gen.DO }

func (s socialFinancingFlowDo) Debug() *socialFinancingFlowDo {
	return s.withDO(s.DO.Debug())
}

func (s socialFinancingFlowDo) WithContext(ctx context.Context) *socialFinancingFlowDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s socialFinancingFlowDo) ReadDB() *socialFinancingFlowDo {
	return s.Clauses(dbresolver.Read)
}

func (s socialFinancingFlowDo) WriteDB() *socialFinancingFlowDo {
	return s.Clauses(dbresolver.Write)
}

func (s socialFinancingFlowDo) Session(config *gorm.Session) *socialFinancingFlowDo {
	return s.withDO(s.DO.Session(config))
}

func (s socialFinancingFlowDo) Clauses(conds ...clause.Expression) *socialFinancingFlowDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s socialFinancingFlowDo) Returning(value interface{}, columns ...string) *socialFinancingFlowDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s socialFinancingFlowDo) Not(conds ...gen.Condition) *socialFinancingFlowDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s socialFinancingFlowDo) Or(conds ...gen.Condition) *socialFinancingFlowDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s socialFinancingFlowDo) Select(conds ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s socialFinancingFlowDo) Where(conds ...gen.Condition) *socialFinancingFlowDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s socialFinancingFlowDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *socialFinancingFlowDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s socialFinancingFlowDo) Order(conds ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s socialFinancingFlowDo) Distinct(cols ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s socialFinancingFlowDo) Omit(cols ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s socialFinancingFlowDo) Join(table schema.Tabler, on ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s socialFinancingFlowDo) LeftJoin(table schema.Tabler, on ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s socialFinancingFlowDo) RightJoin(table schema.Tabler, on ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s socialFinancingFlowDo) Group(cols ...field.Expr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s socialFinancingFlowDo) Having(conds ...gen.Condition) *socialFinancingFlowDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s socialFinancingFlowDo) Limit(limit int) *socialFinancingFlowDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s socialFinancingFlowDo) Offset(offset int) *socialFinancingFlowDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s socialFinancingFlowDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *socialFinancingFlowDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s socialFinancingFlowDo) Unscoped() *socialFinancingFlowDo {
	return s.withDO(s.DO.Unscoped())
}

func (s socialFinancingFlowDo) Create(values ...*model.SocialFinancingFlow) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s socialFinancingFlowDo) CreateInBatches(values []*model.SocialFinancingFlow, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s socialFinancingFlowDo) Save(values ...*model.SocialFinancingFlow) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s socialFinancingFlowDo) First() (*model.SocialFinancingFlow, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SocialFinancingFlow), nil
	}
}

func (s socialFinancingFlowDo) Take() (*model.SocialFinancingFlow, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SocialFinancingFlow), nil
	}
}

func (s socialFinancingFlowDo) Last() (*model.SocialFinancingFlow, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SocialFinancingFlow), nil
	}
}

func (s socialFinancingFlowDo) Find() ([]*model.SocialFinancingFlow, error) {
	result, err := s.DO.Find()
	return result.([]*model.SocialFinancingFlow), err
}

func (s socialFinancingFlowDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SocialFinancingFlow, err error) {
	buf := make([]*model.SocialFinancingFlow, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s socialFinancingFlowDo) FindInBatches(result *[]*model.SocialFinancingFlow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s socialFinancingFlowDo) Attrs(attrs ...field.AssignExpr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s socialFinancingFlowDo) Assign(attrs ...field.AssignExpr) *socialFinancingFlowDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s socialFinancingFlowDo) Joins(fields ...field.RelationField) *socialFinancingFlowDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s socialFinancingFlowDo) Preload(fields ...field.RelationField) *socialFinancingFlowDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s socialFinancingFlowDo) FirstOrInit() (*model.SocialFinancingFlow, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SocialFinancingFlow), nil
	}
}

func (s socialFinancingFlowDo) FirstOrCreate() (*model.SocialFinancingFlow, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SocialFinancingFlow), nil
	}
}

func (s socialFinancingFlowDo) FindByPage(offset int, limit int) (result []*model.SocialFinancingFlow, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s socialFinancingFlowDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s socialFinancingFlowDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s socialFinancingFlowDo) Delete(models ...*model.SocialFinancingFlow) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *socialFinancingFlowDo) withDO(do gen.Dao) *socialFinancingFlowDo {
	s.DO = *do.(*gen.DO)
	return s
}