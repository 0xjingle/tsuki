package cli

import (
	"context"
	"fmt"
	"testing"

	"github.com/TsukiCore/cosmos-sdk/client/flags"

	"github.com/TsukiCore/cosmos-sdk/client"
	"github.com/TsukiCore/cosmos-sdk/testutil"

	"github.com/TsukiCore/cosmos-sdk/baseapp"
	servertypes "github.com/TsukiCore/cosmos-sdk/server/types"
	"github.com/TsukiCore/cosmos-sdk/store/types"
	"github.com/TsukiCore/tsuki/app"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"

	"github.com/TsukiCore/cosmos-sdk/testutil/network"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	//cfg.Codec = app.MakeEncodingConfig().Marshaler

	cfg.NumValidators = 1

	cfg.AppConstructor = func(val network.Validator) servertypes.Application {
		return app.NewInitApp(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			baseapp.SetPruning(types.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestClaimValidatorSet() {
	val := s.network.Validators[0]

	cmd := GetTxClaimValidatorCmd()
	_, out := testutil.ApplyMockIO(cmd)

	clientCtx := val.ClientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs(
		[]string{
			fmt.Sprintf("--%s=%s", flagMoniker, "Moniker"),
			fmt.Sprintf("--%s=%s", flagWebsite, "Website"),
			fmt.Sprintf("--%s=%s", flagSocial, "Social"),
			fmt.Sprintf("--%s=%s", flagIdentity, "Identity"),
			fmt.Sprintf("--%s=%s", flagComission, "10"),
			fmt.Sprintf("--%s=%s", flagPubKey, val.Address.String()),
			fmt.Sprintf("--%s=%s", flagValKey, val.ValAddress.String()),
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Moniker),
		},
	)

	err := cmd.ExecuteContext(ctx)
	s.Require().NoError(err)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
