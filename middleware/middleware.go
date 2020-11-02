package middleware

import (
	feeprocessingkeeper "github.com/TsukiCore/tsuki/x/feeprocessing/keeper"
	customgovkeeper "github.com/TsukiCore/tsuki/x/gov/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	customGovKeeper     customgovkeeper.Keeper
	feeprocessingKeeper feeprocessingkeeper.Keeper
)

// SetKeepers set keepers to be used on middlewares
func SetKeepers(cgk customgovkeeper.Keeper, fk feeprocessingkeeper.Keeper) {
	customGovKeeper = cgk
	feeprocessingKeeper = fk
}

// NewRoute returns an instance of Route.
func NewRoute(p string, h sdk.Handler) sdk.Route {
	newHandler := func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		hResult, hErr := h(ctx, msg)
		if hErr != nil {
			return hResult, hErr
		}
		// handle extra fee based on handler result

		fee := customGovKeeper.GetExecutionFee(ctx, msg.Type())
		if fee == nil {
			return hResult, hErr
		}

		feeprocessingKeeper.SetExecutionStatusSuccess(ctx, msg)
		return hResult, hErr
	}
	return sdk.NewRoute(p, newHandler)
}
