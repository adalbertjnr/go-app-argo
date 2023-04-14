FROM golang:alpine3.17 AS build

WORKDIR /app
COPY ./src/go.mod ./src/go.sum .
COPY ./src/main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:3.17
WORKDIR /app
RUN addgroup --system alpinegroup && adduser --system alpineuser -g alpinegroup
RUN chown -R alpineuser:alpinegroup /app
USER alpineuser
COPY --from=build /app/app .
EXPOSE 5000
CMD [ "./app" ]