{
  description = "feeder - go project";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    systems.url = "github:nix-systems/default";
    process-compose-flake.url = "github:Platonic-Systems/process-compose-flake";
    services-flake.url = "github:juspay/services-flake";

    # northwind.url = "github:pthom/northwind_psql";
    # northwind.flake = false;
  };
  outputs = inputs:
    inputs.flake-parts.lib.mkFlake {inherit inputs;} {
      systems = import inputs.systems;
      imports = [
        inputs.process-compose-flake.flakeModule
      ];
      perSystem = {
        self',
        pkgs,
        config,
        ...
      }: {
        # `process-compose.foo` will add a flake package output called "foo".
        # Therefore, this will add a default package that you can build using
        # `nix build` and run using `nix run`.
        process-compose."feeder" = {config, ...}: let
          dbName = "gator";
        in {
          imports = [
            inputs.services-flake.processComposeModules.default
          ];

          services.postgres."pg1" = {
            enable = true;
            # initialDatabase = [
            #   {
            #     name = dbName;
            #   }
            # ];
          };

          settings.processes.pgweb = let
            pgcfg = config.services.postgres.pg1;
          in {
            environment.PGWEB_DATABASE_URL = pgcfg.connectionURI {inherit dbName;};
            command = pkgs.pgweb;
            depends_on."pg1".condition = "process_healthy";
          };
        };
        packages.default = self'.packages.feeder;

        devShells.default = pkgs.mkShell {
          inputsFrom = [
            # Add the packages of the enabled services in the devShell
            #
            # For example: `psql` to interact with `postgres` server or `redis-cli` with `redis-server`
            config.process-compose."feeder".services.outputs.devShell
          ];
          packages = [
            pkgs.goose
            pkgs.sqlc
            # Add the process-compose app in the devShell
            #
            # In the devShell, run `simple-example` to run the app
            self'.packages.feeder
          ];
          # nativeBuildInputs = [pkgs.just pkgs.goose];
        };
      };
    };
}
