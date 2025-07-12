# planck.lol

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

Use `nix develop` to activate the shell within this repo if you're not already using `direnv`. Run `migrate` to test
that everything is setup correctly.

### Run Migrations

```sh
migrate -source file://migrations -database postgres://postgres:postgres@localhost:5432/postgres\?sslmode=disable -verbose up
```

### Run in Kubernetes

To run in a k8s environment, we use k3d, Tilt and helm provided by Nix.

#### Setup Redis + PostgreSQL

```sh
$ helm install planck-lol-db bitnami/postgresql \
  --set auth.postgresPassword=secretpassword
$ helm install planck-lol-redis bitnami/redis \
  --set architecture=standalone \
  --set auth.enabled=false
```

#### Run Service

```sh
$ tilt up
```

## Motivation

There was [a video](https://www.youtube.com/watch?v=xFeWVugaouk) that motivated me to try making a link shortener. I'd
like to continue practicing Go (+ systems design) and this seems like a worthwhile project that won't take too much
time.
