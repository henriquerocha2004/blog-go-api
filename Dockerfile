FROM golang:1.18

WORKDIR /app

COPY ./domain /app/
COPY ./infra /app/
COPY .gitignore /app/
COPY go.mod /app/
COPY go.sum /app/
COPY main.go /app/


RUN go build -o server

EXPOSE 8080