FROM golang:1.22.0

WORKDIR /app

ENV GOOS=linux
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN make build

EXPOSE ${PORT}

CMD ["./bin/app"]
