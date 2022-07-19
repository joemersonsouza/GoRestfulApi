# Simple API in Go

# Author: Joemerson Souza
# Version: 1.0
# License BSD

Fell free to change better approaches.

You can use this project as a template to create your own API using the following technologies:

- Go
- Gin
- Gorm
- Postgres
- Swagger

# Testing

This project doesn't use a test framework or library, so the test use the class substitution to override the injection.


# Main llibraries (please take a look on the code to be sure that you have the environment set correctly)
go get github.com/gin-gonic/gin
go get github.com/swaggo/gin-swagger
go get -u github.com/swaggo/swag/cmd/swag

# Database Postgres libraries
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres@v1.9.16

go install github.com/swaggo/swag/cmd/swag@latest ´Really important to run the swag command´

# Initialize or recreate the Swagger file
swag init

# Initialize Docker with Dependencies
docker-compose -f docker/docker-compose.dev.yaml up

# Run the API inside of root path (../api-sample)
go run .

# Run Services Test (../api-sample/services)

go test