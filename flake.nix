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
            vendorHash = "sha256-kBOdLecRiWiBA8xvNHFcDZNKiLMiN3Vww6CNumMLW2M=";
            subPackages = [ "cmd/scip-go" ];
          };
          default = self.packages.${system}.scip-go;
        };

        checks = {
          snapshots = pkgs.buildGoModule {
            pname = "scip-go-snapshot-check";
            version = "0.1.26";
            src = ./.;
            vendorHash = "sha256-AqJ9tVDlSMiT/uPI0K0OliE2mTsFl6bwp1fS7w+PfLU=";
            proxyVendor = true;
            nativeBuildInputs = [ pkgs.git ];
            buildPhase = "true";
            checkPhase = ''
              # The snapshot test uses `git rev-parse --show-toplevel`
              # to locate testdata, so we need a git repo.
              git init -q
              git add -A
              go test ./internal/index -count=1
            '';
            installPhase = "touch $out";
          };
        };

        devShells = {
          default = pkgs.mkShellNoCC {
            buildInputs = [ pkgs.go ];
          };
        };
      }
    );
}
