FROM golang:1.14.6-alpine3.12 AS builder


# cd <some directory>
WORKDIR $GOPATH/src/github.com/arschles/gosearch

COPY . .

# CGO_ENABLED=0 is IMPORTANT because it tells Go to assume
# that libc will NOT be enabled. That lets us run the program
# on alpine linux, where libc is missing. it has an alternative
# called musl, which go doesn't support out of the box
RUN GO111MODULE=on CGO_ENABLED=0 go build -o /bin/gosearch .

FROM alpine:3.11.5

COPY --from=builder /bin/gosearch /bin/gosearch

EXPOSE 3000

CMD ["/bin/gosearch"]
