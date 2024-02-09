migrate-up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank" -verbose up