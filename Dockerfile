# Stage 1: Build

FROM golang:1.23.2-alpine AS build

RUN apk add --update \
    curl tar \
    && rm -rf /var/cache/apk/*

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz

WORKDIR /app

COPY go.mod ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# Copy the entire application source code
COPY . .

# Compile the application during build and statically link the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o goshu-server ./cmd/goshu

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM scratch AS release

# Copy the compiled application binary
COPY --from=build /app/goshu-server /goshu-server
COPY --from=build /app/config.yml /config.yml
COPY --from=build /app/views /views

EXPOSE 8080

# Define the command to run the application
ENTRYPOINT ["./goshu-server"]
