# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.25.3 AS build-stage

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /http-server

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static AS build-release-stage

WORKDIR /

COPY --from=build-stage /http-server /http-server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/http-server"]