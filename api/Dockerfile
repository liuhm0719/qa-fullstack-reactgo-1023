FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN GONOSUMDB=* GOFLAGS=-mod=mod go mod download
COPY main.go ./
RUN CGO_ENABLED=0 GONOSUMDB=* GOFLAGS=-mod=mod go build -o /out/app .

FROM alpine:3.20
COPY --from=build /out/app /app
EXPOSE 8080
CMD ["/app"]
