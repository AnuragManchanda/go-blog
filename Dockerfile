FROM golang:1.12

ENV GO111MODULE=on
ENV PORT=9000
WORKDIR /app/blog
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build
CMD ["blog"]