FROM golang:1.25rc1-alpine3.21 AS builder
WORKDIR /auto-complete-service

COPY go.mod go.sum ./
RUN go.mod download

COPY . .
RUN make build

FROM alpine:3.21
WORKDIR /auto-complete-service

ENV CONFIG_DIR=internal/config
COPY --from=builder /auto-complete-service/internal/migrations /auto-complete-service/internal/migrations
COPY --from=builder /auto-complete-service/internal/config /auto-complete-service/internal/config
COPY --from=builder /auto-complete-service/bin/server /auto-complete-service/bin/server
CMD [ "./bin/server" ]
