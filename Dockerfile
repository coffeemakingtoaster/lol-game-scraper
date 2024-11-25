FROM golang:1.23-alpine3.19 AS build 

WORKDIR /build

COPY . .

# CGO for sqlite3
ARG CGO_ENABLED=1
RUN apk add --no-cache --update go gcc g++

RUN go build -o lol-scraper

## -- RUNTIME STAGE --
FROM alpine:3.19 AS runtime 

WORKDIR /app

ARG USER=docker
ARG UID=5432
ARG GID=5433

RUN apk update && apk upgrade
RUN apk add --no-cache sqlite

RUN addgroup -g $GID $USER 

RUN adduser \
    --disabled-password \
    --gecos "" \
    --ingroup "$USER" \
    --no-create-home \
	--uid "$UID" \
    "$USER"

# Copy build with permissions
COPY --from=build --chown=$USER:$USER /build/lol-scraper /app/lol-scraper

# Ensure that backend can be run
RUN chmod +x /app/lol-scraper

USER $USER 

CMD ["/app/lol-scraper"]
