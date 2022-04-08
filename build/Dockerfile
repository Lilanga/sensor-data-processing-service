#Builder image
FROM golang:alpine3.15 as builder
RUN apk add git
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o sensor-processing-service ./cmd/sensor-data-processing-service

#Runner image
FROM alpine:3.15
COPY --from=builder /build/sensor-processing-service .

ENTRYPOINT [ "./sensor-processing-service" ]