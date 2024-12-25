# Golang Project Layout

## Specification
- Init Project using Golang [Cobra](https://github.com/spf13/cobra)
- Parse and Get Env Configuration using [Viper](https://github.com/spf13/viper)
- Implement Http Server Router [Echo](https://echo.labstack.com/docs)
- Connect to MySQL Data via [Gorm](https://gorm.io/docs/)
- Prepare local testing and start up with [Docker Compose](https://docs.docker.com/manuals/)

# Dependencies
```go install github.com/swaggo/swag/cmd/swag@v1.16.4```

# Install Cobra CLI


# Project Layout
This project uses Monolithic Modular Architecture.

# Swagger
* Play url:
```
http://localhost:8080/swagger/index.html
```
* Generate swagger specification
```bash
make swag
```

# How to start the server
* Add to the environment variable with these:
  * SERVER_ENV=dev

* Run by
```bash
env SERVER_ENV=dev go run cmd/main.go server
```
or
```bash
make run
```