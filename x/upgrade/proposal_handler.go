package upgrade

import (
	"github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/upgrade/keeper"
	upgradetypes "github.com/TsukiCore/tsuki/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplySoftwareUpgradeProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySoftwareUpgradeProposalHandler(keeper keeper.Keeper) *ApplySoftwareUpgradeProposalHandler {
	return &ApplySoftwareUpgradeProposalHandler{
		keeper: keeper,
	}
}

func (a ApplySoftwareUpgradeProposalHandler) ProposalType() string {
	return upgradetypes.ProposalTypeSoftwareUpgrade
}

func (a ApplySoftwareUpgradeProposalHandler) Apply(ctx sdk.Context, proposal types.Content) error {
	p := proposal.(*upgradetypes.ProposalSoftwareUpgrade)

	plan := upgradetypes.NewUpgradePlan(p.Name, p.Resources, p.MinUpgradeTime, p.MaxEnrolmentDuration, p.RollbackChecksum, p.InstateUpgrade)
	a.keeper.SaveUpgradePlan(ctx, plan)
	return nil
}

type ApplyCancelSoftwareUpgradeProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCancelSoftwareUpgradeProposalHandler(keeper keeper.Keeper) *ApplyCancelSoftwareUpgradeProposalHandler {
	return &ApplyCancelSoftwareUpgradeProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyCancelSoftwareUpgradeProposalHandler) ProposalType() string {
	return upgradetypes.ProposalTypeCancelSoftwareUpgrade
}

func (a ApplyCancelSoftwareUpgradeProposalHandler) Apply(ctx sdk.Context, proposal types.Content) error {
	a.keeper.ClearUpgradePlan(ctx)
	return nil
}
