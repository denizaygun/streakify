FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/denizaygun/streakify/
WORKDIR /go/src/github.com/denizaygun/streakify
RUN go mod download
COPY . /go/src/github.com/denizaygun/streakify
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/streakify github.com/denizaygun/streakify

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder  /go/src/github.com/denizaygun/streakify/build/streakify /usr/bin/streakify
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/streakify"]