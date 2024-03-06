FROM golang:1.21.5-alpine as builder

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/app .

FROM cgr.dev/chainguard/static
WORKDIR /app
COPY --from=builder /app/app /app/app
ENTRYPOINT ["/app/app"]
