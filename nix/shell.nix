{ system ? builtins.currentSystem, sources ? import ./sources.nix }:
let
  pkgs = import sources.nixpkgs { 
    inherit system;

    config = { };

    overlays = [
      (self: super: 
        let
          sources = super.callPackages ./_sources/generated.nix { };
        in
        {
          nvfetcher = super.nvfetcher.overrideAttrs (old: {
            inherit (sources.nvfetcher) src version;
          });
          golangci-lint = super.golangci-lint.overrideAttrs (old: {
            inherit (sources.golangci-lint) src version;
          });
          frugal = super.frugal.overrideAttrs (old: {
            inherit (sources.frugal) src version;
          });
          # TODO: Add edgedb-cli here
        }
      )
    ];
  };
  src = fetchTarball "https://github.com/numtide/devshell/archive/master.tar.gz";
  devshell = import src { 
    inherit pkgs system;
  };
in
devshell.fromTOML ./devshell.toml
