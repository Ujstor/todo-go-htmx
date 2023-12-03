# Todo app Go-HTMx-Air-Tailwind

Todo-Go is a simple web-based to-do list application built with Go, using the Chi router and HTMx for dynamic content updates without full page reloads. The application allows users to create, mark as done, and delete to-do items. 

![](https://i.imgur.com/4HPp8n0.png)

## Getting Started

To set up Todo-Go using Docker:
1. Clone the repository
2. Copy the `.env.example` file to create a `.env` file and fill in your environment variables
3. Start Docker services using `make docker-run`

The application and the PostgreSQL database will be started. The app will be accessible at `http://localhost:8080`.

## Usage

- Access the Todo-Go app at `http://localhost:8080`.
- You can add, mark as done, and delete to-do items using the web interface.

### Using Docker
When you start the application with Docker Compose, it automatically sets up a PostgreSQL database and connects the Todo-Go application to it. The database schema will be initialized based on the `init.sql` script.

To stop the application and remove the containers, you can use: `make docker-down`

To persist the PostgreSQL data between runs, a Docker volume `psql_volume` is used as defined in the `docker-compose.yml`

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB and App container
```bash
make docker-run
```

Shutdown containers
```bash
make docker-down
```

live reload the application
```bash
make watch
```

clean up binary from the last build
```bash
make clean
```