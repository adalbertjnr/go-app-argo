FROM golang:alpine3.17 AS builder

WORKDIR /app
COPY ./src/go.mod ./src/go.sum .
RUN go mod download
COPY ./src .                
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.17
WORKDIR /app
RUN addgroup --system alpinegroup && adduser --system alpineuser -g alpinegroup
RUN chown -R alpineuser:alpinegroup /app
USER alpineuser
COPY --from=builder /app/ .
EXPOSE 5000
CMD [ "./app" ]