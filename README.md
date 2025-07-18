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

2. Build Docker iamge:
    ```bash
    docker build -t go-fiber:latest .
    ```

3. Init swagger docs:
    ```bash
    docker run --rm \
        -v "$(pwd):/usr/src/some-api" \
        -w /usr/src/some-api \
        go-fiber:latest \
        swag init -g cmd/main/main.go -q
    ```
4. Start container
    ```bash
    docker-compose up -d
    ```

5. Check health

    [http://localhost:9000/api/health](http://localhost:9000/api/health)
