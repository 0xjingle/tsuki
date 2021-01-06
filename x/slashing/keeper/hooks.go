package keeper

import (
	"time"

	"github.com/tendermint/tendermint/crypto"

	"github.com/TsukiCore/tsuki/x/slashing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitValidatorSigningInfo update the signing info start height or create a new signing info
func (k Keeper) InitValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress, _ sdk.ValAddress) {
	// Update the signing info start height or create a new signing info
	_, found := k.GetValidatorSigningInfo(ctx, address)
	if !found {
		signingInfo := types.NewValidatorSigningInfo(
			address,
			ctx.BlockHeight(),
			0,
			time.Unix(0, 0),
			false,
			0,
		)
		k.SetValidatorSigningInfo(ctx, address, signingInfo)
	}
}

// AfterValidatorCreated adds the address-pubkey relation when a validator is created.
func (k Keeper) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) error {
	validator, err := k.sk.GetValidator(ctx, valAddr)
	if err != nil {
		return err
	}

	consPk, err := validator.TmConsPubKey()
	if err != nil {
		return err
	}
	k.AddPubkey(ctx, consPk)

	// initialize validator signing info on validator creation
	k.InitValidatorSigningInfo(ctx, validator.GetConsAddr(), valAddr)
	return nil
}

// AfterValidatorRemoved deletes the address-pubkey relation when a validator is removed,
func (k Keeper) AfterValidatorRemoved(ctx sdk.Context, address sdk.ConsAddress) {
	k.deleteAddrPubkeyRelation(ctx, crypto.Address(address))
}

//_________________________________________________________________________________________

// Hooks wrapper struct for slashing keeper
type Hooks struct {
	k Keeper
}

var _ types.StakingHooks = Hooks{}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// Implements sdk.ValidatorHooks
func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, _ sdk.ValAddress) {
	h.k.AfterValidatorRemoved(ctx, consAddr)
}

// Implements sdk.ValidatorHooks
func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.k.AfterValidatorCreated(ctx, valAddr)
}

func (h Hooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) {}
