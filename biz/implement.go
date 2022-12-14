// 业务层 只处理业务数据

package biz

import (
	"context"

	"github.com/chenhaonan-eth/dao/config"
	"github.com/chenhaonan-eth/dao/dal/query"
	pb "github.com/chenhaonan-eth/dao/proto/server"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

var q = query.Q

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) GetSocialFinancingStock(ctx context.Context, r *emptypb.Empty) (*pb.SocialFinancingStockResponse, error) {
	config.G_LOG.Debug("Start GetSocialFinancingStock ...")
	t := q.SocialFinancingStock
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.SocialFinancingStockResponse)
	resp.Results = make([]*pb.SocialFinancingStock, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.SocialFinancingStock{
			Date:                            v.Date,
			Ndbab:                           v.Ndbab,
			Entrustloan:                     v.Entrustloan,
			Forcloan:                        v.Forcloan,
			Rmblaon:                         v.Rmblaon,
			Bibae:                           v.Bibae,
			Tiosfs:                          v.Tiosfs,
			Sfinfe:                          v.Sfinfe,
			Trustloan:                       v.Trustloan,
			AssetBackedSecurities:           v.AssetBackedSecurities,
			LoansWrittenOff:                 v.LoansWrittenOff,
			GovernmentBonds:                 v.GovernmentBonds,
			Ndbabgrowthrate:                 v.NdbabGrowthRate,
			Entrustloangrowthrate:           v.EntrustloanGrowthRate,
			Forcloangrowthrate:              v.ForcloanGrowthRate,
			Rmblaongrowthrate:               v.RmblaonGrowthRate,
			Bibaegrowthrate:                 v.BibaeGrowthRate,
			Tiosfsgrowthrate:                v.TiosfsGrowthRate,
			Sfinfegrowthrate:                v.SfinfeGrowthRate,
			Trustloangrowthrate:             v.TrustloanGrowthRate,
			AssetBackedSecuritiesgrowthrate: v.AssetBackedSecuritiesGrowthRate,
			LoansWrittenOffgrowthrate:       v.LoansWrittenOffGrowthRate,
			GovernmentBondsgrowthrate:       v.GovernmentBondsGrowthRate,
		})
	}
	config.G_LOG.Debug("End GetSocialFinancingStock ...Results ", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetFuturesForeignHist(ctx context.Context, r *pb.FturesFoewignRequest) (*pb.FturesFoewignResponse, error) {
	config.G_LOG.Debug("Start GetFuturesForeignHist ...")
	t := q.FturesFoewign
	do := t.WithContext(context.Background())
	results, err := do.Where(t.Symbol.Eq(r.Symbol)).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.FturesFoewignResponse)
	resp.Results = make([]*pb.FturesFoewign, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.FturesFoewign{
			Date:   v.Date,
			Symbol: v.Symbol,
			Open:   v.Open,
			High:   v.High,
			Low:    v.Low,
			Lose:   v.Low,
			Volume: v.Volume,
		})
	}
	config.G_LOG.Debug("End GetFuturesForeignHist ...Results ", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetPpi(ctx context.Context, r *emptypb.Empty) (*pb.PpiResponse, error) {
	config.G_LOG.Debug("Start GetPpi ...")
	t := q.ChinaPPI
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.PpiResponse)
	resp.Results = make([]*pb.Ppi, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.Ppi{
			Date:         v.Date,
			Time:         v.Time,
			Ppi:          v.PPI,
			YearOnYear:   v.YearOnYear,
			Accumulative: v.Accumulative,
		})
	}
	config.G_LOG.Debug("End GetPpi ...Results ", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetGdp(ctx context.Context, r *emptypb.Empty) (*pb.GdpResponse, error) {
	config.G_LOG.Debug("Start GetGdp ...")
	t := q.ChinaGDP
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.GdpResponse)
	resp.Results = make([]*pb.Gdp, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.Gdp{
			Date:                          v.Date,
			Time:                          v.Time,
			Gdp:                           v.GDP,
			GdpYearOnYear:                 v.GDPYearOnYear,
			PrimaryIndustry:               v.PrimaryIndustry,
			PrimaryIndustryYearOnYear:     v.PrimaryIndustryYearOnYear,
			SecondaryIndustries:           v.SecondaryIndustries,
			SecondaryIndustriesYearOnYear: v.SecondaryIndustriesYearOnYear,
			TertiaryIndustry:              v.TertiaryIndustry,
			TertiaryIndustryYearOnYear:    v.TertiaryIndustryYearOnYear,
		})
	}
	config.G_LOG.Debug("End GetGdp ...Results ", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetPmi(ctx context.Context, r *emptypb.Empty) (*pb.PmiResponse, error) {
	config.G_LOG.Debug("Start GetPmi ...")
	t := q.ChinaPMI
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.PmiResponse)
	resp.Results = make([]*pb.Pmi, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.Pmi{
			Date:                       v.Date,
			Time:                       v.Time,
			Manufacturing:              v.Manufacturing,
			ManufacturingYearOnYear:    v.ManufacturingYearOnYear,
			NonManufacturing:           v.NonManufacturing,
			NonManufacturingYearOnYear: v.NonManufacturingYearOnYear,
		})
	}
	config.G_LOG.Debug("End GetPmi ...Results ", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetCpi(ctx context.Context, r *emptypb.Empty) (*pb.CpiResponse, error) {
	config.G_LOG.Debug("Start GetCpi ...")
	t := q.ChinaCPI
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.CpiResponse)
	resp.Results = make([]*pb.Cpi, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.Cpi{
			Date:                 v.Date,
			Time:                 v.Time,
			National:             v.National,
			NationalYearOnYear:   v.NationalYearOnYear,
			NationalYearOverYear: v.NationalYearOverYear,
			NationalAccumulative: v.NationalAccumulative,
			City:                 v.City,
			CityYearOnYear:       v.CityYearOnYear,
			CityYearOverYear:     v.CityYearOverYear,
			CityAccumulative:     v.CityAccumulative,
			Rural:                v.Rural,
			RuralYearOnYear:      v.RuralYearOnYear,
			RuralYearOverYear:    v.RuralYearOverYear,
			RuralAccumulative:    v.RuralAccumulative,
		})
	}
	config.G_LOG.Debug("End GetCpi ...Results ", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetBondZhUsRate(ctx context.Context, r *emptypb.Empty) (*pb.BondZhUsRateResponse, error) {
	config.G_LOG.Debug("start GetBondZhUsRate ... ")
	t := q.BondZhUsRate
	do := t.WithContext(context.Background())
	results, err := do.Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.BondZhUsRateResponse)
	resp.Results = make([]*pb.BondZhUsRate, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.BondZhUsRate{
			Date:        v.Date,
			CN2Years:    v.CN2Years,
			CN5Years:    v.CN5Years,
			CN10Years:   v.CN10Years,
			CN30Years:   v.CN30Years,
			CN10_2Years: v.CN10_2Years,
			US2Years:    v.US2Years,
			US5Years:    v.US5Years,
			US10Years:   v.US10Years,
			US30Years:   v.US30Years,
			US10_2Years: v.US10_2Years,
		})
	}
	config.G_LOG.Debug("End GetPmi ...Results ", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetTotalSocialFlows(ctx context.Context, r *emptypb.Empty) (*pb.SocialFinancingFlowsResponse, error) {
	config.G_LOG.Debug("Start GetTotalSocialFlows ...")
	t := q.SocialFinancingFlow
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.SocialFinancingFlowsResponse)
	resp.Results = make([]*pb.SocialFinancingFlow, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.SocialFinancingFlow{
			Date:        v.Date,
			Ndbab:       v.Ndbab,
			Entrustloan: v.Entrustloan,
			Forcloan:    v.Entrustloan,
			Rmblaon:     v.Rmblaon,
			Bibae:       v.Bibae,
			Tiosfs:      v.Tiosfs,
			Sfinfe:      v.Sfinfe,
			Trustloan:   v.Trustloan,
		})
	}
	config.G_LOG.Debug("End GetTotalSocialFlows ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetSH300PE(ctx context.Context, r *emptypb.Empty) (*pb.SH300PEResponse, error) {
	config.G_LOG.Debug("Start GetSH300PE ...")
	t := q.SH300PE
	do := t.WithContext(context.Background())
	results, err := do.Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.SH300PEResponse)
	resp.Results = make([]*pb.SH300PE, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.SH300PE{
			Time:        v.Time,
			Date:        v.Date,
			MiddleLyrPe: v.MiddleLyrPe,
			LyrPe:       v.LyrPe,
			AddLyrPe:    v.AddLyrPe,
			MiddleTtmPe: v.MiddleLyrPe,
			TtmPe:       v.TtmPe,
			AddTtmPe:    v.AddTtmPe,
		})
	}
	config.G_LOG.Debug("End GetSH300PE ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetMoneySupply(ctx context.Context, r *emptypb.Empty) (*pb.MoneySupplyResponse, error) {
	config.G_LOG.Debug("Start GetMoneySupply ...")
	t := q.MacroChinaMoneySupply
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.MoneySupplyResponse)
	resp.Results = make([]*pb.MoneySupply, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.MoneySupply{
			Date:           v.Date,
			M2:             v.M2,
			M2YearOnYear:   v.M2YearOnYear,
			M2YearOverYear: v.M2YearOverYear,
			M1:             v.M1,
			M1YearOnYear:   v.M1YearOnYear,
			M1YearOverYear: v.M1YearOverYear,
			M0:             v.M0,
			M0YearOnYear:   v.M0YearOnYear,
			M0YearOverYear: v.M0YearOverYear,
		})
	}
	config.G_LOG.Debug("End GetMoneySupply ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}
func (s *server) GetConsumerGoodsRetail(ctx context.Context, r *emptypb.Empty) (*pb.ConsumerGoodsRetailResponse, error) {
	config.G_LOG.Debug("Start GetConsumerGoodsRetail ...")
	t := q.MacroChinaConsumerGoodsRetail
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.ConsumerGoodsRetailResponse)
	resp.Results = make([]*pb.ConsumerGoodsRetail, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.ConsumerGoodsRetail{
			Date:                        v.Date,
			TotalRetailSales:            v.TotalRetailSales,
			YearOnYear:                  v.YearOnYear,
			YearOverYear:                v.YearOverYear,
			TotalAccumulation:           v.TotalAccumulation,
			TotalAccumulationYearOnYear: v.TotalAccumulationYearOnYear,
		})
	}
	config.G_LOG.Debug("End GetConsumerGoodsRetail ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetCxPmi(ctx context.Context, r *emptypb.Empty) (*pb.CxPmiResponse, error) {
	config.G_LOG.Debug("Start GetCxPmi ...")
	t := q.PmiCx
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.CxPmiResponse)
	resp.Results = make([]*pb.CxPmi, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.CxPmi{
			Date:                    v.Date,
			Time:                    v.Time,
			Manufacture:             v.Manufacture,
			ManufactureYearOverYear: v.ManufactureYearOverYear,
			Service:                 v.Service,
			ServiceYearOverYear:     v.ServiceYearOverYear,
			Synthesis:               v.Synthesis,
			SynthesisYearOverYear:   v.SynthesisYearOverYear,
		})
	}
	config.G_LOG.Debug("End GetCxPmi ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetValueAddedOfIndustrialProduction(ctx context.Context, r *emptypb.Empty) (*pb.ValueAddedOfIndustrialProductionResponse, error) {
	config.G_LOG.Debug("Start GetValueAddedOfIndustrialProduction ...")
	t := q.ValueAddedOfIndustrialProduction
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.ValueAddedOfIndustrialProductionResponse)
	resp.Results = make([]*pb.ValueAddedOfIndustrialProduction, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.ValueAddedOfIndustrialProduction{
			Date:             v.Date,
			Time:             v.Time,
			YearOnYear:       v.YearOnYear,
			CumulativeGrowth: v.CumulativeGrowth,
		})
	}
	config.G_LOG.Debug("End GetValueAddedOfIndustrialProduction ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetSocialElectricityConsumption(ctx context.Context, r *emptypb.Empty) (*pb.SocialElectricityConsumptionResponse, error) {
	config.G_LOG.Debug("Start GetSocialElectricityConsumption ...")
	t := q.SocialElectricityConsumption
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.SocialElectricityConsumptionResponse)
	resp.Results = make([]*pb.SocialElectricityConsumption, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.SocialElectricityConsumption{
			Date:                        v.Date,
			WholeSociety:                v.WholeSociety,
			WholeSocietyYearOnYear:      v.WholeSocietyYearOnYear,
			AllIndustries:               v.AllIndustries,               //各行业用电量
			AllIndustriesYearOnYear:     v.AllIndustriesYearOnYear,     //各行业用电量合计同比 %
			PrimaryIndustry:             v.PrimaryIndustry,             //第一产业用电量 万千瓦时
			PrimaryIndustryYearOnYear:   v.PrimaryIndustryYearOnYear,   //第一产业用电量同比%
			SecondaryIndustry:           v.SecondaryIndustry,           //第二产业用电量 万千瓦时
			SecondaryIndustryYearOnYear: v.SecondaryIndustryYearOnYear, //第二产业用电量同比%
			TertiaryIndustry:            v.TertiaryIndustry,            //第三产业用电量万千瓦时
			TertiaryIndustryYearOnYear:  v.TertiaryIndustryYearOnYear,  //第三产业用电量同比%
			CitiesAndVillages:           v.CitiesAndVillages,           //城乡居民生活用电量合计/ 万千瓦时
			CitiesAndVillagesYearOnYear: v.CitiesAndVillagesYearOnYear, //城乡居民生活用电量合计同比 %
			Cities:                      v.Cities,                      //城镇居民用电量 万千瓦时
			CitiesYearOnYear:            v.CitiesYearOnYear,            //城镇居民用电量同比%
			Villages:                    v.Villages,                    //乡村居民用电量 万千瓦时
			VillagesYearOnYear:          v.VillagesYearOnYear,          //乡村居民用电量同比 %

		})
	}
	config.G_LOG.Debug("End GetSocialElectricityConsumption ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}
func (s *server) GetPassengerAndFreightTraffic(ctx context.Context, r *emptypb.Empty) (*pb.PassengerAndFreightTrafficResponse, error) {
	config.G_LOG.Debug("Start GetPassengerAndFreightTraffic ...")
	t := q.PassengerAndFreightTraffic
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.PassengerAndFreightTrafficResponse)
	resp.Results = make([]*pb.PassengerAndFreightTraffic, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.PassengerAndFreightTraffic{
			Date:                               v.Date,
			Class:                              pb.CategoryOfTraffic(v.Class),        //运输种类
			FreightVolume:                      v.FreightVolume,                      //货运量/亿吨
			FreightVolumeYearOnYear:            v.FreightVolumeYearOnYear,            //货运量同比增长 %
			FreightTurnover:                    v.FreightTurnover,                    //货物周转量/亿
			FreightTurnoverYearOnYear:          v.FreightTurnoverYearOnYear,          //公里货物周转量同比增长 %
			PassengerCapacity:                  v.PassengerCapacity,                  //客运量/亿人
			PassengerCapacityYearOnYear:        v.PassengerCapacityYearOnYear,        //客运量同比增长 %
			PassengerTurnover:                  v.PassengerTurnover,                  //旅客周转量/亿
			PassengerTurnoverYearOnYear:        v.PassengerTurnoverYearOnYear,        //公里旅客周转量同比增长/%
			CargoThroughputOfMajorCoastalPorts: v.CargoThroughputOfMajorCoastalPorts, //沿海主要港口货物吞吐量/亿吨
			CargoThroughputOfMajorCoastalPortsYearOnYear: v.CargoThroughputOfMajorCoastalPortsYearOnYear, //沿海主要港口货物吞吐量同比增长 %
			ForeignTradeCargoThroughput:                  v.ForeignTradeCargoThroughput,                  //其中:外贸货物吞吐量 /亿吨
			ForeignTradeCargoThroughputYearOnYear:        v.ForeignTradeCargoThroughputYearOnYear,        //其中:外贸货物吞吐量同比增长 %
			TotalTurnoverOfCivilAviation:                 v.TotalTurnoverOfCivilAviation,                 //民航总周转量 /亿
			KmTotalTurnoverOfCivilAviation:               v.KmTotalTurnoverOfCivilAviation,               //公里民航总周转/%
		})
	}
	config.G_LOG.Debug("End GetPassengerAndFreightTraffic ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetNewFinancialCredit(ctx context.Context, r *emptypb.Empty) (*pb.NewFinancialCreditResponse, error) {
	config.G_LOG.Debug("Start GetNewFinancialCredit ...")
	t := q.NewFinancialCredit
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.NewFinancialCreditResponse)
	resp.Results = make([]*pb.NewFinancialCredit, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.NewFinancialCredit{
			Date:               v.Date,
			Time:               v.Time,
			Loany:              v.Loany,
			LoanYearOnYear:     v.LoanYearOnYear,
			LoanyYearOverYear:  v.LoanyYearOverYear,
			LoanyAcc:           v.LoanyAcc,
			LoanyAccYearOnYear: v.LoanyAccYearOnYear,
		})
	}
	config.G_LOG.Debug("End GetNewFinancialCredit ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetForeignReserveAndGold(ctx context.Context, r *emptypb.Empty) (*pb.ForeignReserveAndGoldResponse, error) {
	config.G_LOG.Debug("Start GetNewFinancialCredit ...")
	t := q.ForeignReserveAndGold
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.ForeignReserveAndGoldResponse)
	resp.Results = make([]*pb.ForeignReserveAndGold, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.ForeignReserveAndGold{
			Date:           v.Date,
			Gold:           v.Gold,
			ForeignReserve: v.ForeignReserve,
		})
	}
	config.G_LOG.Debug("End GetNewFinancialCredit ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetInvestmentInFixedAssets(ctx context.Context, r *emptypb.Empty) (*pb.InvestmentInFixedAssetsResponse, error) {
	config.G_LOG.Debug("Start GetInvestmentInFixedAssets ...")
	t := q.InvestmentInFixedAssets
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.InvestmentInFixedAssetsResponse)
	resp.Results = make([]*pb.InvestmentInFixedAssets, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.InvestmentInFixedAssets{
			Date:                                v.Date,
			InvestmentCompletedAmount:           v.InvestmentCompletedAmount,           //投资完成额(亿元)
			InvestmentCompletedAmountYearOnYear: v.InvestmentCompletedAmountYearOnYear, //投资完成额同比增长(%)
			PrimaryIndustry:                     v.PrimaryIndustry,                     //第一产业(亿元)
			PrimaryIndustryYearOnYear:           v.PrimaryIndustryYearOnYear,           //第一产业用同比%
			SecondaryIndustry:                   v.SecondaryIndustry,                   //第二产业
			SecondaryIndustryYearOnYear:         v.SecondaryIndustryYearOnYear,         //第二产业同比%
			TertiaryIndustry:                    v.TertiaryIndustry,                    //第三产业
			TertiaryIndustryYearOnYear:          v.TertiaryIndustryYearOnYear,          //第三产业同比%
			RealEstateDevelopment:               v.RealEstateDevelopment,               //房地产开发(亿元)
			RealEstateDevelopmentYearOnYear:     v.RealEstateDevelopmentYearOnYear,     //房地产开发同比增长(%)
		})
	}
	config.G_LOG.Debug("End GetInvestmentInFixedAssets ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}

func (s *server) GetManufacturingPmiParticulars(ctx context.Context, r *emptypb.Empty) (*pb.ManufacturingPmiParticularsResponse, error) {
	config.G_LOG.Debug("Start GetManufacturingPmiParticulars ...")
	t := q.ManufacturingPmiParticulars
	do := t.WithContext(context.Background())
	results, err := do.Order(t.Date.Desc()).Find()
	if err != nil {
		config.G_LOG.Error("Find err ", zap.Error(err))
		return nil, err
	}
	resp := new(pb.ManufacturingPmiParticularsResponse)
	resp.Results = make([]*pb.ManufacturingPmiParticulars, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.ManufacturingPmiParticulars{
			Date:                              v.Date,
			Pmi:                               v.Pmi,
			ProductionIndex:                   v.ProductionIndex,                   //生产指数
			IndexOfNewOrders:                  v.IndexOfNewOrders,                  //新订单指数
			IndexOfNewExportOrders:            v.IndexOfNewExportOrders,            //新出口订单指数
			BacklogIndex:                      v.BacklogIndex,                      //积压订单指数
			IndexOfInventoriesOfFinishedGoods: v.IndexOfInventoriesOfFinishedGoods, //产成品库存指数
			PurchasingVolumeIndex:             v.PurchasingVolumeIndex,             //采购量指数
			ImportIndex:                       v.ImportIndex,                       //进口指数
			PurchasingPriceIndex:              v.PurchasingPriceIndex,              //购进价格指数
			InventoryIndexOfRawMaterials:      v.InventoryIndexOfRawMaterials,      //原材料库存指数
			EmployeeIndex:                     v.EmployeeIndex,                     //从业人员指数
			SupplierDeliveryTimeIndex:         v.SupplierDeliveryTimeIndex,         //供应商配送时间指数

		})
	}
	config.G_LOG.Debug("End GetManufacturingPmiParticulars ...", zap.Any("len", len(resp.Results)))
	return resp, nil
}
