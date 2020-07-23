package staking

import (
	types2 "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/tsuki/simapp"
	"github.com/TsukiCore/tsuki/x/staking/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestNewHandler_MsgClaimValidator_HappyPath(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, abci.Header{})

	validator, err := types.NewValidator(
		"aMoniker",
		"some-web.com",
		"A Social",
		"My Identity",
		types2.NewDec(1234),
		valAddr,
		addr1,
	)
	require.NoError(t, err)

}
