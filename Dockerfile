FROM golang:1.24@sha256:1ecc479bc712a6bdb56df3e346e33edcc141f469f82840bab9f4bc2bc41bf91d

# Enviroment variable
WORKDIR /usr/src/some-api

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN go mod download && go mod verify

# Run and expose the server on port 80
EXPOSE 80

# CMD [ "nodemon", "build/app.js" ]
