package gov_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/TsukiCore/tsuki/app"
	"github.com/TsukiCore/tsuki/simapp"
	"github.com/TsukiCore/tsuki/x/gov"
	"github.com/TsukiCore/tsuki/x/gov/types"
)

func TestMain(m *testing.M) {
	app.SetConfig()
	m.Run()
}

// When a network actor has not been saved before, it creates one with default params
// and sets the permissions.
func TestHandler_MsgWhitelistPermissions_ActorDoesNotExist(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name             string
		msg              sdk.Msg
		checkWhitelisted bool
	}{
		{
			"Msg Whitelist Permissions",
			&types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			true,
		},
		{
			"Msg Blacklist Permissions",
			&types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			err := setSetPermissionsToAddr(t, app, ctx, proposerAddr)
			require.NoError(t, err)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.NoError(t, err)

			actor, err := app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
			require.NoError(t, err)

			if tt.checkWhitelisted {
				require.True(t, actor.Permissions.IsWhitelisted(types.PermClaimValidator))
			} else {
				require.True(t, actor.Permissions.IsBlacklisted(types.PermClaimValidator))
			}
		})
	}
}

// When a network actor has already permissions it just appends the permission.
func TestNewHandler_SetPermissions_ActorWithPerms(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name             string
		msg              sdk.Msg
		checkWhitelisted bool
	}{
		{
			name: "actor with Whitelist Permissions",
			msg: &types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelisted: true,
		},
		{
			name: "actor with Blacklist Permissions",
			msg: &types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelisted: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			err := setSetPermissionsToAddr(t, app, ctx, proposerAddr)
			require.NoError(t, err)

			// Add some perms before to the actor.
			actor := types.NewDefaultActor(addr)
			if tt.checkWhitelisted {
				err = actor.Permissions.AddToWhitelist(types.PermSetPermissions)
			} else {
				err = actor.Permissions.AddToBlacklist(types.PermSetPermissions)
			}
			require.NoError(t, err)

			app.CustomGovKeeper.SaveNetworkActor(ctx, actor)

			// Call the handler to add some permissions.
			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.NoError(t, err)

			actor, err = app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
			require.NoError(t, err)

			if tt.checkWhitelisted {
				require.True(t, actor.Permissions.IsWhitelisted(types.PermClaimValidator))
				require.True(t, actor.Permissions.IsWhitelisted(types.PermSetPermissions)) // This permission was already set before callid add permission.
			} else {
				require.True(t, actor.Permissions.IsBlacklisted(types.PermClaimValidator))
				require.True(t, actor.Permissions.IsBlacklisted(types.PermSetPermissions)) // This permission was already set before callid add permission.
			}
		})
	}
}

func TestNewHandler_SetPermissionsWithoutSetPermissions(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  sdk.Msg
	}{
		{
			name: "MsgWhitelist",
			msg: &types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
		},
		{
			name: "MsgBlacklist",
			msg: &types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.EqualError(t, err, "PermSetPermissions: not enough permissions")
		})
	}
}

func TestNewHandler_SetPermissions_ProposerHasRoleSudo(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name           string
		msg            sdk.Msg
		checkWhitelist bool
	}{
		{
			name: "MsgWhitelist",
			msg: &types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelist: true,
		},
		{
			name: "MsgBlacklist",
			msg: &types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelist: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			// First we set Role Sudo to proposer Actor
			proposerActor := types.NewDefaultActor(proposerAddr)
			proposerActor.SetRole(types.RoleSudo)
			require.NoError(t, err)
			app.CustomGovKeeper.SaveNetworkActor(ctx, proposerActor)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.NoError(t, err)

			actor, err := app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
			require.NoError(t, err)

			if tt.checkWhitelist {
				require.True(t, actor.Permissions.IsWhitelisted(types.PermClaimValidator))
			} else {
				require.True(t, actor.Permissions.IsBlacklisted(types.PermClaimValidator))
			}
		})
	}
}

func setSetPermissionsToAddr(t *testing.T, app *simapp.SimApp, ctx sdk.Context, addr sdk.AccAddress) error {
	proposerActor := types.NewDefaultActor(addr)
	err := proposerActor.Permissions.AddToWhitelist(types.PermSetPermissions)
	require.NoError(t, err)

	app.CustomGovKeeper.SaveNetworkActor(ctx, proposerActor)

	return nil
}
