# `CosmWasm` SP1 Verifier Contract

This is a `CosmWasm` contract that verifies [SP1](https://github.com/succinctlabs/sp1) proofs. It can be used to verify plonk proofs if built with the `plonk` feature, or groth16 proofs if built with the `groth16` feature.

This contract uses [cosmwasm optimizer](https://github.com/CosmWasm/optimizer) to build both the groth16 and plonk versions of the contract. The optimizer is a tool that allows you to build multiple versions of a contract with different features enabled.

To build the contracts, run the following command:

```sh
just build-optimize
```

This will build the contracts under the `artifacts` directory.

## Usage

This contract is meant to be queried by other contracts/clients to verify SP1 proofs. As such, it only has one query message, `VerifyProof`, which takes the following form:

```rust
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
```
