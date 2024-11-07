package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	interchaintest "github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"

	"github.com/srdtrk/cw-sp1-verifier/v8/e2esuite"
)

// BasicTestSuite is a suite of tests that wraps the TestSuite
// and can provide additional functionality
type BasicTestSuite struct {
	e2esuite.TestSuite
}

// SetupSuite calls the underlying BasicTestSuite's SetupSuite method
func (s *BasicTestSuite) SetupSuite(ctx context.Context) {
	s.TestSuite.SetupSuite(ctx)
}

// TestWithBasicTestSuite is the boilerplate code that allows the test suite to be run
func TestWithBasicTestSuite(t *testing.T) {
	suite.Run(t, new(BasicTestSuite))
}

// TestBasic is an example test function that will be run by the test suite
func (s *BasicTestSuite) TestBasic() {
	ctx := context.Background()

	s.SetupSuite(ctx)

	wasmd1 := s.ChainA

	// Add your test code here. For example, create a new wallet and fund it from the UserA account
	var newUser ibc.Wallet
	s.Run("CreateANewAccountAndSend", func() {
		var err error
		newUser, err = interchaintest.GetAndFundTestUserWithMnemonic(ctx, s.T().Name(), "", 1, wasmd1)
		s.Require().NoError(err)

		/*
			Send funds to the new user

			I am using the broadcaster to send funds to the new user to demonstrate how to use the broadcaster,
			but you can also use the s.FundAddressChainA method to send funds to the new user, or
			use the SendFunds method from the s.ChainA instance.

			I used 200_000 gas to send the funds, but you can use any amount of gas you want.
		*/
		_, err = s.BroadcastMessages(ctx, wasmd1, s.UserA, 200_000, &banktypes.MsgSend{
			FromAddress: s.UserA.FormattedAddress(),
			ToAddress:   newUser.FormattedAddress(),
			Amount:      sdk.NewCoins(sdk.NewInt64Coin(wasmd1.Config().Denom, 100_000)),
		})
		s.Require().NoError(err)
	})

	// Test if the send was successful
	s.Run("VerifySendMessage", func() {
		/*
			Query the new user's balance

			I am using the GRPCQuery to query the new user's balance to demonstrate how to use the GRPCQuery,
			but you can also use the GetBalance method from the s.ChainA instance.
		*/
		balanceResp, err := e2esuite.GRPCQuery[banktypes.QueryBalanceResponse](ctx, wasmd1, &banktypes.QueryBalanceRequest{
			Address: newUser.FormattedAddress(),
			Denom:   wasmd1.Config().Denom,
		})
		s.Require().NoError(err)
		s.Require().NotNil(balanceResp.Balance)

		// Verify the balance
		s.Require().Equal(int64(100_001), balanceResp.Balance.Amount.Int64())
	})
}
