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
	results, err := do.Find()
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
	results, err := do.Find()
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
			PPI:          v.PPI,
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
	results, err := do.Find()
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
			GDP:                           v.GDP,
			GDPYearOnYear:                 v.GDPYearOnYear,
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
	results, err := do.Find()
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
	results, err := do.Find()
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
	results, err := do.Find()
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
	config.G_LOG.Debug("End GetTotalSocialFlows ...")
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
			Date:        v.Date,
			MiddleLyrPe: v.MiddleLyrPe,
			LyrPe:       v.LyrPe,
			AddLyrPe:    v.AddLyrPe,
			MiddleTtmPe: v.MiddleLyrPe,
			TtmPe:       v.TtmPe,
			AddTtmPe:    v.AddTtmPe,
		})
	}
	config.G_LOG.Debug("End GetSH300PE ...")
	return resp, nil
}

func (s *server) GetMoneySupply(ctx context.Context, r *emptypb.Empty) (*pb.MoneySupplyResponse, error) {
	config.G_LOG.Debug("Start GetMoneySupply ...")
	t := q.MacroChinaMoneySupply
	do := t.WithContext(context.Background())
	results, err := do.Find()
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
	config.G_LOG.Debug("End GetMoneySupply ...")
	return resp, nil
}
func (s *server) GetConsumerGoodsRetail(ctx context.Context, r *emptypb.Empty) (*pb.ConsumerGoodsRetailResponse, error) {
	config.G_LOG.Debug("Start GetConsumerGoodsRetail ...")
	t := q.MacroChinaConsumerGoodsRetail
	do := t.WithContext(context.Background())
	results, err := do.Find()
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
	config.G_LOG.Debug("End GetConsumerGoodsRetail ...")
	return resp, nil
}
