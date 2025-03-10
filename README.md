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

## Configuration

Create a `.env` file in the root directory based on the example below:

```
# Database settings when using docker
WEB_SERVER_PORT=8080
MYSQL_ROOT_PASSWORD=secretpassword
MYSQL_DATABASE=app
MYSQL_USER=user_app
MYSQL_PASSWORD=password_app
MYSQL_EXTERNAL_PORT=3316
MYSQL_PORT=3306
MYSQL_HOST=mysqldb

# Database settings when running outside Docker
#MYSQL_PORT=3316
#MYSQL_HOST=127.0.0.1

JWT_SECRET=secret
```

Modify the values according to your environment.

## Running the Application

### Using Docker

1. Make sure Docker and Docker Compose are installed
2. Adjust the Docker-specific database settings in your `.env` file:
    - Set `MYSQL_PORT=3306`
    - Set `MYSQL_HOST=mysqldb`
3. Build and start the containers:

```bash
docker-compose up -d
```

### IMPORTANT
If you update the .env to run inside the containers you need to rebuild the containers:

```bash
docker-compose up -d --build
```

The API will be available at `http://localhost:8080/api/v1`

### Running Locally

1. Make sure MySQL is running locally or accessible
2. Use the outside-Docker settings in your `.env` file:
    - Set `MYSQL_PORT=3316`
    - Set `MYSQL_HOST=127.0.0.1` or `localhost`
3. Start the application:

```bash
go run main.go
```

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

The API will be available at `http://localhost:8080/api/v1`

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