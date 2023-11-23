generate-proto:
	rm -f pb/proto/*.pb.go
	protoc --go_out=pb/ --go_opt=paths=source_relative \
                                        --go-grpc_out=pb/ --go-grpc_opt=paths=source_relative \
                                        proto/*.proto
db-init:
	psql -c 'CREATE DATABASE "user-management-service"' -U $(user)
migrationup:
    migrate -path db/migrations -database "postgres://$(user):$(password)@$(host):$(port)/user-management-service?sslmode=disable" -verbose up
migrationdown:
    migrate -path db/migrations -database "postgres://$(user):$(password)@$(host):$(port)/user-management-service?sslmode=disable" -verbose down
