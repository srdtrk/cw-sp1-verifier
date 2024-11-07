//! # Messages
//!
//! This module defines the messages that this contract receives.

use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::Binary;

/// The message to instantiate the contract.
#[cw_serde]
pub struct InstantiateMsg {}

/// The execute messages supported by the contract.
#[cw_serde]
pub enum ExecuteMsg {
    /// Verifies an SP1 proof.
    VerifyProof(VerifyProofMsg),
}

/// The query messages supported by the contract.
#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    /// Verifies an SP1 proof.
    #[returns(cosmwasm_std::Empty)]
    VerifyProof(VerifyProofMsg),
}

/// Verifies an SP1 proof.
#[cw_serde]
pub struct VerifyProofMsg {
    /// The proof to verify.
    pub proof: Binary,
    /// The public values to verify the proof against.
    pub public_values: Binary,
    /// The verification key of the sp1 program.
    pub vk_hash: String,
}

impl From<VerifyProofMsg> for QueryMsg {
    fn from(msg: VerifyProofMsg) -> Self {
        Self::VerifyProof(msg)
    }
}

impl From<VerifyProofMsg> for ExecuteMsg {
    fn from(msg: VerifyProofMsg) -> Self {
        Self::VerifyProof(msg)
    }
}
