package app

import (
	"encoding/json"

	customstaking "github.com/TsukiCore/tsuki/x/staking"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// ExportAppStateAndValidators export the state of Tsuki for a genesis file
func (app *TsukiApp) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string,
) (servertypes.ExportedApp, error) {

	ctx := app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()})

	// TODO: handle zero height upgrades
	height := app.LastBlockHeight() + 1

	genState := app.mm.ExportGenesis(ctx, app.appCodec)
	appState, err := json.MarshalIndent(genState, "", "  ")
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	validators, err := customstaking.WriteValidators(ctx, app.customStakingKeeper)
	return servertypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: app.BaseApp.GetConsensusParams(ctx),
	}, err
}
