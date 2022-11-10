// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/chenhaonan-eth/guide-sisyphean/dal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPePbPsDvTotalmv(db *gorm.DB, opts ...gen.DOOption) pePbPsDvTotalmv {
	_pePbPsDvTotalmv := pePbPsDvTotalmv{}
	
	db.AutoMigrate(model.PePbPsDvTotalmv{})
	_pePbPsDvTotalmv.pePbPsDvTotalmvDo.UseDB(db, opts...)
	_pePbPsDvTotalmv.pePbPsDvTotalmvDo.UseModel(&model.PePbPsDvTotalmv{})

	tableName := _pePbPsDvTotalmv.pePbPsDvTotalmvDo.TableName()
	_pePbPsDvTotalmv.ALL = field.NewAsterisk(tableName)
	_pePbPsDvTotalmv.Data = field.NewString(tableName, "data")
	_pePbPsDvTotalmv.Pe = field.NewString(tableName, "pe")
	_pePbPsDvTotalmv.PeTtm = field.NewString(tableName, "pe_ttm")
	_pePbPsDvTotalmv.Pb = field.NewString(tableName, "pb")
	_pePbPsDvTotalmv.Ps = field.NewString(tableName, "ps")
	_pePbPsDvTotalmv.PsTtm = field.NewString(tableName, "ps_ttm")
	_pePbPsDvTotalmv.Dv_ratio = field.NewString(tableName, "dv_ratio")
	_pePbPsDvTotalmv.Dv_ttm = field.NewString(tableName, "dv_ttm")
	_pePbPsDvTotalmv.Total_mv = field.NewString(tableName, "total_mv")
	_pePbPsDvTotalmv.Name = field.NewString(tableName, "name")
	_pePbPsDvTotalmv.Symbol = field.NewString(tableName, "symbol")

	_pePbPsDvTotalmv.fillFieldMap()

	return _pePbPsDvTotalmv
}

type pePbPsDvTotalmv struct {
	pePbPsDvTotalmvDo pePbPsDvTotalmvDo

	ALL      field.Asterisk
	Data     field.String
	Pe       field.String
	PeTtm    field.String
	Pb       field.String
	Ps       field.String
	PsTtm    field.String
	Dv_ratio field.String
	Dv_ttm   field.String
	Total_mv field.String
	Name     field.String
	Symbol   field.String

	fieldMap map[string]field.Expr
}

func (p pePbPsDvTotalmv) Table(newTableName string) *pePbPsDvTotalmv {
	p.pePbPsDvTotalmvDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pePbPsDvTotalmv) As(alias string) *pePbPsDvTotalmv {
	p.pePbPsDvTotalmvDo.DO = *(p.pePbPsDvTotalmvDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pePbPsDvTotalmv) updateTableName(table string) *pePbPsDvTotalmv {
	p.ALL = field.NewAsterisk(table)
	p.Data = field.NewString(table, "data")
	p.Pe = field.NewString(table, "pe")
	p.PeTtm = field.NewString(table, "pe_ttm")
	p.Pb = field.NewString(table, "pb")
	p.Ps = field.NewString(table, "ps")
	p.PsTtm = field.NewString(table, "ps_ttm")
	p.Dv_ratio = field.NewString(table, "dv_ratio")
	p.Dv_ttm = field.NewString(table, "dv_ttm")
	p.Total_mv = field.NewString(table, "total_mv")
	p.Name = field.NewString(table, "name")
	p.Symbol = field.NewString(table, "symbol")

	p.fillFieldMap()

	return p
}

func (p *pePbPsDvTotalmv) WithContext(ctx context.Context) *pePbPsDvTotalmvDo {
	return p.pePbPsDvTotalmvDo.WithContext(ctx)
}

func (p pePbPsDvTotalmv) TableName() string { return p.pePbPsDvTotalmvDo.TableName() }

func (p pePbPsDvTotalmv) Alias() string { return p.pePbPsDvTotalmvDo.Alias() }

func (p *pePbPsDvTotalmv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pePbPsDvTotalmv) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["data"] = p.Data
	p.fieldMap["pe"] = p.Pe
	p.fieldMap["pe_ttm"] = p.PeTtm
	p.fieldMap["pb"] = p.Pb
	p.fieldMap["ps"] = p.Ps
	p.fieldMap["ps_ttm"] = p.PsTtm
	p.fieldMap["dv_ratio"] = p.Dv_ratio
	p.fieldMap["dv_ttm"] = p.Dv_ttm
	p.fieldMap["total_mv"] = p.Total_mv
	p.fieldMap["name"] = p.Name
	p.fieldMap["symbol"] = p.Symbol
}

