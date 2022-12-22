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

func newInvestmentInFixedAssets(db *gorm.DB, opts ...gen.DOOption) investmentInFixedAssets {
	_investmentInFixedAssets := investmentInFixedAssets{}

	_investmentInFixedAssets.investmentInFixedAssetsDo.UseDB(db, opts...)
	_investmentInFixedAssets.investmentInFixedAssetsDo.UseModel(&model.InvestmentInFixedAssets{})

	tableName := _investmentInFixedAssets.investmentInFixedAssetsDo.TableName()
	_investmentInFixedAssets.ALL = field.NewAsterisk(tableName)
	_investmentInFixedAssets.Date = field.NewString(tableName, "date")
	_investmentInFixedAssets.InvestmentCompletedAmount = field.NewString(tableName, "investment_completed_amount")
	_investmentInFixedAssets.InvestmentCompletedAmountYearOnYear = field.NewString(tableName, "investment_completed_amount_year_on_year")
	_investmentInFixedAssets.PrimaryIndustry = field.NewString(tableName, "primary_industry")
	_investmentInFixedAssets.PrimaryIndustryYearOnYear = field.NewString(tableName, "primary_industry_year_on_year")
	_investmentInFixedAssets.SecondaryIndustry = field.NewString(tableName, "secondary_industry")
	_investmentInFixedAssets.SecondaryIndustryYearOnYear = field.NewString(tableName, "secondary_industry_year_on_year")
	_investmentInFixedAssets.TertiaryIndustry = field.NewString(tableName, "tertiary_industry")
	_investmentInFixedAssets.TertiaryIndustryYearOnYear = field.NewString(tableName, "tertiary_industry_year_on_year")
	_investmentInFixedAssets.RealEstateDevelopment = field.NewString(tableName, "real_estate_development")
	_investmentInFixedAssets.RealEstateDevelopmentYearOnYear = field.NewString(tableName, "real_estate_development_year_on_year")

	_investmentInFixedAssets.fillFieldMap()

	return _investmentInFixedAssets
}

type investmentInFixedAssets struct {
	investmentInFixedAssetsDo investmentInFixedAssetsDo

	ALL                                 field.Asterisk
	Date                                field.String
	InvestmentCompletedAmount           field.String
	InvestmentCompletedAmountYearOnYear field.String
	PrimaryIndustry                     field.String
	PrimaryIndustryYearOnYear           field.String
	SecondaryIndustry                   field.String
	SecondaryIndustryYearOnYear         field.String
	TertiaryIndustry                    field.String
	TertiaryIndustryYearOnYear          field.String
	RealEstateDevelopment               field.String
	RealEstateDevelopmentYearOnYear     field.String

	fieldMap map[string]field.Expr
}

func (i investmentInFixedAssets) Table(newTableName string) *investmentInFixedAssets {
	i.investmentInFixedAssetsDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i investmentInFixedAssets) As(alias string) *investmentInFixedAssets {
	i.investmentInFixedAssetsDo.DO = *(i.investmentInFixedAssetsDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *investmentInFixedAssets) updateTableName(table string) *investmentInFixedAssets {
	i.ALL = field.NewAsterisk(table)
	i.Date = field.NewString(table, "date")
	i.InvestmentCompletedAmount = field.NewString(table, "investment_completed_amount")
	i.InvestmentCompletedAmountYearOnYear = field.NewString(table, "investment_completed_amount_year_on_year")
	i.PrimaryIndustry = field.NewString(table, "primary_industry")
	i.PrimaryIndustryYearOnYear = field.NewString(table, "primary_industry_year_on_year")
	i.SecondaryIndustry = field.NewString(table, "secondary_industry")
	i.SecondaryIndustryYearOnYear = field.NewString(table, "secondary_industry_year_on_year")
	i.TertiaryIndustry = field.NewString(table, "tertiary_industry")
	i.TertiaryIndustryYearOnYear = field.NewString(table, "tertiary_industry_year_on_year")
	i.RealEstateDevelopment = field.NewString(table, "real_estate_development")
	i.RealEstateDevelopmentYearOnYear = field.NewString(table, "real_estate_development_year_on_year")

	i.fillFieldMap()

	return i
}

func (i *investmentInFixedAssets) WithContext(ctx context.Context) *investmentInFixedAssetsDo {
	return i.investmentInFixedAssetsDo.WithContext(ctx)
}

func (i investmentInFixedAssets) TableName() string { return i.investmentInFixedAssetsDo.TableName() }

func (i investmentInFixedAssets) Alias() string { return i.investmentInFixedAssetsDo.Alias() }

func (i *investmentInFixedAssets) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *investmentInFixedAssets) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 11)
	i.fieldMap["date"] = i.Date
	i.fieldMap["investment_completed_amount"] = i.InvestmentCompletedAmount
	i.fieldMap["investment_completed_amount_year_on_year"] = i.InvestmentCompletedAmountYearOnYear
	i.fieldMap["primary_industry"] = i.PrimaryIndustry
	i.fieldMap["primary_industry_year_on_year"] = i.PrimaryIndustryYearOnYear
	i.fieldMap["secondary_industry"] = i.SecondaryIndustry
	i.fieldMap["secondary_industry_year_on_year"] = i.SecondaryIndustryYearOnYear
	i.fieldMap["tertiary_industry"] = i.TertiaryIndustry
	i.fieldMap["tertiary_industry_year_on_year"] = i.TertiaryIndustryYearOnYear
	i.fieldMap["real_estate_development"] = i.RealEstateDevelopment
	i.fieldMap["real_estate_development_year_on_year"] = i.RealEstateDevelopmentYearOnYear
}

