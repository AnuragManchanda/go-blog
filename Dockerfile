FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go build

EXPOSE 80

# Command to run the executable
CMD ["./go-blog"]
