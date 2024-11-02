//! This module defines [`ContractError`].

use cosmwasm_std::StdError;
use thiserror::Error;

/// `ContractError` is the error type returned by contract's functions.
#[allow(missing_docs)]
#[allow(clippy::module_name_repetitions)]
#[derive(Error, Debug)]
pub enum ContractError {
    #[error("{0}")]
    Std(#[from] StdError),

    #[error("Unauthorized")]
    Unauthorized {},

    /// Error returned when groth16 verification fails.
    #[error("{0}")]
    Groth16Error(String),

    /// Error returned when plonk verification fails.
    #[error("{0}")]
    PlonkError(String),
}