func (i investmentInFixedAssets) clone(db *gorm.DB) investmentInFixedAssets {
	i.investmentInFixedAssetsDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i investmentInFixedAssets) replaceDB(db *gorm.DB) investmentInFixedAssets {
	i.investmentInFixedAssetsDo.ReplaceDB(db)
	return i
}

type investmentInFixedAssetsDo struct{ gen.DO }

func (i investmentInFixedAssetsDo) Debug() *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Debug())
}

func (i investmentInFixedAssetsDo) WithContext(ctx context.Context) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i investmentInFixedAssetsDo) ReadDB() *investmentInFixedAssetsDo {
	return i.Clauses(dbresolver.Read)
}

func (i investmentInFixedAssetsDo) WriteDB() *investmentInFixedAssetsDo {
	return i.Clauses(dbresolver.Write)
}

func (i investmentInFixedAssetsDo) Session(config *gorm.Session) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Session(config))
}

func (i investmentInFixedAssetsDo) Clauses(conds ...clause.Expression) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i investmentInFixedAssetsDo) Returning(value interface{}, columns ...string) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i investmentInFixedAssetsDo) Not(conds ...gen.Condition) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i investmentInFixedAssetsDo) Or(conds ...gen.Condition) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i investmentInFixedAssetsDo) Select(conds ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i investmentInFixedAssetsDo) Where(conds ...gen.Condition) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i investmentInFixedAssetsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *investmentInFixedAssetsDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i investmentInFixedAssetsDo) Order(conds ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i investmentInFixedAssetsDo) Distinct(cols ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i investmentInFixedAssetsDo) Omit(cols ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i investmentInFixedAssetsDo) Join(table schema.Tabler, on ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i investmentInFixedAssetsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i investmentInFixedAssetsDo) RightJoin(table schema.Tabler, on ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i investmentInFixedAssetsDo) Group(cols ...field.Expr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i investmentInFixedAssetsDo) Having(conds ...gen.Condition) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i investmentInFixedAssetsDo) Limit(limit int) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i investmentInFixedAssetsDo) Offset(offset int) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i investmentInFixedAssetsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i investmentInFixedAssetsDo) Unscoped() *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Unscoped())
}

func (i investmentInFixedAssetsDo) Create(values ...*model.InvestmentInFixedAssets) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i investmentInFixedAssetsDo) CreateInBatches(values []*model.InvestmentInFixedAssets, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i investmentInFixedAssetsDo) Save(values ...*model.InvestmentInFixedAssets) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i investmentInFixedAssetsDo) First() (*model.InvestmentInFixedAssets, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvestmentInFixedAssets), nil
	}
}

func (i investmentInFixedAssetsDo) Take() (*model.InvestmentInFixedAssets, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvestmentInFixedAssets), nil
	}
}

func (i investmentInFixedAssetsDo) Last() (*model.InvestmentInFixedAssets, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvestmentInFixedAssets), nil
	}
}

func (i investmentInFixedAssetsDo) Find() ([]*model.InvestmentInFixedAssets, error) {
	result, err := i.DO.Find()
	return result.([]*model.InvestmentInFixedAssets), err
}

func (i investmentInFixedAssetsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InvestmentInFixedAssets, err error) {
	buf := make([]*model.InvestmentInFixedAssets, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i investmentInFixedAssetsDo) FindInBatches(result *[]*model.InvestmentInFixedAssets, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i investmentInFixedAssetsDo) Attrs(attrs ...field.AssignExpr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i investmentInFixedAssetsDo) Assign(attrs ...field.AssignExpr) *investmentInFixedAssetsDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i investmentInFixedAssetsDo) Joins(fields ...field.RelationField) *investmentInFixedAssetsDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i investmentInFixedAssetsDo) Preload(fields ...field.RelationField) *investmentInFixedAssetsDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i investmentInFixedAssetsDo) FirstOrInit() (*model.InvestmentInFixedAssets, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvestmentInFixedAssets), nil
	}
}

func (i investmentInFixedAssetsDo) FirstOrCreate() (*model.InvestmentInFixedAssets, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvestmentInFixedAssets), nil
	}
}

func (i investmentInFixedAssetsDo) FindByPage(offset int, limit int) (result []*model.InvestmentInFixedAssets, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i investmentInFixedAssetsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i investmentInFixedAssetsDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i investmentInFixedAssetsDo) Delete(models ...*model.InvestmentInFixedAssets) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *investmentInFixedAssetsDo) withDO(do gen.Dao) *investmentInFixedAssetsDo {
	i.DO = *do.(*gen.DO)
	return i
}