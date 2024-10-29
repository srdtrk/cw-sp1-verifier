//! # Messages
//!
//! This module defines the messages that this contract receives.

use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::{Binary, Empty};

/// The message to instantiate the contract.
#[cw_serde]
pub struct InstantiateMsg {}

/// The execute messages supported by the contract.
#[cw_serde]
pub enum ExecuteMsg {}

/// The query messages supported by the contract.
#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    /// Verifies an SP1 proof.
    #[returns(Empty)]
    VerifyProof {
        /// The proof to verify.
        proof: Binary,
        /// The public values to verify the proof against.
        public_values: Binary,
        /// The verification key of the sp1 program.
        vk: Binary,
    },
}
