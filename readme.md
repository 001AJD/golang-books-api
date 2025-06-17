# Books API

A simple REST API built with Go. This API demonstrates basic CRUD operations on a "Books" resource using [GORM](https://gorm.io) as the ORM and [PostgreSQL](https://www.postgresql.org) for database storage. Routes are handled with [Gorilla Mux](https://github.com/gorilla/mux).

## Project Structure

```
books-api/
├── db/
│   └── connection.go       # Database connection
|── db-migrations/           # Database DDL changesets
|   └──
├── handlers/
│   ├── base-handler.go     # Provides the base handler with a DB instance.
│   └── handlers.go         # Contains the GetBooks HTTP endpoint implementation.
├── models/
│   └── Books.model.go      # Defines the Books model.
├── go.mod                  # Module definition and dependencies.
├── go.sum                  # Dependency checksums.
├── docker-compose.yml      # Docker Compose file to spin up a PostgreSQL container.
└── main.go                 # Entrypoint for the API server.
```

## Requirements

- Go 1.24+
- PostgreSQL (or use the provided Docker Compose/Podman Compose configuration)
- Docker/Docker Compose or Podman/Podman Compose (optional)

## Getting Started

### 1. Running the PostgreSQL Database

You can use the provided Docker Compose file to run a PostgreSQL instance.

#### Using Docker Compose

Open a terminal in the project root and run:

```
docker-compose up -d
```

#### Using Podman Compose

If you prefer Podman, run:

```
podman-compose up -d
```

The PostgreSQL container will be set up with the following credentials:

- **User:** gorm
- **Password:** gorm
- **Database:** books
- **Port:** 5432

### 2. Running the API Server

Run the following command in the project root:

```
go run main.go
```

The server will start and listen on port `3000`. You can access the books API at:

```
GET http://localhost:3000/api/v1/books
```

## API Endpoints

- **GET /api/v1/books**  
  Retrieves all books from the database.

## Database Migration

On startup, the API automatically performs a migration for the `Books` model using GORM.

## Todo

- [ ] Implement APIs for insert/update/delete
- [ ] Configure environment file
- [ ] Add authentication mechanism
