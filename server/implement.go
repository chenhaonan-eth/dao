package server

import (
	"context"

	"github.com/chenhaonan-eth/dao/dal/query"
	pb "github.com/chenhaonan-eth/dao/proto/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	q = query.Q
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) GetSocialFinancingStock(ctx context.Context, r *emptypb.Empty) (*pb.SocialFinancingStockResponse, error) {
	t := q.SocialFinancingStock
	do := t.WithContext(context.Background())
	results, err := do.Find()
	if err != nil {
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
	return resp, nil
}

func (s *server) GetFuturesForeignHist(ctx context.Context, r *pb.FturesFoewignRequest) (*pb.FturesFoewignResponse, error) {
	t := q.FturesFoewign
	do := t.WithContext(context.Background())
	results, err := do.Where(t.Symbol.Eq(r.Symbol)).Find()
	if err != nil {
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
	return resp, nil
}

func (s *server) GetPpi(ctx context.Context, r *pb.PpiRequest) (*pb.PpiResponse, error) {
	t := q.MacroPpi
	do := t.WithContext(context.Background())
	results, err := do.Where(t.Country.Eq(r.Country)).Find()
	if err != nil {
		return nil, err
	}
	resp := new(pb.PpiResponse)
	resp.Results = make([]*pb.Ppi, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.Ppi{
			Date:    v.Date,
			Country: v.Country,
			Ppi:     v.Ppi,
		})
	}
	return resp, nil
}

func (s *server) GetGdp(ctx context.Context, r *pb.GdpRequest) (*pb.GdpResponse, error) {
	t := q.MacroGDP
	do := t.WithContext(context.Background())
	results, err := do.Where(t.Country.Eq(r.Country)).Find()
	if err != nil {
		return nil, err
	}
	resp := new(pb.GdpResponse)
	resp.Results = make([]*pb.Gdp, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.Gdp{
			Date:    v.Date,
			Country: v.Country,
			Gdp:     v.Gdp,
		})
	}
	return resp, nil
}

func (s *server) GetPmi(ctx context.Context, r *pb.PmiRequest) (*pb.PmiResponse, error) {
	t := q.MacroPMI
	do := t.WithContext(context.Background())
	results, err := do.Where(t.Country.Eq(r.Country)).Find()
	if err != nil {
		return nil, err
	}
	resp := new(pb.PmiResponse)
	resp.Results = make([]*pb.Pmi, 0)
	for _, v := range results {
		resp.Results = append(resp.Results, &pb.Pmi{
			Date:    v.Date,
			Country: v.Country,
			Pmi:     v.Pmi,
		})
	}
	return resp, nil
}
func (s *server) GetBondZhUsRate(ctx context.Context, r *emptypb.Empty) (*pb.BondZhUsRateResponse, error) {
	t := q.BondZhUsRate
	do := t.WithContext(context.Background())
	results, err := do.Find()
	if err != nil {
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
	return resp, nil
}
