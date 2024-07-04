# Api Rest Go

## Description
This api is my frist api rest in go, it's a simple api that allows you to create, read, update and delete personalities.
is made it for study purposes.

## Installation
You need to have go installed in your machine, you can download it [here](https://golang.org/dl/) chose the 1.23^ Version.
After that you can clone this repository and run the command `go run main.go` in the root of the project.

for database i'm using a postgres database, in docker container, you can run the command `docker-compose up` in the root of the project to start the database.
obs: you nedd to have docker and docker-compose installed in your machine.

## Endpoints
- GET /personalities
- GET /personalities/{id}
- POST /personalities
- PUT /personalities/{id}
- DELETE /personalities/{id}

## Technologies
- Go
- Gorm
- Gorilla Mux
- Postgres
- Docker
- Docker-compose

## Final considerations
This project was made for study purposes, if you have any questions or suggestions, please contact me.


