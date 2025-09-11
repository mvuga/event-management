FROM golang:1.24.5-alpine3.22 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage
ENV GIN_MODE=release
WORKDIR /
COPY --from=build-stage /docker-gs-ping /docker-gs-ping
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/docker-gs-ping"]
