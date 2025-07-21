# Go Fiber Template

## Requirements

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/dl/) 1.24 or higher

## Setup

1. 建立 docker 映像檔:
    ```bash
    docker build -t go-fiber:latest .
    ```

2. 初始化 swagger 文件:
    ```bash
    docker run --rm \
        -v "$(pwd):/usr/src/some-api" \
        -w /usr/src/some-api \
        go-fiber:latest \
        swag init -g cmd/main/main.go -q
    ```
3. 啟動容器
    ```bash
    docker compose up -d
    ```

4. 檢查服務是否啟動

    [http://localhost:9000/livez](http://localhost:9000/livez)

## Migrate

- 建立 migration

    ```
    docker compose exec app goose -dir ./databases/migrations create create_users_table go
    ```