# Planck.lol

## Development

We are using Nix to manage the `migrate` CLI tool used for DB migrations.

#### Requirements:
- Nix
- direnv (optional)

>[!NOTE]
>
>If using `direnv`, run `direnv allow` to whitelist the current directory.

### Activate Shell

Use `nix develop` to activate the shell within this repo if you're not already using `direnv`. Run `migrate` to make
sure it installed correctly.
