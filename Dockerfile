## We specify the base image we need for our
## go application
FROM golang:1.18.3-buster
MAINTAINER "Matt Townsend"
## We create an /app directory within our
## image that will hold our application source
## files
## We copy everything in the root directory
## into our /app directory
COPY . /app
## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
# Expose port 5050 to the outside world
EXPOSE 5050
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]
