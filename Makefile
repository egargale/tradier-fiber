migrate_local:
	migrate -verbose -path ./db/migrations \
		-database postgres://root:secret@localhost:5432/tradier?sslmode=disable up

rollback_local:
	migrate -verbose -path ./db/migrations \
		-database postgres://root:secret@localhost:5432/tradier?sslmode=disable down

drop_local:
	migrate -verbose -path ./db/migrations \
		-database postgres://root:secret@localhost:5432/tradier?sslmode=disable drop
		
migrate:
	docker-compose run --rm pg_helpers migrate -verbose -path ./migrations \
		-database postgres://root:secret@postgres:5432/tradier?sslmode=disable up

rollback:
	docker-compose run --rm pg_helpers migrate -verbose -path ./migrations \
		-database postgres://root:secret@postgres:5432/tradier?sslmode=disable down

drop:
	docker-compose run --rm pg_helpers migrate -verbose -path ./migrations \
		-database postgres://root:secret@postgres:5432/tradier?sslmode=disable drop

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir db/migrations $$name

install_schema:
	migrate create -ext sql -dir db/migrations -seq init_schema

schema:
	dbml2sql ./db/seeds/schema.dbml -o ./db/seeds/schema.sql

data:
	cat db/seeds/data.sql | docker-compose exec -T postgres psql  postgres://root:secret@postgres:5432/tradier?sslmode=disable

test:
	docker-compose run --rm pg_helpers PGPASSWORD=secret psql -U root -d tradier -p 5432 -h postgres -f ./seeds/data.sql

sqlc:
	sqlc generate