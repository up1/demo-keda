FROM golang:1.13.7-alpine3.11 as build
WORKDIR /src
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-demo-api

FROM alpine:3.11.3
COPY --from=build /src/go-demo-api /
EXPOSE 8080
CMD ["/go-demo-api"]