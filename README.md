# Go Fiber Template

## Requirements

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/dl/) 1.24 or higher

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/gofiber/recipes.git
    cd recipes/auth-docker-postgres-jwt
    ```

2. Set the environment variables in a `.env` file:
    ```env
    DB_PORT=5432
    DB_USER=example_user
    DB_PASSWORD=example_password
    DB_NAME=example_db
    SECRET=example_secret
    ```

3. Build and start the Docker containers:
    ```bash
    docker-compose build
    docker-compose up
    ```

The API and the database should now be running.
