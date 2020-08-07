FROM golang:1.15rc1-alpine3.12

WORKDIR /go/src/blog
COPY . .

EXPOSE 80

RUN go install -v ./...

CMD ["blog"]