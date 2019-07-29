FROM golang:1.12-alpine

ENV GIN_MODE=release

RUN apk add git

WORKDIR /app

ADD . /app

RUN go mod download
RUN go build .


FROM alpine:latest  
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/todo-api .
CMD ["./todo-api"] 