# Create database service using helm
db-create:
	helm install planck-lol-db bitnami/postgresql \
		--set auth.postgresPassword=secretpassword

# Migrate database
db-migrate:
	migrate -source file://migrations \
		-database postgres://postgres:secretpassword@localhost:5432/postgres\?sslmode=disable \
		-verbose up

# Expose postgresql service to localhost:5432
db-expose:
	kubectl port-forward \
		svc/planck-lol-db-postgresql 5432:5432

# Create Redis service using helm
redis-create:
	helm install planck-lol-redis bitnami/redis \
		--set architecture=standalone \
		--set auth.enabled=false

# Expose redis service to localhost:6379
redis-expose:
	kubectl port-forward \
		svc/planck-lol-redis 6379:6379
