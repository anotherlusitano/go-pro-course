# Workout API

A backend written in Go from scratch: a RESTful API for creating and managing workouts, with user registration, token-based authentication, and per-user ownership of data.

Built while following the [Complete Go](https://master.dev/courses/complete-go/) course by Melkey, as a hands-on review of Go fundamentals and a practical introduction to backend development.

## Features

- Full CRUD for workouts, each with an ordered list of exercise entries
- User registration with validated input and bcrypt-hashed passwords
- Token-based authentication (Bearer tokens, stored hashed with an expiry and scope)
- Middleware that authenticates requests and protects routes
- Ownership checks: only the creator of a workout can update or delete it
- SQL migrations managed with [Goose](https://github.com/pressly/goose) and applied on startup
- Unit tests against a dedicated test database
- Health check endpoint

## Tech Stack

- **Language:** Go
- **Router:** [Chi](https://github.com/go-chi/chi)
- **Database:** PostgreSQL (running in Docker) via [pgx](https://github.com/jackc/pgx)
- **Migrations:** [Goose](https://github.com/pressly/goose)
- **Passwords:** bcrypt
- **Tooling:** Docker Compose, cURL, psql

## Architecture

The application is split into layers, each with a single responsibility:

```
.
├── main.go            # Entry point: starts the HTTP server
├── internal/
│   ├── app/           # Application struct: wires stores, handlers, and middleware together
│   ├── routes/        # Route definitions and middleware grouping (Chi)
│   ├── api/           # HTTP handlers (workouts, users, tokens)
│   ├── store/         # Data layer: interfaces + PostgreSQL implementations
│   ├── middleware/    # Authentication and route protection
│   ├── tokens/        # Token generation and hashing
│   └── utils/         # JSON responses, logging, URL param helpers
├── migrations/        # Goose SQL migrations
└── docker-compose.yml # PostgreSQL containers (development and test)
```

Request flow: **router → middleware → handler → store → database**.

Handlers depend on store *interfaces* (`WorkoutStore`, `UserStore`, `TokenStore`) rather than on Postgres directly, so the database layer can be swapped or mocked without touching the application logic.

## Getting Started

### Prerequisites

You need to have installed:

- [Mise](https://mise.jdx.dev/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/download/) (for the `psql` client)

The project uses **Go 1.24.1**, which is installed by Mise.

### Setup

```bash
# Clone the repository
git clone https://github.com/anotherlusitano/go-pro-course.git
cd https://github.com/anotherlusitano/go-pro-course.git

# Install the tools declared in mise.toml (Go, Goose)
mise install
```

### Database

The Docker Compose file starts two PostgreSQL containers:

| Port   | Database    |
| ------ | ----------- |
| `5432` | Production  |
| `5433` | Test        |

```bash
# Start the databases (Postgres)
docker compose up --build

# Run the migrations with Goose to build the database schema
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
goose -dir migrations postgres "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable" up


# Connect to Postgres (the password is "postgres")
psql -U postgres -h localhost -p 5432
```

### Run the app

```bash
go run main.go
```

The API is served at `http://localhost:8080`.

### Run the tests

The tests run against the test database (port `5433`).

```bash
# Run all the tests
go test ./...

# Or run only the store tests
cd internal/store && go test .
```

## API Reference

| Method | Route                    | Auth | Description                    |
| ------ | ------------------------ | ---- | ------------------------------ |
| GET    | `/health`                | No   | Check that the server is up    |
| POST   | `/users`                 | No   | Register a new user            |
| POST   | `/tokens/authentication` | No   | Log in and receive a token     |
| POST   | `/workouts`              | Yes  | Create a workout               |
| GET    | `/workouts/{id}`         | Yes  | Get a workout by ID            |
| PUT    | `/workouts/{id}`         | Yes  | Update a workout (owner only)  |
| DELETE | `/workouts/{id}`         | Yes  | Delete a workout (owner only)  |

Protected routes expect the header `Authorization: Bearer <token>`.

### Examples

**Register a user**

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"username": "alice", "email": "alice@example.com", "password": "secret123", "bio": "Lifting enthusiast"}'
```

**Get a token**

```bash
curl -X POST http://localhost:8080/tokens/authentication \
  -H "Content-Type: application/json" \
  -d '{"username": "alice", "password": "secret123"}'
```

**Create a workout**

```bash
curl -X POST http://localhost:8080/workouts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "title": "Push day",
    "description": "Chest and triceps",
    "duration_minutes": 60,
    "calories_burned": 450,
    "entries": [
      {"exercise_name": "Bench press", "sets": 4, "reps": 8, "weight": 80.0, "notes": "Felt strong", "order_index": 1},
      {"exercise_name": "Plank", "sets": 3, "duration_seconds": 60, "notes": "", "order_index": 2}
    ]
  }'
```

**Get, update, and delete**

```bash
curl http://localhost:8080/workouts/1 -H "Authorization: Bearer <token>"

curl -X PUT http://localhost:8080/workouts/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"title": "Push day (updated)", "duration_minutes": 70, "calories_burned": 500, "entries": []}'

curl -X DELETE http://localhost:8080/workouts/1 -H "Authorization: Bearer <token>"
```

Updating a workout replaces all of its entries with the ones sent in the request.

## What I Learned

- Writing an HTTP server and organizing a Go project into modular packages
- Routing and middleware with Chi, including request context and protected route groups
- Working with PostgreSQL: transactions, rollbacks, and versioned migrations
- Designing a data layer around interfaces to keep it decoupled from the database
- Implementing user registration, password hashing, and token authentication
- Writing unit tests with table-driven testing against a dedicated test database

## Credits

Based on the [Complete Go](https://master.dev/courses/complete-go/) course by Melkey. The original course repository is [Melkeydev/fem-project-live](https://github.com/Melkeydev/fem-project-live).
