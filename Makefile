# Set env SERVER_ENV=dev before run
run:
	env SERVER_ENV=dev go run ./... server
swag:
	swag init --parseDependency --parseDependencyLevel 3 -g main.go -d ./cmd,./servers,./v1/users -o ./docs
migrate:
	env SERVER_ENV=dev go run ./... migrate