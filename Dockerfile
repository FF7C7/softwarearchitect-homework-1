FROM golang:1.14-alpine as builder
WORKDIR /app
COPY . .
WORKDIR /app/bin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ../cmd/...

FROM alpine:latest  
RUN addgroup -S http-server && adduser -S http-server -G http-server
USER http-server
WORKDIR /
COPY --from=builder /app/bin/* .
EXPOSE 8000
CMD ["/http-server"] 
