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

        mkScipGo =
          {
            GOOS ? null,
            GOARCH ? null,
            doCheck ? true,
          }:
          pkgs.buildGoModule {
            pname = "scip-go";
            inherit version;
            src = ./.;
            vendorHash = "sha256-AqJ9tVDlSMiT/uPI0K0OliE2mTsFl6bwp1fS7w+PfLU=";
            subPackages = [ "cmd/scip-go" ];
            env = {
              CGO_ENABLED = 0;
            } // pkgs.lib.optionalAttrs (GOOS != null) { inherit GOOS; }
              // pkgs.lib.optionalAttrs (GOARCH != null) { inherit GOARCH; };
            inherit doCheck;
            checkPhase = "go test ./...";
            # Use proxyVendor so deps go into the module cache instead
            # of a vendor/ directory that conflicts with test sub-modules
            # having their own go.mod.
            proxyVendor = true;
            nativeCheckInputs = [ pkgs.git ];
          };
      in
      {
        packages = {
          scip-go = mkScipGo { };
          default = self.packages.${system}.scip-go;

          scip-go-linux-amd64 = mkScipGo {
            GOOS = "linux";
            GOARCH = "amd64";
            doCheck = false;
          };
          scip-go-linux-arm64 = mkScipGo {
            GOOS = "linux";
            GOARCH = "arm64";
            doCheck = false;
          };
          scip-go-darwin-amd64 = mkScipGo {
            GOOS = "darwin";
            GOARCH = "amd64";
            doCheck = false;
          };
          scip-go-darwin-arm64 = mkScipGo {
            GOOS = "darwin";
            GOARCH = "arm64";
            doCheck = false;
          };
          scip-go-windows-amd64 = mkScipGo {
            GOOS = "windows";
            GOARCH = "amd64";
            doCheck = false;
          };
          scip-go-windows-arm64 = mkScipGo {
            GOOS = "windows";
            GOARCH = "arm64";
            doCheck = false;
          };

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

        devShells = {
          default = pkgs.mkShellNoCC {
            buildInputs = [ pkgs.go ];
          };
        };
      }
    );
}
