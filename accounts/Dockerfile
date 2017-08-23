# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/context

# copy the local package files to the container workspace
ADD . /go/src/github.com/mattstratton/blondie/accounts

# Setting up working directory
WORKDIR /go/src/github.com/mattstratton/blondie/accounts

# Build the accounts command inside the container.
RUN go install github.com/mattstratton/blondie/accounts

# Run the accounts microservice when the container starts.
ENTRYPOINT /go/bin/accounts

# Service listens on port 8080.
EXPOSE 8080
