
FROM golang:1.17-alpine as build
WORKDIR /app
COPY go.mod /go.sum /app/
RUN go mod download
COPY . /app/
RUN CGO_ENABLED=0 go build ./cmd/main.go

FROM alpine:3.10 as runtime
COPY --from=build /app/main /usr/local/bin/webhook
RUN chmod +x /usr/local/bin/webhook
EXPOSE 8443
ENTRYPOINT ["webhook"]