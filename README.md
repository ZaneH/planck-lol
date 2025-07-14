# planck.lol

## Development

We are using Nix to provide Go and the [migrate](https://github.com/golang-migrate/migrate) CLI tool used for DB
migrations.

##### Requirements:
- Nix
- direnv (optional)

>[!NOTE]
>
>If using `direnv`, run `direnv allow` to whitelist the current directory.

#### Check Your Environment

If you're not using `direnv`, run `nix develop` to activate the Nix shell in this repo. Then run `migrate`, `k3d` or
`kubectl` to test that the required tools are available.

### Run in Kubernetes

We use k3d + Tilt to manage the local K8s environment. To start the app + infra, run:

```sh
$ tilt up
```

### Run Migrations

```sh
$ make db-migrate
```

## Motivation

There was [a video](https://www.youtube.com/watch?v=xFeWVugaouk) that motivated me to try making a link shortener. I'd
like to continue practicing Go (+ systems design) and this seems like a worthwhile project that won't take too much
time.
