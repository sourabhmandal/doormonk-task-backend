migrate-up:
	migrate -path database/migrations -database "sqlite3://database/doormonk" -verbose up

migrate-down:
	migrate -path database/migrations -database "sqlite3://database/doormonk" -verbose up

.PHONY: migrate-up migrate-down