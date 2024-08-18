# A Simple Go Web Service

This is a sample of a simple web service in Golang that accepts a message and stores it to a sqlite database.

It uses the Gin framework in Go which gives it the benefit of being robust and highly concurrent. The project is separated into buisness logic, and a domain layer which also acts as a data access layer, allowing for easy maintainability. Integration tests were created using Testify to ensure reliability of the service. A dockerfile can be used to quickly get the application up and running without the need of compiling to a specific target machine.

If you would like to compile for your specific device, you can use the built go functions

## Setup

### Target Machine

To setup on a specific target machine:

Run `go mod download`

Then `go run .`

### Docker

To setup through docker ensure you have docker setup then build the docker container:

`docker build -t go-message-app .`

Then run the container exposing port 8080

`docker run -d -p 8080:8080 --name go-message-app-container go-message-app`
