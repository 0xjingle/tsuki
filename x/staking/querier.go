package staking

import (
	"context"

	sdk "github.com/TsukiCore/cosmos-sdk/types"

	"github.com/TsukiCore/tsuki/x/staking/keeper"
	"github.com/TsukiCore/tsuki/x/staking/types"
)

type Querier struct {
	keeper keeper.Keeper
}

func NewQuerier(keeper keeper.Keeper) types.QueryServer {
	return &Querier{keeper: keeper}
}

func (q Querier) ValidatorByAddress(ctx context.Context, request *types.ValidatorByAddressRequest) (*types.ValidatorResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	return &types.ValidatorResponse{
		Validator: q.keeper.GetValidator(c, request.ValAddr),
	}, nil
}

func (q Querier) ValidatorByMoniker(ctx context.Context, request *types.ValidatorByMonikerRequest) (*types.ValidatorResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	return &types.ValidatorResponse{
		Validator: q.keeper.GetValidatorByMoniker(c, request.Moniker),
	}, nil
}
