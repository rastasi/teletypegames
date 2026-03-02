# stage 1: building application binary file
FROM --platform=linux/amd64 golang:1.19-alpine as build

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main ./main.go

# stage 2: copy only the application binary file and necessary files to the alpine container
FROM --platform=linux/amd64 alpine:latest
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/http/openapi.yml .

# run the service on container startup.
CMD ["/app/main"]