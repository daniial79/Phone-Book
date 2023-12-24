createdb:
	docker exec -it phone-book-db-1 createdb --username=postgres --owner=postgres phonebook

dropdb:
	docker exec -it phone-book-db-1 dropdb phonebook

migrateup:
	migrate -path src/db/migrations -database "postgresql://postgres:postgres@localhost:5432/phonebook?sslmode=disable" -verbose up

migratedown:
	migrate -path src/db/migrations -database "postgresql://postgres:postgres@localhost:5432/phonebook?sslmode=disable" -verbose down

.PHONY: createdb dropdb migrateup migratedown