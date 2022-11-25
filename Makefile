run:
	go run app/cmd/main.go

swag:
	swag init -g app/api/router.go -o app/api/docs