package multistaking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	govkeeper "github.com/TsukiCore/tsuki/x/gov/keeper"
	customkeeper "github.com/TsukiCore/tsuki/x/multistaking/keeper"
	"github.com/TsukiCore/tsuki/x/multistaking/types"
)

func NewHandler(ck customkeeper.Keeper, govkeeper govkeeper.Keeper) sdk.Handler {
	msgServer := customkeeper.NewMsgServerImpl(ck, govkeeper)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgUpsertStakingPool:
			res, err := msgServer.UpsertStakingPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDelegate:
			res, err := msgServer.Delegate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUndelegate:
			res, err := msgServer.Undelegate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgClaimRewards:
			res, err := msgServer.ClaimRewards(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgClaimUndelegation:
			res, err := msgServer.ClaimUndelegation(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, errors.Wrapf(errors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}
