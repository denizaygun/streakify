FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum C:/Go/streakify/
WORKDIR C:/Go/streakify/
RUN go mod download
COPY . C:/Go/streakify
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/streakify streakify

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder C:/Go/streakify/build/streakify /usr/bin/streakify
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/streakify"]