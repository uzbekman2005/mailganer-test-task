run:
	go run app/cmd/main.go

swag:
	swag init -g app/api/router.go -o app/api/docs

migrate_up:
	migrate -path app/migrations/ -database postgres://azizbek:Azizbek@localhost:5432/maildb up

migrate_up:
	migrate -path app/migrations/ -database postgres://azizbek:Azizbek@localhost:5432/maildb down
	