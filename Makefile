createdb:
	docker exec -it phone-book-db-1 createdb --username=postgres --owner=postgres postgres

dropdb:
	docker exec -it phone-book-db-1 dropdb postgres

migrateup:
	migrate -path src/db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path src/db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: createdb dropdb migrateup migratedown