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

func newSH300PEModel(db *gorm.DB, opts ...gen.DOOption) sH300PEModel {
	_sH300PEModel := sH300PEModel{}

	_sH300PEModel.sH300PEModelDo.UseDB(db, opts...)
	_sH300PEModel.sH300PEModelDo.UseModel(&model.SH300PEModel{})

	tableName := _sH300PEModel.sH300PEModelDo.TableName()
	_sH300PEModel.ALL = field.NewAsterisk(tableName)
	_sH300PEModel.Data = field.NewFloat64(tableName, "data")
	_sH300PEModel.MiddleLyrPe = field.NewFloat64(tableName, "middle_lyr_pe")
	_sH300PEModel.LyrPe = field.NewFloat64(tableName, "lyr_pe")
	_sH300PEModel.AddLyrPe = field.NewFloat64(tableName, "add_lyr_pe")
	_sH300PEModel.MiddleTtmPe = field.NewFloat64(tableName, "middle_ttm_pe")
	_sH300PEModel.TtmPe = field.NewFloat64(tableName, "ttm_pe")
	_sH300PEModel.AddTtmPe = field.NewFloat64(tableName, "add_ttm_pe")

	_sH300PEModel.fillFieldMap()

	return _sH300PEModel
}

type sH300PEModel struct {
	sH300PEModelDo sH300PEModelDo

	ALL         field.Asterisk
	Data        field.Float64
	MiddleLyrPe field.Float64
	LyrPe       field.Float64
	AddLyrPe    field.Float64
	MiddleTtmPe field.Float64
	TtmPe       field.Float64
	AddTtmPe    field.Float64

	fieldMap map[string]field.Expr
}

func (s sH300PEModel) Table(newTableName string) *sH300PEModel {
	s.sH300PEModelDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sH300PEModel) As(alias string) *sH300PEModel {
	s.sH300PEModelDo.DO = *(s.sH300PEModelDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sH300PEModel) updateTableName(table string) *sH300PEModel {
	s.ALL = field.NewAsterisk(table)
	s.Data = field.NewFloat64(table, "data")
	s.MiddleLyrPe = field.NewFloat64(table, "middle_lyr_pe")
	s.LyrPe = field.NewFloat64(table, "lyr_pe")
	s.AddLyrPe = field.NewFloat64(table, "add_lyr_pe")
	s.MiddleTtmPe = field.NewFloat64(table, "middle_ttm_pe")
	s.TtmPe = field.NewFloat64(table, "ttm_pe")
	s.AddTtmPe = field.NewFloat64(table, "add_ttm_pe")

	s.fillFieldMap()

	return s
}

func (s *sH300PEModel) WithContext(ctx context.Context) *sH300PEModelDo {
	return s.sH300PEModelDo.WithContext(ctx)
}

func (s sH300PEModel) TableName() string { return s.sH300PEModelDo.TableName() }

func (s sH300PEModel) Alias() string { return s.sH300PEModelDo.Alias() }

func (s *sH300PEModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sH300PEModel) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["data"] = s.Data
	s.fieldMap["middle_lyr_pe"] = s.MiddleLyrPe
	s.fieldMap["lyr_pe"] = s.LyrPe
	s.fieldMap["add_lyr_pe"] = s.AddLyrPe
	s.fieldMap["middle_ttm_pe"] = s.MiddleTtmPe
	s.fieldMap["ttm_pe"] = s.TtmPe
	s.fieldMap["add_ttm_pe"] = s.AddTtmPe
}

func (s sH300PEModel) clone(db *gorm.DB) sH300PEModel {
	s.sH300PEModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sH300PEModel) replaceDB(db *gorm.DB) sH300PEModel {
	s.sH300PEModelDo.ReplaceDB(db)
	return s
}

type sH300PEModelDo struct{ gen.DO }

func (s sH300PEModelDo) Debug() *sH300PEModelDo {
	return s.withDO(s.DO.Debug())
}

func (s sH300PEModelDo) WithContext(ctx context.Context) *sH300PEModelDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sH300PEModelDo) ReadDB() *sH300PEModelDo {
	return s.Clauses(dbresolver.Read)
}

func (s sH300PEModelDo) WriteDB() *sH300PEModelDo {
	return s.Clauses(dbresolver.Write)
}

func (s sH300PEModelDo) Session(config *gorm.Session) *sH300PEModelDo {
	return s.withDO(s.DO.Session(config))
}

func (s sH300PEModelDo) Clauses(conds ...clause.Expression) *sH300PEModelDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sH300PEModelDo) Returning(value interface{}, columns ...string) *sH300PEModelDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sH300PEModelDo) Not(conds ...gen.Condition) *sH300PEModelDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sH300PEModelDo) Or(conds ...gen.Condition) *sH300PEModelDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sH300PEModelDo) Select(conds ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sH300PEModelDo) Where(conds ...gen.Condition) *sH300PEModelDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sH300PEModelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sH300PEModelDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sH300PEModelDo) Order(conds ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sH300PEModelDo) Distinct(cols ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sH300PEModelDo) Omit(cols ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sH300PEModelDo) Join(table schema.Tabler, on ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sH300PEModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sH300PEModelDo) RightJoin(table schema.Tabler, on ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sH300PEModelDo) Group(cols ...field.Expr) *sH300PEModelDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sH300PEModelDo) Having(conds ...gen.Condition) *sH300PEModelDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sH300PEModelDo) Limit(limit int) *sH300PEModelDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sH300PEModelDo) Offset(offset int) *sH300PEModelDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sH300PEModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sH300PEModelDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sH300PEModelDo) Unscoped() *sH300PEModelDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sH300PEModelDo) Create(values ...*model.SH300PEModel) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sH300PEModelDo) CreateInBatches(values []*model.SH300PEModel, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sH300PEModelDo) Save(values ...*model.SH300PEModel) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sH300PEModelDo) First() (*model.SH300PEModel, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SH300PEModel), nil
	}
}

func (s sH300PEModelDo) Take() (*model.SH300PEModel, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SH300PEModel), nil
	}
}

func (s sH300PEModelDo) Last() (*model.SH300PEModel, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SH300PEModel), nil
	}
}

func (s sH300PEModelDo) Find() ([]*model.SH300PEModel, error) {
	result, err := s.DO.Find()
	return result.([]*model.SH300PEModel), err
}

func (s sH300PEModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SH300PEModel, err error) {
	buf := make([]*model.SH300PEModel, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sH300PEModelDo) FindInBatches(result *[]*model.SH300PEModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sH300PEModelDo) Attrs(attrs ...field.AssignExpr) *sH300PEModelDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sH300PEModelDo) Assign(attrs ...field.AssignExpr) *sH300PEModelDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sH300PEModelDo) Joins(fields ...field.RelationField) *sH300PEModelDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sH300PEModelDo) Preload(fields ...field.RelationField) *sH300PEModelDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sH300PEModelDo) FirstOrInit() (*model.SH300PEModel, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SH300PEModel), nil
	}
}

func (s sH300PEModelDo) FirstOrCreate() (*model.SH300PEModel, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SH300PEModel), nil
	}
}

func (s sH300PEModelDo) FindByPage(offset int, limit int) (result []*model.SH300PEModel, count int64, err error) {
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

func (s sH300PEModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sH300PEModelDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sH300PEModelDo) Delete(models ...*model.SH300PEModel) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sH300PEModelDo) withDO(do gen.Dao) *sH300PEModelDo {
	s.DO = *do.(*gen.DO)
	return s
}