{ pkgs ? import <nixpkgs> {}, lib ? pkgs.lib, stdenv ? pkgs.stdenv }:
let
  unstable = import
    (builtins.fetchTarball https://github.com/nixos/nixpkgs/tarball/ccc0c2126893dd20963580b6478d1a10a4512185)
    # reuse the current configuration
    { config = pkgs.config; };
in
  pkgs.mkShell {
    nativeBuildInputs = with pkgs.buildPackages; [
      openssl openssl.dev pkg-config
      just unstable.golangci-lint unstable.go rustup
    ];
    # Run a command after entering the shell
    shellHook = ''
      echo "Entering shell with stable rust"
      rustup toolchain install stable
      # Check if tool is installed
      if [ -z "$(which cargo-prove)" ]; then
        echo "SP1 toolchain is not installed. Please follow the instructions at"
        echo "https://docs.succinct.xyz/getting-started/install.html"
      else
        echo "SP1 toolchain is already installed."
      fi
    '';
}
