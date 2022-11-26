run:
	go run cmd/main.go

swag:
	swag init -g api/router.go -o api/docs

migrate_up:
	migrate -path migrations/ -database postgres://azizbek:Azizbek@localhost:5432/maildb up

migrate_down:
	migrate -path migrations/ -database postgres://azizbek:Azizbek@localhost:5432/maildb down
