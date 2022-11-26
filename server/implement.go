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
