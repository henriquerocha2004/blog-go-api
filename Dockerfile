FROM golang:1.18

WORKDIR /app
COPY . ./

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

ENTRYPOINT [ "air" ]