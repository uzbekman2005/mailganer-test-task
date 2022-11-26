run_main:
	go run app/cmd/mainapi/main.go

run_cron:
	go run app/cmd/cron_job/main.go

swag:
	swag init -g app/api/router.go -o app/api/docs

migrate_up:
	migrate -path app/migrations/ -database postgres://azizbek:Azizbek@localhost:5432/maildb up

migrate_down:
	migrate -path app/migrations/ -database postgres://azizbek:Azizbek@localhost:5432/maildb down
