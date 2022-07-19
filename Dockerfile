FROM golang:alpine as builder
WORKDIR $GOPATH/src/build
COPY go.mod go.sum .
RUN go get ./...
ADD . .
RUN go build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/build/appserver .
RUN mkdir public
ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080
CMD ["./appserver"]
