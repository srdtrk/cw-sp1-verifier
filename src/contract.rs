//! This module handles the execution logic of the contract.

use cosmwasm_std::{Binary, Deps, DepsMut, Env, MessageInfo, Response};

use crate::types::ContractError;
use crate::types::{
    keys,
    msg::{ExecuteMsg, InstantiateMsg, QueryMsg},
};

#[cfg(all(feature = "plonk", feature = "groth16", feature = "export"))]
compile_error!("This contract cannot be built with both `plonk` and `groth16` features enabled.");

#[cfg(all(not(feature = "plonk"), not(feature = "groth16"), feature = "export"))]
compile_error!(
    "This contract cannot be built without either `plonk` or `groth16` feature enabled."
);

/// `CONTRACT_NAME` is the name of the contract recorded with [`cw2`]
#[cfg(feature = "groth16")]
pub const CONTRACT_NAME: &str = "crates.io:cw-sp1-verifier-groth16";

/// `CONTRACT_NAME` is the name of the contract recorded with [`cw2`]
#[cfg(feature = "plonk")]
pub const CONTRACT_NAME: &str = "crates.io:cw-sp1-verifier-plonk";

/// Instantiates a new contract.
///
/// # Errors
/// Will return an error if the instantiation fails.
#[allow(clippy::needless_pass_by_value)]
#[cosmwasm_std::entry_point]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InstantiateMsg,
) -> Result<Response, ContractError> {
    cw2::set_contract_version(deps.storage, CONTRACT_NAME, keys::CONTRACT_VERSION)?;
    Ok(Response::default())
}

/// Handles the execution of the contract by routing the messages to the respective handlers.
///
/// # Errors
/// Will return an error if the handler returns an error.
#[allow(clippy::needless_pass_by_value)]
#[cosmwasm_std::entry_point]
pub fn execute(
    _deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: ExecuteMsg,
) -> Result<Response, ContractError> {
    unimplemented!()
}

/// Handles the query messages by routing them to the respective handlers.
///
/// # Errors
/// Will return an error if the handler returns an error.
#[allow(clippy::needless_pass_by_value)]
#[cosmwasm_std::entry_point]
pub fn query(_deps: Deps, _env: Env, msg: QueryMsg) -> Result<Binary, ContractError> {
    match msg {
        QueryMsg::VerifyProof {
            proof,
            public_values,
            vk_hash,
        } => query::verify_proof(proof, public_values, vk_hash).map(|_| b"{}".into()),
    }
}

mod query {
    use cosmwasm_std::Empty;

    use super::{Binary, ContractError};

    /// Verifies an SP1 proof.
    ///
    /// # Errors
    /// Will return an error if the proof verification fails.
    /// # Panics
    /// Will panic if the proof or vk cannot be read.
    #[allow(clippy::needless_pass_by_value)]
    pub fn verify_proof(
        proof: Binary,
        public_values: Binary,
        vk_hash: String,
    ) -> Result<Empty, ContractError> {
        #[cfg(feature = "groth16")]
        sp1_verifier::Groth16Verifier::verify(
            proof.as_slice(),
            public_values.as_slice(),
            &vk_hash,
            &sp1_verifier::GROTH16_VK_BYTES,
        )?;

        #[cfg(feature = "plonk")]
        sp1_verifier::PlonkVerifier::verify(
            proof.as_slice(),
            public_values.as_slice(),
            &vk_hash,
            &sp1_verifier::PLONK_VK_BYTES,
        )?;

        Ok(Empty::default())
    }
}

#[cfg(test)]
mod tests {}
