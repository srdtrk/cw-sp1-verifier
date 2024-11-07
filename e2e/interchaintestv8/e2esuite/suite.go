package e2esuite

import (
	"context"

	dockerclient "github.com/docker/docker/client"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	sdkmath "cosmossdk.io/math"

	interchaintest "github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"

	"github.com/srdtrk/cw-sp1-verifier/v8/chainconfig"
	"github.com/srdtrk/cw-sp1-verifier/v8/testvalues"
)

// TestSuite is a suite of tests that require two chains and a relayer
type TestSuite struct {
	suite.Suite

	ChainA       *cosmos.CosmosChain
	UserA        ibc.Wallet
	dockerClient *dockerclient.Client
	network      string
	logger       *zap.Logger
}

// SetupSuite sets up the chains, relayer, user accounts, clients, and connections
func (s *TestSuite) SetupSuite(ctx context.Context) {
	chainSpecs := chainconfig.DefaultChainSpecs

	if len(chainSpecs) != 1 {
		panic("TestSuite requires exactly 1 chain specs")
	}

	t := s.T()

	s.logger = zaptest.NewLogger(t)
	s.dockerClient, s.network = interchaintest.DockerSetup(t)

	cf := interchaintest.NewBuiltinChainFactory(s.logger, chainSpecs)

	chains, err := cf.Chains(t.Name())
	s.Require().NoError(err)
	s.ChainA = chains[0].(*cosmos.CosmosChain)

	ic := interchaintest.NewInterchain().AddChain(s.ChainA)
	s.Require().NoError(ic.Build(ctx, nil, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           s.dockerClient,
		NetworkID:        s.network,
		SkipPathCreation: true,
	}))

	// map all query request types to their gRPC method paths
	s.Require().NoError(populateQueryReqToPath(ctx, s.ChainA))

	// Fund a user accounts
	userFunds := sdkmath.NewInt(testvalues.StartingTokenAmount)
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), userFunds, s.ChainA)
	s.UserA = users[0]

	t.Cleanup(
		func() {
			// Collect diagnostics
			chains := []string{chainSpecs[0].ChainConfig.Name}
			collect(t, s.dockerClient, false, chains...)
		},
	)
}
