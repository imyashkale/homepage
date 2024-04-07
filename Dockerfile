FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o homepage .

FROM alpine:latest

WORKDIR /app

ENV GIN_MODE=release

COPY --from=build /app/homepage .

EXPOSE 80

CMD ["./myapp"]
