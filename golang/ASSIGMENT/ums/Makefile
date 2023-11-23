db-init:
	psql -c 'CREATE DATABASE "user-management-service"' -U $(user)
migrationup:
    migrate -path db/migrations -database "postgres://$(user):$(password)@$(host):$(port)/user-management-service?sslmode=disable" -verbose up
migrationdown:
    migrate -path db/migrations -database "postgres://$(user):$(password)@$(host):$(port)/user-management-service?sslmode=disable" -verbose down