func (p pePbPsDvTotalmv) clone(db *gorm.DB) pePbPsDvTotalmv {
	p.pePbPsDvTotalmvDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pePbPsDvTotalmv) replaceDB(db *gorm.DB) pePbPsDvTotalmv {
	p.pePbPsDvTotalmvDo.ReplaceDB(db)
	return p
}

type pePbPsDvTotalmvDo struct{ gen.DO }

func (p pePbPsDvTotalmvDo) Debug() *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Debug())
}

func (p pePbPsDvTotalmvDo) WithContext(ctx context.Context) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pePbPsDvTotalmvDo) ReadDB() *pePbPsDvTotalmvDo {
	return p.Clauses(dbresolver.Read)
}

func (p pePbPsDvTotalmvDo) WriteDB() *pePbPsDvTotalmvDo {
	return p.Clauses(dbresolver.Write)
}

func (p pePbPsDvTotalmvDo) Session(config *gorm.Session) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Session(config))
}

func (p pePbPsDvTotalmvDo) Clauses(conds ...clause.Expression) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pePbPsDvTotalmvDo) Returning(value interface{}, columns ...string) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pePbPsDvTotalmvDo) Not(conds ...gen.Condition) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pePbPsDvTotalmvDo) Or(conds ...gen.Condition) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pePbPsDvTotalmvDo) Select(conds ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pePbPsDvTotalmvDo) Where(conds ...gen.Condition) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pePbPsDvTotalmvDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pePbPsDvTotalmvDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pePbPsDvTotalmvDo) Order(conds ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pePbPsDvTotalmvDo) Distinct(cols ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pePbPsDvTotalmvDo) Omit(cols ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pePbPsDvTotalmvDo) Join(table schema.Tabler, on ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pePbPsDvTotalmvDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pePbPsDvTotalmvDo) RightJoin(table schema.Tabler, on ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pePbPsDvTotalmvDo) Group(cols ...field.Expr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pePbPsDvTotalmvDo) Having(conds ...gen.Condition) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pePbPsDvTotalmvDo) Limit(limit int) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pePbPsDvTotalmvDo) Offset(offset int) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pePbPsDvTotalmvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pePbPsDvTotalmvDo) Unscoped() *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pePbPsDvTotalmvDo) Create(values ...*model.PePbPsDvTotalmv) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pePbPsDvTotalmvDo) CreateInBatches(values []*model.PePbPsDvTotalmv, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pePbPsDvTotalmvDo) Save(values ...*model.PePbPsDvTotalmv) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pePbPsDvTotalmvDo) First() (*model.PePbPsDvTotalmv, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PePbPsDvTotalmv), nil
	}
}

func (p pePbPsDvTotalmvDo) Take() (*model.PePbPsDvTotalmv, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PePbPsDvTotalmv), nil
	}
}

func (p pePbPsDvTotalmvDo) Last() (*model.PePbPsDvTotalmv, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PePbPsDvTotalmv), nil
	}
}

func (p pePbPsDvTotalmvDo) Find() ([]*model.PePbPsDvTotalmv, error) {
	result, err := p.DO.Find()
	return result.([]*model.PePbPsDvTotalmv), err
}

func (p pePbPsDvTotalmvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PePbPsDvTotalmv, err error) {
	buf := make([]*model.PePbPsDvTotalmv, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pePbPsDvTotalmvDo) FindInBatches(result *[]*model.PePbPsDvTotalmv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pePbPsDvTotalmvDo) Attrs(attrs ...field.AssignExpr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pePbPsDvTotalmvDo) Assign(attrs ...field.AssignExpr) *pePbPsDvTotalmvDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pePbPsDvTotalmvDo) Joins(fields ...field.RelationField) *pePbPsDvTotalmvDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pePbPsDvTotalmvDo) Preload(fields ...field.RelationField) *pePbPsDvTotalmvDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pePbPsDvTotalmvDo) FirstOrInit() (*model.PePbPsDvTotalmv, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PePbPsDvTotalmv), nil
	}
}

func (p pePbPsDvTotalmvDo) FirstOrCreate() (*model.PePbPsDvTotalmv, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PePbPsDvTotalmv), nil
	}
}

func (p pePbPsDvTotalmvDo) FindByPage(offset int, limit int) (result []*model.PePbPsDvTotalmv, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pePbPsDvTotalmvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pePbPsDvTotalmvDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pePbPsDvTotalmvDo) Delete(models ...*model.PePbPsDvTotalmv) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pePbPsDvTotalmvDo) withDO(do gen.Dao) *pePbPsDvTotalmvDo {
	p.DO = *do.(*gen.DO)
	return p
}
