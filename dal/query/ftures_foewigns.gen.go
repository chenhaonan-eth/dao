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

func newFturesFoewign(db *gorm.DB, opts ...gen.DOOption) fturesFoewign {
	_fturesFoewign := fturesFoewign{}

	_fturesFoewign.fturesFoewignDo.UseDB(db, opts...)
	_fturesFoewign.fturesFoewignDo.UseModel(&model.FturesFoewign{})

	tableName := _fturesFoewign.fturesFoewignDo.TableName()
	_fturesFoewign.ALL = field.NewAsterisk(tableName)
	_fturesFoewign.Date = field.NewString(tableName, "date")
	_fturesFoewign.Open = field.NewString(tableName, "open")
	_fturesFoewign.High = field.NewString(tableName, "high")
	_fturesFoewign.Low = field.NewString(tableName, "low")
	_fturesFoewign.Close = field.NewString(tableName, "close")
	_fturesFoewign.Volume = field.NewString(tableName, "volume")
	_fturesFoewign.Symbol = field.NewString(tableName, "symbol")

	_fturesFoewign.fillFieldMap()

	return _fturesFoewign
}

type fturesFoewign struct {
	fturesFoewignDo fturesFoewignDo

	ALL    field.Asterisk
	Date   field.String
	Open   field.String
	High   field.String
	Low    field.String
	Close  field.String
	Volume field.String
	Symbol field.String

	fieldMap map[string]field.Expr
}

func (f fturesFoewign) Table(newTableName string) *fturesFoewign {
	f.fturesFoewignDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fturesFoewign) As(alias string) *fturesFoewign {
	f.fturesFoewignDo.DO = *(f.fturesFoewignDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fturesFoewign) updateTableName(table string) *fturesFoewign {
	f.ALL = field.NewAsterisk(table)
	f.Date = field.NewString(table, "date")
	f.Open = field.NewString(table, "open")
	f.High = field.NewString(table, "high")
	f.Low = field.NewString(table, "low")
	f.Close = field.NewString(table, "close")
	f.Volume = field.NewString(table, "volume")
	f.Symbol = field.NewString(table, "symbol")

	f.fillFieldMap()

	return f
}

func (f *fturesFoewign) WithContext(ctx context.Context) *fturesFoewignDo {
	return f.fturesFoewignDo.WithContext(ctx)
}

func (f fturesFoewign) TableName() string { return f.fturesFoewignDo.TableName() }

func (f fturesFoewign) Alias() string { return f.fturesFoewignDo.Alias() }

func (f *fturesFoewign) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fturesFoewign) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["date"] = f.Date
	f.fieldMap["open"] = f.Open
	f.fieldMap["high"] = f.High
	f.fieldMap["low"] = f.Low
	f.fieldMap["close"] = f.Close
	f.fieldMap["volume"] = f.Volume
	f.fieldMap["symbol"] = f.Symbol
}

func (f fturesFoewign) clone(db *gorm.DB) fturesFoewign {
	f.fturesFoewignDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fturesFoewign) replaceDB(db *gorm.DB) fturesFoewign {
	f.fturesFoewignDo.ReplaceDB(db)
	return f
}

type fturesFoewignDo struct{ gen.DO }

func (f fturesFoewignDo) Debug() *fturesFoewignDo {
	return f.withDO(f.DO.Debug())
}

func (f fturesFoewignDo) WithContext(ctx context.Context) *fturesFoewignDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fturesFoewignDo) ReadDB() *fturesFoewignDo {
	return f.Clauses(dbresolver.Read)
}

func (f fturesFoewignDo) WriteDB() *fturesFoewignDo {
	return f.Clauses(dbresolver.Write)
}

func (f fturesFoewignDo) Session(config *gorm.Session) *fturesFoewignDo {
	return f.withDO(f.DO.Session(config))
}

func (f fturesFoewignDo) Clauses(conds ...clause.Expression) *fturesFoewignDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fturesFoewignDo) Returning(value interface{}, columns ...string) *fturesFoewignDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fturesFoewignDo) Not(conds ...gen.Condition) *fturesFoewignDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fturesFoewignDo) Or(conds ...gen.Condition) *fturesFoewignDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fturesFoewignDo) Select(conds ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fturesFoewignDo) Where(conds ...gen.Condition) *fturesFoewignDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fturesFoewignDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *fturesFoewignDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f fturesFoewignDo) Order(conds ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fturesFoewignDo) Distinct(cols ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fturesFoewignDo) Omit(cols ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fturesFoewignDo) Join(table schema.Tabler, on ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fturesFoewignDo) LeftJoin(table schema.Tabler, on ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fturesFoewignDo) RightJoin(table schema.Tabler, on ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fturesFoewignDo) Group(cols ...field.Expr) *fturesFoewignDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fturesFoewignDo) Having(conds ...gen.Condition) *fturesFoewignDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fturesFoewignDo) Limit(limit int) *fturesFoewignDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fturesFoewignDo) Offset(offset int) *fturesFoewignDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fturesFoewignDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *fturesFoewignDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fturesFoewignDo) Unscoped() *fturesFoewignDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fturesFoewignDo) Create(values ...*model.FturesFoewign) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fturesFoewignDo) CreateInBatches(values []*model.FturesFoewign, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fturesFoewignDo) Save(values ...*model.FturesFoewign) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fturesFoewignDo) First() (*model.FturesFoewign, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FturesFoewign), nil
	}
}

func (f fturesFoewignDo) Take() (*model.FturesFoewign, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FturesFoewign), nil
	}
}

func (f fturesFoewignDo) Last() (*model.FturesFoewign, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FturesFoewign), nil
	}
}

func (f fturesFoewignDo) Find() ([]*model.FturesFoewign, error) {
	result, err := f.DO.Find()
	return result.([]*model.FturesFoewign), err
}

func (f fturesFoewignDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FturesFoewign, err error) {
	buf := make([]*model.FturesFoewign, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fturesFoewignDo) FindInBatches(result *[]*model.FturesFoewign, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fturesFoewignDo) Attrs(attrs ...field.AssignExpr) *fturesFoewignDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fturesFoewignDo) Assign(attrs ...field.AssignExpr) *fturesFoewignDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fturesFoewignDo) Joins(fields ...field.RelationField) *fturesFoewignDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fturesFoewignDo) Preload(fields ...field.RelationField) *fturesFoewignDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fturesFoewignDo) FirstOrInit() (*model.FturesFoewign, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FturesFoewign), nil
	}
}

func (f fturesFoewignDo) FirstOrCreate() (*model.FturesFoewign, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FturesFoewign), nil
	}
}

func (f fturesFoewignDo) FindByPage(offset int, limit int) (result []*model.FturesFoewign, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fturesFoewignDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fturesFoewignDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fturesFoewignDo) Delete(models ...*model.FturesFoewign) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fturesFoewignDo) withDO(do gen.Dao) *fturesFoewignDo {
	f.DO = *do.(*gen.DO)
	return f
}
