package keeper_test

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"

	types2 "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/tsuki/simapp"
	"github.com/TsukiCore/tsuki/x/staking/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestKeeper_AddValidator(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, abci.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, types2.TokensFromConsensusPower(10))
	addr1 := addrs[0]
	valAddr := types2.ValAddress(addr1)

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

	app.CustomStakingKeeper.AddValidator(ctx, validator)
	getValidator := app.CustomStakingKeeper.GetValidator(ctx, validator.ValKey)

	assert.Equal(t, validator, getValidator)
}

func TestKeeper_GetValidatorSet(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, abci.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 2, types2.TokensFromConsensusPower(10))
	addr1 := addrs[0]
	valAddr1 := types2.ValAddress(addr1)

	addr2 := addrs[1]
	valAddr2 := types2.ValAddress(addr2)

	fmt.Printf("val addrs %s\n", addr1.String())
	fmt.Printf("val addrs %s\n", valAddr1.String())

	validator1, err := types.NewValidator(
		"validator 1",
		"some-web.com",
		"A Social",
		"My Identity",
		types2.NewDec(1234),
		valAddr1,
		addr1,
	)
	require.NoError(t, err)

	validator2, err := types.NewValidator(
		"validator 2",
		"some-web.com",
		"A Social",
		"My Identity",
		types2.NewDec(1234),
		valAddr2,
		addr2,
	)
	require.NoError(t, err)

	app.CustomStakingKeeper.AddValidator(ctx, validator1)
	app.CustomStakingKeeper.AddValidator(ctx, validator2)

	validatorSet := app.CustomStakingKeeper.GetValidatorSet(ctx)
	require.Equal(t, 2, len(validatorSet))
}
