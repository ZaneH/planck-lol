# Planck.lol

## Development

We are using Nix to provide Go and the [migrate](https://github.com/golang-migrate/migrate) CLI tool used for DB
migrations.

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

## Motivation

There was [a video](https://www.youtube.com/watch?v=xFeWVugaouk) that motivated me to try making a link shortener. I'd
like to continue practicing Go (+ systems design) and this seems like a worthwhile project that won't take too much
time.
