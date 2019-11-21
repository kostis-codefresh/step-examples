FROM golang:1.12-alpine AS build_base

WORKDIR /tmp/

COPY . .

# Unit tests
RUN go test -v

# Build the Go app
RUN go build -o ./out/mvncoords .

# Start fresh from a smaller image
FROM alpine:3.9 

COPY --from=build_base /tmp/out/mvncoords /usr/local/bin/mvncoords

CMD ["mvncoords"]