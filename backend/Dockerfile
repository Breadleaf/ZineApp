#
# stage 1: build the go binary
#

# init Golang image
FROM golang:1.24.1-alpine3.21 AS builder
WORKDIR /app

# copy files and install dependencies
# create a cache layer for less frequently changing code
COPY go.mod .
# TODO: uncomment below when go.sum becomes relevant
# COPY go.sum .
RUN go mod download
# create a cache layer for more frequently changing code
COPY . .

# build app for production
RUN CGO_ENABLED=0 go build -o backend .

#
# stage 2: serve the build app using a new Alpine instance
#

# init Alpine image
FROM alpine:3.21
WORKDIR /app

# copy the builder output to Alpine's FS
COPY --from=builder /app/backend .

# expose the port 8080 to be able to access externally
EXPOSE 8080

# run the binary when the container launches
CMD ["./backend"]