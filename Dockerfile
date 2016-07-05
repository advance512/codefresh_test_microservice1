# Standard Dockerfile for Codefresh test microservices

# Instructions are for Docker daemon, Docker Compose configuration available.
# Make sure you have a Docker network called codefresh_test up:
# docker network create --driver bridge codefresh_test

FROM golang:1.6-alpine

# We need gcc to compile native stuff
RUN apk add -U gcc alpine-sdk

# We need git to use go get
RUN apk add -U git

# Make a /microservice directory
RUN mkdir -p /go/src/microservice 
WORKDIR /go/src/microservice

# Add the source files
ADD microservice1.go /go/src/microservice/

# Fetch all the project's dependencies using go get ./...
RUN go get

# Build the project
RUN go build -o main .

EXPOSE 3000 

# Run the executable by default when the container starts.
ENTRYPOINT ["/go/src/microservice/main"]

# Then run:
# docker run --net=codefresh_test -p 3000:3000 --rm --name ms1 -it microservice1
