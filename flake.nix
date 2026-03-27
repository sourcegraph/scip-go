{
  inputs = {
    flake-utils = {
      url = "github:numtide/flake-utils";
    };
    nixpkgs = {
      url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    };
  };

  outputs =
    {
      self,
      flake-utils,
      nixpkgs,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages = {
          scip-go = pkgs.buildGoModule {
            pname = "scip-go";
            version = "0.1.26";
            src = ./.;
            vendorHash = "sha256-AqJ9tVDlSMiT/uPI0K0OliE2mTsFl6bwp1fS7w+PfLU=";
            subPackages = [ "cmd/scip-go" ];
            checkPhase = "go test ./...";
            # Use proxyVendor so deps go into the module cache instead
            # of a vendor/ directory that conflicts with test sub-modules
            # having their own go.mod.
            proxyVendor = true;
            nativeCheckInputs = [ pkgs.git ];
          };
          default = self.packages.${system}.scip-go;
        };

        devShells = {
          default = pkgs.mkShellNoCC {
            buildInputs = [ pkgs.go ];
          };
        };
      }
    );
}
