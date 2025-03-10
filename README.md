# Tasks API

A RESTful API for task management with role-based access (technicians and managers).

## Features

- Task management system
- Role-based authorization (Technicians/Managers)
- JWT authentication
- MySQL database storage
- Docker support

## Prerequisites

- Go 1.18+
- MySQL 8.0+
- Docker and Docker Compose (optional)

## Installation

### Clone the repository

```bash
git clone https://github.com/mayckol/tasks-api.git
cd tasks-api
```

### Install dependencies

```bash
go mod download
```

### Running Locally

1. Make sure MySQL and RabbitMQ is running locally or accessible
2. Create the `.env` file according to your chosen configuration, based on `.env.example` or `.env.docker-example`
3. Start the application:

```bash
go run main.go
```

### Using Docker

1. Make sure Docker and Docker Compose are installed
2. Build and start the containers:

```bash
docker-compose up -d --build
```

### IMPORTANT
If you update the .env to run inside the containers you need to rebuild the containers:

```bash
docker-compose up -d --build
```

The API will be available at `http://localhost:8080/api/v1`

## Migration
To create the database schema, run the following command:

```bash
make migrateup
```

To rollback the database schema, run the following command:

```bash
make migratedown
```

## Seed
To seed the database with the initial data, run the following command:

```bash
make seed
```
This will create a manager and a technician user with the following credentials as you can see in the `seed/main.go` file

The API will be available at `http://localhost:8080/api/v1`

## Reading messages from the queue
To read messages from the queue, run the following command:

```bash
make readtasks
```

## API Usage

The API provides endpoints for:

- User authentication (login to get JWT token)
- Task management for technicians
- Administrative task overview for managers

### Authentication

All API endpoints (except login) require a valid JWT token in the Authorization header.

### Role-Based Access

- **Technicians** can manage their own tasks
- **Managers** can view and manage all tasks in the system

## Documentation

API documentation is available via Swagger at `http://localhost:8080/api/v1/docs/index.html` when the application is running.

## Collection
There is a Insomnia collection in the root of the project called `tasks-api-insomnia.json` that you can import into Insomnia to test the API.

## Testing

Run the tests with:

```bash
make test
```

## Credits
- The docker initial setup uses [docker init](https://docs.docker.com/reference/cli/docker/init/)
- The gracefull shutdown idea comes from [go-blueprint](https://github.com/Melkeydev/go-blueprint)
- The messaging broker idea comes from [RabbitMQ](https://www.rabbitmq.com/tutorials/tutorial-one-go)
- The Clean Architecture was inspired by [Full Cycle](https://fullcycle.com.br/clean-architecture-trabalhe-em-aplicacoes-de-grande-porte/)
