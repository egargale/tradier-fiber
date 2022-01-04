FROM golang:1.16.3-alpine3.13 AS builder

RUN apk --no-cache add \
    alpine-sdk \
    librdkafka-dev \
    pkgconf && \
    rm -rf /var/cache/apk/*

WORKDIR /build/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go install -a -installsuffix cgo -ldflags "-s -w" -tags 'postgres' \
    github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#-

FROM alpine:latest

WORKDIR /api/
ENV PATH=/api/bin/:$PATH
RUN apk --no-cache add postgresql-client \
    ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/migrate ./bin/migrate
COPY --from=builder /build/env.example .
COPY --from=builder /build/db/ .

# EXPOSE 3300
CMD ["/bin/sh"]
# ENTRYPOINT [ "/bin/bash" ]