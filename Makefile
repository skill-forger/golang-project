# Set env SERVER_ENV=dev before run
run:
	go run main.go serve
swag:
	swag init --parseDependency --parseDependencyLevel 3 -g main.go -d ./cmd,./servers,./v1/users -o ./docs
migrate:
	env SERVER_ENV=local go run ./... migrate