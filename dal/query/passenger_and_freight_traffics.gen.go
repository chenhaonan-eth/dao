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

func newPassengerAndFreightTraffic(db *gorm.DB, opts ...gen.DOOption) passengerAndFreightTraffic {
	_passengerAndFreightTraffic := passengerAndFreightTraffic{}

	_passengerAndFreightTraffic.passengerAndFreightTrafficDo.UseDB(db, opts...)
	_passengerAndFreightTraffic.passengerAndFreightTrafficDo.UseModel(&model.PassengerAndFreightTraffic{})

	tableName := _passengerAndFreightTraffic.passengerAndFreightTrafficDo.TableName()
	_passengerAndFreightTraffic.ALL = field.NewAsterisk(tableName)
	_passengerAndFreightTraffic.Date = field.NewString(tableName, "date")
	_passengerAndFreightTraffic.Class = field.NewInt64(tableName, "class")
	_passengerAndFreightTraffic.FreightVolume = field.NewString(tableName, "freight_volume")
	_passengerAndFreightTraffic.FreightVolumeYearOnYear = field.NewString(tableName, "freight_volume_year_on_year")
	_passengerAndFreightTraffic.FreightTurnover = field.NewString(tableName, "freight_turnover")
	_passengerAndFreightTraffic.FreightTurnoverYearOnYear = field.NewString(tableName, "freight_turnover_year_on_year")
	_passengerAndFreightTraffic.PassengerCapacity = field.NewString(tableName, "passenger_capacity")
	_passengerAndFreightTraffic.PassengerCapacityYearOnYear = field.NewString(tableName, "passenger_capacity_year_on_year")
	_passengerAndFreightTraffic.PassengerTurnover = field.NewString(tableName, "passenger_turnover")
	_passengerAndFreightTraffic.PassengerTurnoverYearOnYear = field.NewString(tableName, "passenger_turnover_year_on_year")
	_passengerAndFreightTraffic.CargoThroughputOfMajorCoastalPorts = field.NewString(tableName, "cargo_throughput_of_major_coastal_ports")
	_passengerAndFreightTraffic.CargoThroughputOfMajorCoastalPortsYearOnYear = field.NewString(tableName, "cargo_throughput_of_major_coastal_ports_year_on_year")
	_passengerAndFreightTraffic.ForeignTradeCargoThroughput = field.NewString(tableName, "foreign_trade_cargo_throughput")
	_passengerAndFreightTraffic.ForeignTradeCargoThroughputYearOnYear = field.NewString(tableName, "foreign_trade_cargo_throughput_year_on_year")
	_passengerAndFreightTraffic.TotalTurnoverOfCivilAviation = field.NewString(tableName, "total_turnover_of_civil_aviation")
	_passengerAndFreightTraffic.KmTotalTurnoverOfCivilAviation = field.NewString(tableName, "km_total_turnover_of_civil_aviation")

	_passengerAndFreightTraffic.fillFieldMap()

	return _passengerAndFreightTraffic
}

type passengerAndFreightTraffic struct {
	passengerAndFreightTrafficDo passengerAndFreightTrafficDo

	ALL                                          field.Asterisk
	Date                                         field.String
	Class                                        field.Int64
	FreightVolume                                field.String
	FreightVolumeYearOnYear                      field.String
	FreightTurnover                              field.String
	FreightTurnoverYearOnYear                    field.String
	PassengerCapacity                            field.String
	PassengerCapacityYearOnYear                  field.String
	PassengerTurnover                            field.String
	PassengerTurnoverYearOnYear                  field.String
	CargoThroughputOfMajorCoastalPorts           field.String
	CargoThroughputOfMajorCoastalPortsYearOnYear field.String
	ForeignTradeCargoThroughput                  field.String
	ForeignTradeCargoThroughputYearOnYear        field.String
	TotalTurnoverOfCivilAviation                 field.String
	KmTotalTurnoverOfCivilAviation               field.String

	fieldMap map[string]field.Expr
}

