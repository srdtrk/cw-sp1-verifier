# Build optimized wasm using the cosmwasm/optimizer:0.15.1 docker image
build-optimize:
  @echo "Compiling optimized wasm..."
  docker run --rm -t -v "$(pwd)":/code \
    --mount type=volume,source="$(basename "$(pwd)")_cache",target=/code/target \
    --mount type=volume,source=registry_cache,target=/usr/local/cargo/registry \
    cosmwasm/optimizer:0.16.1

# Run cargo fmt and clippy checks
lint:
  cargo fmt --all -- --check
  cargo clippy --all-targets --all-features -- -D warnings

# Generate JSON schema files for all contracts in the project
generate-schemas:
  @echo "Generating JSON schema files..."
  cargo run --bin schema
  @echo "Done."

# Run the unit tests
unit-tests:
  cargo test --locked --no-default-features --features export,groth16
  cargo test --locked --no-default-features --features export,plonk
