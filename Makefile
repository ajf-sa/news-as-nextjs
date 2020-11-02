build:
	docker build . -t news-ajf-sa
run:
	docker run -p 3000:3000 news-ajf-sa
createdb:
	docker exec -it db01 createdb --username=admin --owner=admin news-ajf-sa-db
dropdb:
	docker exec db01 dropdb news-ajf-sa-db -U admin
migration:
	@read -p "Enter migration name " name; \
	migrate create -ext sql -dir app/db/migration -seq $$name

migrate:
	migrate -path app/db/migration -database "postgresql://admin:secret@localhost:5432/news-ajf-sa-db?sslmode=disable" -verbose up
rollback:
	migrate -path app/db/migration -database "postgresql://admin:secret@localhost:5432/news-ajf-sa-db?sslmode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY:	build run createdb dropdb migration migrate rollback sqlc