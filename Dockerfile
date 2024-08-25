FROM golang:1.22.3-alpine3.20 AS builder

WORKDIR /src
COPY . .

RUN go mod download
RUN go build -ldflags "-s -w" -o build/osm-bot cmd/main.go


FROM alpine:3.20.0 AS runner

COPY --from=builder /src/build/osm-bot /osm-bot

ENTRYPOINT [ "/osm-bot" ]
