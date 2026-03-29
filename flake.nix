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
        version = pkgs.lib.fileContents ./internal/index/version.txt;
      in
      {
        packages = {
          scip-go = pkgs.buildGoModule {
            pname = "scip-go";
            inherit version;
            src = ./.;
            vendorHash = "sha256-DcK4ifLt5884X6xwdVbZdarrOJGa/tslzSXIPuPKD3Q=";
            subPackages = [ "cmd/scip-go" ];
            env.CGO_ENABLED = 0;
            checkPhase = "go test ./...";
            # Use proxyVendor so deps go into the module cache instead
            # of a vendor/ directory that conflicts with test sub-modules
            # having their own go.mod.
            proxyVendor = true;
            nativeCheckInputs = [ pkgs.git ];
          };
          default = self.packages.${system}.scip-go;

          docker = pkgs.dockerTools.buildLayeredImage {
            name = "scip-go";
            tag = version;
            contents = [
              self.packages.${system}.scip-go
              pkgs.go
              pkgs.git
              pkgs.cacert
            ];
            config = {
              Cmd = [ "scip-go" ];
              Env = [ "GOTOOLCHAIN=auto" ];
            };
          };
        };

        checks = {
          nixfmt = pkgs.runCommand "check-nixfmt" { } ''
            ${pkgs.nixfmt}/bin/nixfmt --check ${./flake.nix}
            touch $out
          '';
          goimports = pkgs.runCommand "check-goimports" { } ''
            cd ${./.}
            bad=$(
              # Snapshot outputs are generated with modified
              # indentation for alignment with annotations.
              find . -name '*.go' \
                -not -path '*/testdata/snapshots/output/*' \
                -exec ${pkgs.gotools}/bin/goimports -l {} +
            )
            if [ -n "$bad" ]; then
              echo "goimports check failed on:"
              echo "$bad"
              exit 1
            fi
            touch $out
          '';
        };

        devShells = {
          default = pkgs.mkShellNoCC {
            buildInputs = [
              pkgs.go
              pkgs.gotools
              pkgs.nixfmt
            ];
          };
        };
      }
    );
}
