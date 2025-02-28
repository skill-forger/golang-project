# Set env SERVER_ENV=dev before run -g ./server/engine.go
serve:
	go run main.go serve
swag:
	swag init --parseDependency --parseDependencyLevel 3 -g main.go -g handler.go -d ./internal/handler -o ./docs/swagger
migrate:
	env SERVER_ENV=local go run ./... migrate