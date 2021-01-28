package keeper

import (
	"context"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	govkeeper "github.com/TsukiCore/tsuki/x/gov/keeper"
	customgovtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/staking/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper    Keeper
	govKeeper govkeeper.Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, govKeeper govkeeper.Keeper) types.MsgServer {
	return &msgServer{
		keeper:    keeper,
		govKeeper: govKeeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ClaimValidator(goCtx context.Context, msg *types.MsgClaimValidator) (*types.MsgClaimValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := govkeeper.CheckIfAllowedPermission(ctx, k.govKeeper, sdk.AccAddress(msg.ValKey), customgovtypes.PermClaimValidator)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, "PermClaimValidator")
	}

	pk, ok := msg.PubKey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrap(errors.ErrInvalidPubKey, "Invalid pub key")
	}

	validator, err := types.NewValidator(msg.Moniker, msg.Website, msg.Social, msg.Identity, msg.Commission, msg.ValKey, pk)
	if err != nil {
		return nil, err
	}

	k.keeper.AddPendingValidator(ctx, validator)

	return &types.MsgClaimValidatorResponse{}, nil
}

func (k msgServer) ProposalUnjailValidator(goCtx context.Context, msg *types.MsgProposalUnjailValidator) (*types.MsgProposalUnjailValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := govkeeper.CheckIfAllowedPermission(ctx, k.govKeeper, msg.Proposer, customgovtypes.PermCreateUnjailValidatorProposal)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, customgovtypes.PermCreateUnjailValidatorProposal.String())
	}

	proposalID, err := k.CreateAndSaveProposalWithContent(ctx, types.NewProposalUnjailValidator(
		msg.Proposer,
	))
	if err != nil {
		return nil, err
	}

	return &types.MsgProposalUnjailValidatorResponse{
		ProposalID: proposalID,
	}, nil
}

func (k msgServer) CreateAndSaveProposalWithContent(ctx sdk.Context, content customgovtypes.Content) (uint64, error) {
	blockTime := ctx.BlockTime()
	proposalID, err := k.govKeeper.GetNextProposalID(ctx)
	if err != nil {
		return 0, err
	}

	properties := k.govKeeper.GetNetworkProperties(ctx)

	proposal, err := customgovtypes.NewProposal(
		proposalID,
		content,
		blockTime,
		blockTime.Add(time.Minute*time.Duration(properties.ProposalEndTime)),
		blockTime.Add(time.Minute*time.Duration(properties.ProposalEnactmentTime)),
	)

	k.govKeeper.SaveProposal(ctx, proposal)
	k.govKeeper.AddToActiveProposals(ctx, proposal)

	return proposalID, nil
}
