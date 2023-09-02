FROM golang:1.19-alpine

LABEL maintainer="John Chukwuma <jchukwuma@mac.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/user/main.go

RUN rm -rf cfg pkg Dockerfile Makefile cmd go*

CMD ["./main"]
