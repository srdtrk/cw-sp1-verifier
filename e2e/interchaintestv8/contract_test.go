package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/srdtrk/cw-sp1-verifier/v8/e2esuite"
	"github.com/srdtrk/cw-sp1-verifier/v8/types"
	"github.com/srdtrk/cw-sp1-verifier/v8/types/cwsp1verifier"
)

// ContactTestSuite is a suite of tests that wraps the TestSuite
// and can provide additional functionality
type ContractTestSuite struct {
	e2esuite.TestSuite

	// this line is used by go-codegen # suite/contract
}

// SetupSuite calls the underlying ContractTestSuite's SetupSuite method
func (s *ContractTestSuite) SetupSuite(ctx context.Context) {
	s.TestSuite.SetupSuite(ctx)
}

// TestWithContractTestSuite is the boilerplate code that allows the test suite to be run
func TestWithContractTestSuite(t *testing.T) {
	suite.Run(t, new(ContractTestSuite))
}

// TestContract is an example test function that will be run by the test suite
func (s *ContractTestSuite) TestPlonkVerifier() {
	ctx := context.Background()

	s.SetupSuite(ctx)

	wasmd1 := s.ChainA

	// Add your test code here. For example, upload and instantiate a contract:
	// This boilerplate may be moved to SetupSuite if it is common to all tests.
	var contract *cwsp1verifier.Contract
	s.Require().True(s.Run("UploadAndInstantiateContract", func() {
		// Upload the contract code to the chain.
		codeID, err := wasmd1.StoreContract(ctx, s.UserA.KeyName(), "../../artifacts/cw_sp1_verifier-plonk.wasm")
		s.Require().NoError(err)

		// Instantiate the contract using contract helpers.
		// This will an error if the instantiate message is invalid.
		contract, err = cwsp1verifier.Instantiate(ctx, s.UserA.KeyName(), codeID, "", wasmd1, cwsp1verifier.InstantiateMsg{})
		s.Require().NoError(err)

		s.Require().NotEmpty(contract.Address)
	}))

	s.Require().True(s.Run("VerifyPlonk", func() {
		fixture := types.GetPlonkFixture()

		queryMsg := &cwsp1verifier.QueryMsg_VerifyProof{
			Proof:        cwsp1verifier.ToBinary(fixture.DecodedProof()),
			PublicValues: cwsp1verifier.ToBinary(fixture.DecodedPublicValues()),
			VkHash:       fixture.Vkey,
		}

		// Call the contract with the proof and public values.
		_, err := contract.QueryClient().VerifyProof(ctx, queryMsg)
		s.Require().NoError(err)
	}))
}
