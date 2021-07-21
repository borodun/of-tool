FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/borodun/tick-counter
WORKDIR $GOPATH/src/borodun/tick-counter/cmd
RUN go get -d -v
RUN go build -o /tmp/tick-counter

FROM alpine
RUN addgroup -S appgroup && adduser -S appuser -G appgroup && mkdir -p /app
COPY --from=builder /tmp/tick-counter /app
RUN chmod a+rx /app/tick-counter

USER appuser
WORKDIR /app
ENV LISTENING_PORT 8080

CMD ["./tick-counter"]