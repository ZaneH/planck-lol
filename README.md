# Planck.lol

## Development

We are using Nix to provide Go and the [migrate](github.com/golang-migrate/migrate) CLI tool used for DB migrations.

#### Requirements:
- Nix
- direnv (optional)

>[!NOTE]
>
>If using `direnv`, run `direnv allow` to whitelist the current directory.

### Check Your Environment

Use `nix develop` to activate the shell within this repo if you're not already using `direnv`. Run `migrate` to make
sure everything is setup correctly.

### Run Migrations

```sh
migrate -source file://migrations -database postgres://postgres:postgres@localhost:5432/postgres\?sslmode=disable -verbose up
```
