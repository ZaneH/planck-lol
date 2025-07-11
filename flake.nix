{
  description = "Planck URL shortener, written in Golang";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        migrate = pkgs.buildGoModule rec {
          pname = "migrate";
          version = "v4.18.3";

          src = pkgs.fetchFromGitHub {
            owner = "golang-migrate";
            repo = "migrate";
            rev = version;
            hash = "sha256-aM8okSrLj2oIb3Ey2KkHu3UQY7mSnPjMfwNsdL2Fz28=";
          };

          subPackages = [ "cmd/migrate" ];

          vendorHash = "sha256-yJ1D0uhUQXBjw7ikRmcW8DFn0wvnBCw6jDYl3UQepXU=";

          tags = [ "postgres" ];
        };
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            migrate

            pkgs.tilt
            pkgs.k3d
            pkgs.kubectl
            pkgs.kubernetes-helm
            pkgs.redis
          ];
        };
      });
}
