FROM ubuntu:18.04

# Docker builder for Golang
FROM golang:1.19-alpine AS builder

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN  go mod download

# Import the code from the context.
COPY ./ ./

# Create the `/.cache/` directory if it doesn't exist.
RUN mkdir -p /.cache/

# Change the ownership of the `/.cache/` directory to the unprivileged user.
RUN chown nobody:nobody /.cache/

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app 

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8090

# Perform any further action as an unprivileged user.
USER nobody:nobody


# Run the compiled binary.
CMD ["./app"]

RUN go run main.go

