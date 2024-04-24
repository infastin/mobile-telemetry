FROM golang:1.22.2-alpine3.18 as build
RUN apk add make
WORKDIR /app
COPY . .
RUN make build

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/build/server /app/server
ENTRYPOINT ["/app/server"]
