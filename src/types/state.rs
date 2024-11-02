//! This module defines the state storage of the Contract.

/// Plonk verification key for SP1
#[cfg(feature = "plonk")]
pub const PLONK_VK_BYTES: &[u8] = include_bytes!("../../circuit-artifacts/plonk_vk.bin");
/// Groth16 verification key for SP1
#[cfg(feature = "groth16")]
pub const GROTH16_VK_BYTES: &[u8] = include_bytes!("../../circuit-artifacts/groth16_vk.bin");
