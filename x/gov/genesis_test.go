package gov_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/TsukiCore/tsuki/simapp"
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov"
	"github.com/TsukiCore/tsuki/x/gov/keeper"
	"github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

// test for gov export / init genesis process
func TestSimappExportGenesis(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	genesis := gov.ExportGenesis(ctx, app.CustomGovKeeper)
	bz, err := app.AppCodec().MarshalJSON(genesis)
	require.NoError(t, err)
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, []byte(`{
		"starting_proposal_id":"1",
		"permissions":{
			"1":{"blacklist":[],"whitelist":[1,2,3,6,8,9,4,5,12,13,10,11,14,15,18,19,20,21,22,23,24,25,16,17]},
			"2":{"blacklist":[],"whitelist":[2]}
		},
		"network_actors":[],
		"network_properties":{
			"min_tx_fee":"100",
			"max_tx_fee":"1000000",
			"vote_quorum":"33",
			"proposal_end_time":"600",
			"proposal_enactment_time":"300",
			"min_proposal_end_blocks":"2",
			"min_proposal_enactment_blocks":"1",
			"enable_foreign_fee_payments":true,
			"mischance_rank_decrease_amount":"10",
			"max_mischance":"110",
			"mischance_confidence":"10",
			"inactive_rank_decrease_percent":"50",
			"min_validators":"1",
			"poor_network_max_bank_send":"1000000",
			"jail_max_time":"600",
			"enable_token_whitelist":false,
			"enable_token_blacklist":true
		},
		"execution_fees":[],
		"poor_network_messages":{
			"messages":["proposal-assign-permission","proposal-set-network-property","set-network-properties","vote-proposal","claim-councilor","whitelist-permissions","blacklist-permissions","create-role","assign-role","remove-role","whitelist-role-permission","blacklist-role-permission","remove-whitelist-role-permission","remove-blacklist-role-permission","claim-validator","activate","pause","unpause"]
		},
		"proposals":[],
		"votes":[],
		"data_registry":{},
		"last_identity_record_id":"0",
		"last_id_record_verify_request_id":"0"
	}`))
	require.NoError(t, err)
	require.Equal(t, string(bz), buffer.String())
}

func TestExportInitGenesis(t *testing.T) {
	// Initialize store keys
	keyGovernance := sdk.NewKVStoreKey(types.ModuleName)

	// Initialize memory database and mount stores on it
	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyGovernance, sdk.StoreTypeDB, nil)
	ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, tmproto.Header{
		Height: 1,
		Time:   time.Date(2020, time.March, 1, 1, 0, 0, 0, time.UTC),
	}, false, log.TestingLogger())

	k := keeper.NewKeeper(keyGovernance, simapp.MakeEncodingConfig().Marshaler, nil)

	genState := types.GenesisState{
		Permissions: map[uint64]*types.Permissions{
			uint64(types.RoleSudo): types.NewPermissions([]types.PermValue{
				types.PermSetPermissions,
				types.PermClaimValidator,
				types.PermClaimCouncilor,
				types.PermUpsertTokenAlias,
			}, nil),
		},
		StartingProposalId: 1,
		NetworkProperties: &types.NetworkProperties{
			MinTxFee:                   100,
			MaxTxFee:                   1000000,
			VoteQuorum:                 33,
			ProposalEndTime:            600, // 600 seconds / 10 mins
			ProposalEnactmentTime:      300, // 300 seconds / 5 mins
			MinProposalEndBlocks:       2,
			MinProposalEnactmentBlocks: 1,
			EnableForeignFeePayments:   true,
		},
		ExecutionFees: []*types.ExecutionFee{
			{
				Name:              "Claim Validator Seat",
				TransactionType:   "claim-validator-seat",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
		},
		PoorNetworkMessages: &types.AllowedMessages{
			Messages: []string{
				tsukitypes.MsgTypeProposalAssignPermission,
				tsukitypes.MsgTypeProposalSetNetworkProperty,
				tsukitypes.MsgTypeSetNetworkProperties,
			},
		},
	}
	err := gov.InitGenesis(ctx, k, genState)
	require.NoError(t, err)

	genesis := gov.ExportGenesis(ctx, k)
	bz, err := simapp.MakeEncodingConfig().Marshaler.MarshalJSON(genesis)
	require.NoError(t, err)
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, []byte(`{
		"starting_proposal_id":"1",
		"permissions":{
			"1":{"blacklist":[],"whitelist":[1,2,3,6]}
		},
		"network_actors":[],
		"network_properties":{
			"min_tx_fee":"100",
			"max_tx_fee":"1000000",
			"vote_quorum":"33",
			"proposal_end_time":"600",
			"proposal_enactment_time":"300",
			"min_proposal_end_blocks":"2",
			"min_proposal_enactment_blocks":"1",
			"enable_foreign_fee_payments":true,
			"mischance_rank_decrease_amount":"0",
			"max_mischance":"0",
			"mischance_confidence":"0",
			"inactive_rank_decrease_percent":"0",
			"min_validators":"0",
			"poor_network_max_bank_send":"0",
			"jail_max_time":"0",
			"enable_token_whitelist":false,
			"enable_token_blacklist":false
		},
		"execution_fees":[],
		"poor_network_messages":{
			"messages":["proposal-assign-permission","proposal-set-network-property","set-network-properties"]
		},
		"proposals":[],
		"votes":[],
		"data_registry":{},
		"last_identity_record_id":"0",
		"last_id_record_verify_request_id":"0"
	}`))
	require.NoError(t, err)
	require.Equal(t, string(bz), buffer.String())
}
