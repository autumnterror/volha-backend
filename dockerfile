FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/volha-gateway ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/volha-gateway /app/volha-gateway
COPY --from=builder /app /app/cmd/docs

RUN mkdir -p /app/configs
VOLUME /app/configs

EXPOSE 8080
ENTRYPOINT ["sh", "-c", "if [ -f \"/app/configs/${CONFIG_FILE}\" ]; then ./volha-gateway --config /app/configs/${CONFIG_FILE}; else echo \"Error: Config file not found. Please mount your config file to /app/configs/ and set CONFIG_FILE env variable\"; exit 1; fi"]