func (p passengerAndFreightTraffic) Table(newTableName string) *passengerAndFreightTraffic {
	p.passengerAndFreightTrafficDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p passengerAndFreightTraffic) As(alias string) *passengerAndFreightTraffic {
	p.passengerAndFreightTrafficDo.DO = *(p.passengerAndFreightTrafficDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *passengerAndFreightTraffic) updateTableName(table string) *passengerAndFreightTraffic {
	p.ALL = field.NewAsterisk(table)
	p.Date = field.NewString(table, "date")
	p.Class = field.NewInt64(table, "class")
	p.FreightVolume = field.NewString(table, "freight_volume")
	p.FreightVolumeYearOnYear = field.NewString(table, "freight_volume_year_on_year")
	p.FreightTurnover = field.NewString(table, "freight_turnover")
	p.FreightTurnoverYearOnYear = field.NewString(table, "freight_turnover_year_on_year")
	p.PassengerCapacity = field.NewString(table, "passenger_capacity")
	p.PassengerCapacityYearOnYear = field.NewString(table, "passenger_capacity_year_on_year")
	p.PassengerTurnover = field.NewString(table, "passenger_turnover")
	p.PassengerTurnoverYearOnYear = field.NewString(table, "passenger_turnover_year_on_year")
	p.CargoThroughputOfMajorCoastalPorts = field.NewString(table, "cargo_throughput_of_major_coastal_ports")
	p.CargoThroughputOfMajorCoastalPortsYearOnYear = field.NewString(table, "cargo_throughput_of_major_coastal_ports_year_on_year")
	p.ForeignTradeCargoThroughput = field.NewString(table, "foreign_trade_cargo_throughput")
	p.ForeignTradeCargoThroughputYearOnYear = field.NewString(table, "foreign_trade_cargo_throughput_year_on_year")
	p.TotalTurnoverOfCivilAviation = field.NewString(table, "total_turnover_of_civil_aviation")
	p.KmTotalTurnoverOfCivilAviation = field.NewString(table, "km_total_turnover_of_civil_aviation")

	p.fillFieldMap()

	return p
}

func (p *passengerAndFreightTraffic) WithContext(ctx context.Context) *passengerAndFreightTrafficDo {
	return p.passengerAndFreightTrafficDo.WithContext(ctx)
}

func (p passengerAndFreightTraffic) TableName() string {
	return p.passengerAndFreightTrafficDo.TableName()
}

func (p passengerAndFreightTraffic) Alias() string { return p.passengerAndFreightTrafficDo.Alias() }

func (p *passengerAndFreightTraffic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *passengerAndFreightTraffic) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 16)
	p.fieldMap["date"] = p.Date
	p.fieldMap["class"] = p.Class
	p.fieldMap["freight_volume"] = p.FreightVolume
	p.fieldMap["freight_volume_year_on_year"] = p.FreightVolumeYearOnYear
	p.fieldMap["freight_turnover"] = p.FreightTurnover
	p.fieldMap["freight_turnover_year_on_year"] = p.FreightTurnoverYearOnYear
	p.fieldMap["passenger_capacity"] = p.PassengerCapacity
	p.fieldMap["passenger_capacity_year_on_year"] = p.PassengerCapacityYearOnYear
	p.fieldMap["passenger_turnover"] = p.PassengerTurnover
	p.fieldMap["passenger_turnover_year_on_year"] = p.PassengerTurnoverYearOnYear
	p.fieldMap["cargo_throughput_of_major_coastal_ports"] = p.CargoThroughputOfMajorCoastalPorts
	p.fieldMap["cargo_throughput_of_major_coastal_ports_year_on_year"] = p.CargoThroughputOfMajorCoastalPortsYearOnYear
	p.fieldMap["foreign_trade_cargo_throughput"] = p.ForeignTradeCargoThroughput
	p.fieldMap["foreign_trade_cargo_throughput_year_on_year"] = p.ForeignTradeCargoThroughputYearOnYear
	p.fieldMap["total_turnover_of_civil_aviation"] = p.TotalTurnoverOfCivilAviation
	p.fieldMap["km_total_turnover_of_civil_aviation"] = p.KmTotalTurnoverOfCivilAviation
}

func (p passengerAndFreightTraffic) clone(db *gorm.DB) passengerAndFreightTraffic {
	p.passengerAndFreightTrafficDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p passengerAndFreightTraffic) replaceDB(db *gorm.DB) passengerAndFreightTraffic {
	p.passengerAndFreightTrafficDo.ReplaceDB(db)
	return p
}

type passengerAndFreightTrafficDo struct{ gen.DO }

func (p passengerAndFreightTrafficDo) Debug() *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Debug())
}

