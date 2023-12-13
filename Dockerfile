FROM golang:1.21 AS build

RUN mkdir /app
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./src/main.go

FROM alpine:latest AS production
RUN mkdir /app
COPY --from=build /app/app /app
WORKDIR /app
CMD ["./app"]
