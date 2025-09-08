# build
FROM golang:1.24 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/airthings_exporter ./cmd/airthings_exporter

# final
FROM gcr.io/distroless/static:nonroot
USER nonroot:nonroot
COPY --from=build /out/airthings_exporter /airthings_exporter
EXPOSE 9000
ENTRYPOINT ["/airthings_exporter"]