func (p passengerAndFreightTrafficDo) WithContext(ctx context.Context) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p passengerAndFreightTrafficDo) ReadDB() *passengerAndFreightTrafficDo {
	return p.Clauses(dbresolver.Read)
}

func (p passengerAndFreightTrafficDo) WriteDB() *passengerAndFreightTrafficDo {
	return p.Clauses(dbresolver.Write)
}

func (p passengerAndFreightTrafficDo) Session(config *gorm.Session) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Session(config))
}

func (p passengerAndFreightTrafficDo) Clauses(conds ...clause.Expression) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p passengerAndFreightTrafficDo) Returning(value interface{}, columns ...string) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p passengerAndFreightTrafficDo) Not(conds ...gen.Condition) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p passengerAndFreightTrafficDo) Or(conds ...gen.Condition) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p passengerAndFreightTrafficDo) Select(conds ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p passengerAndFreightTrafficDo) Where(conds ...gen.Condition) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p passengerAndFreightTrafficDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *passengerAndFreightTrafficDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p passengerAndFreightTrafficDo) Order(conds ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p passengerAndFreightTrafficDo) Distinct(cols ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p passengerAndFreightTrafficDo) Omit(cols ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p passengerAndFreightTrafficDo) Join(table schema.Tabler, on ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p passengerAndFreightTrafficDo) LeftJoin(table schema.Tabler, on ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p passengerAndFreightTrafficDo) RightJoin(table schema.Tabler, on ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p passengerAndFreightTrafficDo) Group(cols ...field.Expr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p passengerAndFreightTrafficDo) Having(conds ...gen.Condition) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p passengerAndFreightTrafficDo) Limit(limit int) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p passengerAndFreightTrafficDo) Offset(offset int) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p passengerAndFreightTrafficDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p passengerAndFreightTrafficDo) Unscoped() *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Unscoped())
}

func (p passengerAndFreightTrafficDo) Create(values ...*model.PassengerAndFreightTraffic) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p passengerAndFreightTrafficDo) CreateInBatches(values []*model.PassengerAndFreightTraffic, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p passengerAndFreightTrafficDo) Save(values ...*model.PassengerAndFreightTraffic) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p passengerAndFreightTrafficDo) First() (*model.PassengerAndFreightTraffic, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerAndFreightTraffic), nil
	}
}

func (p passengerAndFreightTrafficDo) Take() (*model.PassengerAndFreightTraffic, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerAndFreightTraffic), nil
	}
}

func (p passengerAndFreightTrafficDo) Last() (*model.PassengerAndFreightTraffic, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerAndFreightTraffic), nil
	}
}

func (p passengerAndFreightTrafficDo) Find() ([]*model.PassengerAndFreightTraffic, error) {
	result, err := p.DO.Find()
	return result.([]*model.PassengerAndFreightTraffic), err
}

func (p passengerAndFreightTrafficDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PassengerAndFreightTraffic, err error) {
	buf := make([]*model.PassengerAndFreightTraffic, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p passengerAndFreightTrafficDo) FindInBatches(result *[]*model.PassengerAndFreightTraffic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p passengerAndFreightTrafficDo) Attrs(attrs ...field.AssignExpr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p passengerAndFreightTrafficDo) Assign(attrs ...field.AssignExpr) *passengerAndFreightTrafficDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p passengerAndFreightTrafficDo) Joins(fields ...field.RelationField) *passengerAndFreightTrafficDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p passengerAndFreightTrafficDo) Preload(fields ...field.RelationField) *passengerAndFreightTrafficDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p passengerAndFreightTrafficDo) FirstOrInit() (*model.PassengerAndFreightTraffic, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerAndFreightTraffic), nil
	}
}

func (p passengerAndFreightTrafficDo) FirstOrCreate() (*model.PassengerAndFreightTraffic, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerAndFreightTraffic), nil
	}
}

func (p passengerAndFreightTrafficDo) FindByPage(offset int, limit int) (result []*model.PassengerAndFreightTraffic, count int64, err error) {
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

func (p passengerAndFreightTrafficDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p passengerAndFreightTrafficDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p passengerAndFreightTrafficDo) Delete(models ...*model.PassengerAndFreightTraffic) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *passengerAndFreightTrafficDo) withDO(do gen.Dao) *passengerAndFreightTrafficDo {
	p.DO = *do.(*gen.DO)
	return p
}
