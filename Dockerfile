FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o F1ResultsApi ./cmd/api

RUN chmod +x F1ResultsApi

FROM alpine:latest

WORKDIR /app

COPY .env .

COPY --from=builder /app/F1ResultsApi .

EXPOSE 80

CMD [ "/app/F1ResultsApi" ]
