postgres:
	docker run -dit -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret --name=db postgres

createdb:
	docker exec -it db createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it db dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/darbitman/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server make
