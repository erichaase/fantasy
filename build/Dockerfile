FROM golang:1.15 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server github.com/erichaase/fantasy/cmd/server

FROM alpine:3.12
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server /server
COPY --from=builder /app/web /web

CMD ["/server"]