# go-galileosky-test

This is a simple command-line ToDo list application written in Go. It allows you to manage your tasks efficiently.

## Features

- Add a new task with a title and description.
- Mark tasks as completed.
- List all tasks.
- Delete a task.

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/doxanocap/go-galileosky-test.git
   ```

## Run using docker:

1. Compose up:
   ```shell
   docker-compose up
   ```
2. Open docs:
   ```shell
   http://localhost:8080/doc/index.html
   ```

## Run locally

1. Change directory:

   ```shell
   cd go-galileosky-test
   ```

2. Create .env file:
   ```shell
   echo "" >> .env
   
   PORT=8080
   ENV_MODE=development 
   ZAP_JSON=false
    
   PSQL_HOST=
   PSQL_PORT=
   PSQL_USER=
   PSQL_PASSWORD=
   PSQL_DB=gotodo
   PSQL_SSL=disable
   ```
3. Run migrations:
   ```shell
   migrate -path migrations -database "postgres://localhost:5432/db_name?sslmode=disable" up
   ```
4. Run application:
   ```shell
   make run
   ```
5. Run tests:
   ```shell
   make test
   ```