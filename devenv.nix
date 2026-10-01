{
  pkgs,
  config,
  ...
}: {
  services.postgres.enable = true;
  packages = with pkgs; [
    goose
  ];
}
