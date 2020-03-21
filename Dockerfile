
ARG GO_VERSION=1.13
ARG LDFLAGS="-s -w"

## Stage 0 - Get latest trusted root certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates

## Stage 1 - Prepare
FROM golang:${GO_VERSION}-alpine AS dev
ENV APP_PATH="/var/app"

# Download all dependencies (cached between runs if go.mod and go.sum are not updated)
WORKDIR ${APP_PATH}
COPY go.mod go.sum ./
RUN go mod download

# Copy source from current dir to working dir
COPY . .

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux
ENTRYPOINT ["sh"]

## Stage 2 - Build
FROM dev as build
ARG LDFLAGS
RUN go build -ldflags="${LDFLAGS}" -o main .
RUN chmod +x main

## Stage 3 - Package
FROM scratch AS prod
ARG GIT_COMMIT
ARG VERSION

LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION
ENV APP_BUILD_PATH="/var/app"
WORKDIR ${APP_BUILD_PATH}

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build ${APP_BUILD_PATH}/main .

ENTRYPOINT [ "./main" ]
